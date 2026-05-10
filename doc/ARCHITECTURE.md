# Architecture

## Scope

`snake-knot-picker` is a Go library that validates tokenized CLI-like input (`[]string`) against an admin-authored command schema.

The code is organized around three main phases:

1. document transport (`json` <-> `CommandDocument`)
2. compile-time validation (`CommandDocument` -> `CompiledCommand`)
3. runtime validation (`CompiledCommand` + user argv -> `ParseResult`)

## Package Layout

### Root package `picker`

- Core public types: [`command.go`](../command.go)
  - `CommandDocument`, `CommandFlagDef`, `CompiledCommand`, `CompiledFlag`, `ParseResult`, `Value`
- Document transport: [`command_document_transport.go`](../command_document_transport.go)
  - `ParseCommandDocumentJSON`, `ToJSON`
- Document compile pipeline:
  - [`command_document_compile.go`](../command_document_compile.go)
  - [`command_document_validate.go`](../command_document_validate.go)
  - [`command_document_rules.go`](../command_document_rules.go)
  - [`command_flag_name_validator.go`](../command_flag_name_validator.go)
- Runtime parser/validator: [`validate.go`](../validate.go)
- Convenience wrappers: [`command_document_runtime.go`](../command_document_runtime.go)

### Internal packages

- `internal/schema`: schema token parsing/compilation helpers and rules
- `internal/validators`: concrete string/number/url/time/email/arn/color validators
- `internal/argv`: thin parser wrapper used by tests/integration points

## End-to-End Data Flow

### 1) Transport: JSON to in-memory document

Input:

```json
{
  "version": "1",
  "commandPath": ["wash", "start"],
  "flags": [
    {"kind":"string","name":"mode","schema":["schema","string","--required"]},
    {"kind":"number","name":"spin","schema":["schema","number","--int"]},
    {"kind":"tuple","name":"range","schema":["schema","tuple","--size","2"],"schemas":[
      ["schema","number","--tuple","0","--int"],
      ["schema","number","--tuple","1","--int"]
    ]}
  ]
}
```

Output:

- `CommandDocument` (Go struct), or schema error if JSON shape is invalid.

### 2) Compile: document to immutable runtime command

Entry points:

- `CompileCommandDocument(doc)`
- `CompileCommandDocumentWithOptions(doc, options)`

Compile responsibilities:

- validates document structure and semantics
- validates flag names (default allowlist, configurable validator factory)
- validates tuple child slots, repeatable placement, duplicate names, etc.
- extracts runtime shape into `CompiledCommand`/`CompiledFlag`

Example output (`CompiledCommand`):

```go
CompiledCommand{
  CommandPath: []string{"wash", "start"},
  Flags: []CompiledFlag{
    {Kind: "string", Name: "mode"},
    {Kind: "number", Name: "spin"},
    {Kind: "tuple", Name: "range", TupleSize: 2},
  },
}
```

### 3) Runtime validation: argv to typed values

Entry points:

- `Validate(compiled, argv)`
- `ValidateWithDocument(doc, argv)`
- `ValidateWithDocumentJSON(data, argv)`

Input (must be tokenized):

```go
[]string{"wash", "start", "--mode", "normal", "--spin", "1200", "--range", "10,20"}
```

Output:

```go
&ParseResult{
  CommandPath: []string{"wash", "start"},
  Values: map[string]Value{
    "mode":  {String: ptr("normal")},
    "spin":  {Number: ptr(1200.0)},
    "range": {Tuple: []Value{{String: ptr("10")}, {String: ptr("20")}}},
  },
}
```

On error:

- returns `*ValidationError` with stable detail IDs and `schema` or `validation` kind.

## Security Boundaries

Runtime parser intentionally enforces strict token rules:

- no `--key=value` syntax
- value tokens that look like flags (including leading whitespace then `--`) are rejected
- flag-like CSV segments inside tuple/repeatable values are rejected
- schema command tokens (`schema`, `custom`) are forbidden in runtime argv

This keeps value-vs-flag interpretation deterministic and reduces injection ambiguity.

## Extension Points

### Custom flag-name policy

`CompileOptions.FlagNameValidator` allows injecting a custom policy.

Provided factories:

- regex-based validator factory
- manual rune-based validator factory

Both support min/max length constraints and can be benchmarked independently.

### Validator registry

[`registry.go`](../registry.go) defines a registry for built-in and custom validator factories by operator name.

## Typical Usage Sequence

1. Admin authoring flow stores `CommandDocument` JSON.
2. Service loads JSON via `ParseCommandDocumentJSON`.
3. Service compiles once with `CompileCommandDocument`.
4. For each request, service validates tokenized argv (`[]string`) with `Validate`.
