# snake-knot-picker

![snake-knot-picker illustration](doc/snake-knot-picker-hero.png)

`snake-knot-picker` is a validation library for Go with a compact schema language built from CLI-style arguments.

It is designed for two use cases:

1. trusted admin-authored schemas that define what is allowed
2. strict user-side validation for untrusted input

## What it supports

- String validation for character sets, scripts, formats, prefixes, and custom rules
- Number validation and string-to-number conversion
- Tuple validation for fixed ordered values
- Repeatable flags with length bounds
- Formatters such as trim, lowercase, and uppercase
- Custom validators registered in Go and exposed through the same schema language
- A JSON-compatible schema shape for storing and exchanging command definitions

## Schema language

The schema language uses argv-style tokens so it stays easy to read, store, and parse.

Examples:

```text
schema string --enum cold,warm,hot --required
schema number --int --required
custom postal-code --country US --required
schema repeatable --min-length 1 --max-length 5
```

The same shape is used for built-in validators and custom Go-registered validators.

## Validation building blocks

### String validation

String validation covers:

- character classes like `digit`, `alphabetic`, `whitespace`, `blank`, and `hexa`
- Unicode classes like `unicodeLetter`, `unicodeNumber`, `unicodePunctuation`, `unicodeSymbol`, and `unicodeSeparator`
- script checks like `latin`, `han`, `arabic`, `hiragana`, `katakana`, `hangul`, `tamil`, and others
- format checks like `email`, `uri`, `date`, `datetime`, `time`, `duration`, `color`, and `base64`
- composition checks like `matchesFormatter(...)` and `number(...)`

### Number validation

Number validation supports:

- `min`
- `max`
- `multipleOf`
- `int`
- `float`
- parsing from string with dedicated converters

### Collections

- `tuple` validates fixed-position values
- `list` validates homogeneous collections

Tuple schema authoring guideline:

- Put tuple-level directives in `schema` (for example `schema tuple --size 2 --required`)
- Put each tuple slot validation in `schemas` as separate commands
- Make `--tuple <index>` mandatory in every tuple slot command

### Formatters

Formatters transform strings without being validations themselves.

- `trim`
- `lowercase`
- `uppercase`

You can also validate whether a formatter would change a string.

## Custom validators

Custom validators are registered in Go and exposed through the same schema registry as built-ins.

That means a validator like `postal-code` can be treated as a first-class schema operator rather than a special case.

## Design goals

- keep the schema surface compact
- make validation strict by default
- avoid ambiguous positional parsing
- keep the schema JSON-friendly
- allow custom validators without inventing a separate mini-language

## Project layout

- `doc/design-meta/examples/` contains the source-of-truth examples
- `doc/snake-knot-picker-hero.png` is the README illustration
