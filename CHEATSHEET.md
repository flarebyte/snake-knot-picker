# Cheat Sheet

Quick reference for schema commands, runtime argv patterns, and errors.

## Runtime rules

- Runtime input must be tokenized `[]string`.
- `--key=value` is rejected.
- Values that look like flags (including leading space + `--`) are rejected.
- Runtime input cannot include schema authoring tokens (`schema`, `custom`).

## Built-in operators

From registry:

- `boolean`
- `number`
- `string`
- `tuple`
- `repeatable`
- `postal-code` (custom/built-in registered operator name)

## Schema heads

- `schema <operator> ...`
- `custom <operator> ...`

`custom` currently allows only:

- `--required`
- `--country <value>`

## Supported schema flags

### Arity 0 (switch flags)

- `--required`
- `--int`
- `--secure`
- `--allow-query`
- `--allow-alpha`
- `--alphabetic`
- `--hexa`
- `--whitespace`
- `--lowercase`
- `--uppercase`
- `--punctuation`
- `--blank`
- `--unicode-letter`
- `--unicode-number`
- `--unicode-punctuation`
- `--unicode-symbol`
- `--unicode-separator`
- `--latin`
- `--han`
- `--devanagari`
- `--arabic`
- `--hiragana`
- `--katakana`
- `--hangul`
- `--tamil`
- `--gujarati`
- `--ethiopic`
- `--email`
- `--datetime`
- `--duration`
- `--base64`
- `--boolean`
- `--color`
- `--date`
- `--time`
- `--uri`
- `--arn`
- `--allow-timezone`

### Arity 1 (flag + one value)

- `--enum <value>`
- `--enum-separator <value>`
- `--size <value>`
- `--tuple <value>`
- `--min-length <value>`
- `--max-length <value>`
- `--scheme <value>`
- `--allow-domains <value>`
- `--allow-partition <value>`
- `--allow-service <value>`
- `--allow-region <value>`
- `--allow-account-id <value>`
- `--allow-resource <value>`
- `--country <value>`
- `--layout <value>`
- `--location <value>`
- `--min-duration <value>`
- `--max-duration <value>`
- `--starts-with <value>`
- `--format <value>`
- `--min <value>`
- `--max <value>`
- `--multiple-of <value>`

### Arity 2

- `--codepoint-range <start> <end>`

## Common schema command examples

```text
schema string --required
schema string --enum cold,warm,hot
schema number --int --min 0 --max 2000
schema tuple --size 2 --required
schema number --tuple 0 --int
schema number --tuple 1 --int
schema repeatable --min-length 1 --max-length 5
custom postal-code --country US --required
```

## Runtime argv examples

Valid:

```go
[]string{"wash", "start", "--mode", "normal", "--spin", "1200", "--range", "10,20"}
[]string{"--mode", "normal", "--spin", "1200", "--range", "10,20"} // command path omitted
[]string{"wash", "start", "--add", "soap", "--add", "bleach,softener"}
```

Rejected:

```go
[]string{"wash", "start", "--spin=1200"}        // inline form rejected
[]string{"wash", "start", "--mode", "--x"}      // value looks like flag
[]string{"wash", "start", "schema", "string"}   // schema tokens forbidden at runtime
[]string{"wash", "start", "--unknown", "x"}     // unknown flag
```

## Output shape reminder

- `ParseResult.CommandPath []string`
- `ParseResult.Values map[string]Value`
- `Value.String`, `Value.Number`, `Value.Bool`, `Value.Tuple`, `Value.List`

## Error ID quick map

Schema-side:

- `schema.unknown_operator`
- `schema.unknown_flag`
- `schema.missing_value`
- `schema.invalid_value`
- `schema.invalid_combination`
- `schema.duplicate_registration`
- `schema.enum_whitespace`
- `schema.enum_empty`
- `schema.tuple_missing_index`
- `schema.tuple_index_out_of_range`
- `schema.tuple_duplicate_slot`
- `schema.tuple_missing_slot`

Runtime-side:

- `validation.required`
- `validation.unexpected_flag`
- `validation.schema_command_forbidden`
- `validation.invalid_type`
- `validation.string`
- `validation.number`
- `validation.tuple`
- `validation.list`
- `validation.format`
- `validation.range`
