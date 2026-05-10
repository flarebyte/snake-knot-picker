# Usage

This guide shows the typical runtime usage of `snake-knot-picker`.

The runtime input contract is tokenized argv: `[]string`.

## Quick start

```go
package main

import (
	"fmt"

	picker "github.com/flarebyte/snake-knot-picker"
)

func main() {
	doc := picker.CommandDocument{
		Version:     "1",
		CommandPath: []string{"wash", "start"},
		Flags: []picker.CommandFlagDef{
			{Kind: "string", Name: "mode", Schema: []string{"schema", "string", "--required"}},
			{Kind: "number", Name: "spin", Schema: []string{"schema", "number", "--int"}},
			{Kind: "tuple", Name: "range", Schema: []string{"schema", "tuple", "--size", "2"}, Schemas: [][]string{
				{"schema", "number", "--tuple", "0", "--int"},
				{"schema", "number", "--tuple", "1", "--int"},
			}},
			{Kind: "string", Name: "add", Schema: []string{"schema", "string"}, Schemas: [][]string{
				{"schema", "repeatable", "--min-length", "1", "--max-length", "5"},
			}},
		},
	}

	compiled, err := picker.CompileCommandDocument(doc)
	if err != nil {
		panic(err)
	}

	argv := []string{
		"wash", "start",
		"--mode", "normal",
		"--spin", "1200",
		"--range", "10,20",
		"--add", "soap",
		"--add", "bleach,softener",
	}

	result, err := picker.Validate(compiled, argv)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.CommandPath) // [wash start]
}
```

## Main entry points

- `ParseCommandDocumentJSON(data []byte) (CommandDocument, error)`
- `CompileCommandDocument(doc CommandDocument) (CompiledCommand, error)`
- `CompileCommandDocumentWithOptions(doc CommandDocument, options CompileOptions) (CompiledCommand, error)`
- `Validate(compiled CompiledCommand, argv []string) (*ParseResult, error)`
- `ValidateWithDocument(doc CommandDocument, argv []string) (*ParseResult, error)`
- `ValidateWithDocumentJSON(data []byte, argv []string) (*ParseResult, error)`

## Recommended production flow

1. Load admin-authored JSON once at startup.
2. Compile once with `CompileCommandDocument`.
3. Reuse `CompiledCommand` for each request.
4. Validate incoming tokenized argv with `Validate`.

This keeps runtime fast and avoids re-validating schema definitions on every request.

## Using JSON documents

```go
raw := []byte(`{
  "version":"1",
  "commandPath":["wash","start"],
  "adminOnly":false,
  "flags":[
    {"kind":"string","name":"mode","schema":["schema","string","--required"]},
    {"kind":"number","name":"spin","schema":["schema","number","--int"]}
  ]
}`)

doc, err := picker.ParseCommandDocumentJSON(raw)
if err != nil {
	// JSON/document schema issue
}

compiled, err := picker.CompileCommandDocument(doc)
if err != nil {
	// schema compile/validation issue
}
```

## Using builder API

```go
doc := picker.NewCommandBuilder("wash", "start").
	AddFlag(picker.CommandFlagDef{
		Kind:   "string",
		Name:   "mode",
		Schema: []string{"schema", "string", "--required"},
	}).
	AddFlag(picker.CommandFlagDef{
		Kind:   "number",
		Name:   "spin",
		Schema: []string{"schema", "number", "--int"},
	}).
	Build()
```

## Reading parse output

`ParseResult.Values` is `map[string]Value`.

- `Value.String` for string flags
- `Value.Number` for number flags
- `Value.Bool` for boolean flags
- `Value.Tuple` for tuple flags
- `Value.List` for repeatable flags

Example:

```go
mode := *result.Values["mode"].String
spin := *result.Values["spin"].Number
```

## Error handling

Errors are `*ValidationError` with stable IDs.

```go
result, err := picker.Validate(compiled, argv)
if err != nil {
	if verr, ok := err.(*picker.ValidationError); ok {
		for _, d := range verr.Details {
			// d.ID and d.Kind are stable for programmatic handling
			fmt.Printf("%s (%s): %s\n", d.ID, d.Kind, d.Message)
		}
	}
}
_ = result
```

Common IDs:

- `schema.*` for authoring/compile-time issues
- `validation.*` for runtime argv issues

## Security and input rules

- Runtime input is tokenized argv (`[]string`) only.
- Inline syntax `--key=value` is rejected.
- Values that look like flags (including leading whitespace before `--`) are rejected.
- Flag-like CSV segments are rejected for tuple/repeatable parsing.
- Runtime input cannot include schema authoring commands (`schema`, `custom`).

## Important behavior notes

- Boolean flags are presence-based (`--extra-rinse` => `true`).
- Repeatable flags use repeated tokens (`--add a --add b,c`).
- Tuple parsing uses comma-separated values (`--range 10,20`).
- Command path prefix can be included (`wash start ...`) or omitted (`--mode normal ...`) if the rest matches.
