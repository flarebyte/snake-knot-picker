# Snake Knot Picker Design Specification

Canonical specification generated from `doc/design-meta/examples` artefacts.

## 01 Overview

Purpose, scope, and design goals.

### 01 Project Intent

What the library optimizes for.

#### Design Goals

1. Keep the schema surface compact.
2. Keep validation strict by default.
3. Avoid ambiguous positional parsing.
4. Keep schemas JSON-compatible for storage and interchange.
5. Support custom validators without a separate mini-language.

#### Validation Library Purpose

`snake-knot-picker` provides strict validation with a compact schema language built from CLI-style argument tokens.

The same schema shape supports both built-in validators and custom registered operators.

## 02 Schema Language

Authoring and transport forms for schema commands.

### 01 Command JSON Shape

JSON-compatible command representation.

#### Args Command JSON Example

```json
{
  "version": "1",
  "commandPath": ["wash", "start"],
  "adminOnly": true,
  "flags": [
    {
      "kind": "boolean",
      "name": "extra-rinse",
      "schema": ["schema", "boolean"]
    },
    {
      "kind": "string",
      "name": "mode",
      "schema": ["schema", "string", "--enum", "normal,delicate,whites", "--required"]
    },
    {
      "kind": "number",
      "name": "spin",
      "schema": ["schema", "number", "--int", "--required"]
    },
    {
      "kind": "tuple",
      "name": "range",
      "schema": [
        "schema",
        "tuple",
        "--size",
        "2",
        "--required"
      ],
      "schemas": [
        ["schema", "number", "--tuple", "0", "--int"],
        ["schema", "number", "--tuple", "1", "--int"]
      ]
    },
    {
      "kind": "string",
      "name": "add",
      "schema": ["schema", "string", "--alphabetic"],
      "schemas": [
        ["schema", "repeatable", "--min-length", "1", "--max-length", "5"]
      ]
    },
    {
      "kind": "tuple",
      "name": "pair",
      "schema": [
        "schema",
        "tuple",
        "--size",
        "2"
      ],
      "schemas": [
        ["schema", "string", "--tuple", "0", "--alphabetic"],
        ["schema", "string", "--tuple", "1", "--hexa"],
        ["schema", "repeatable", "--min-length", "1", "--max-length", "5"]
      ]
    },
    {
      "kind": "number",
      "name": "dose",
      "schema": ["schema", "number", "--int"],
      "schemas": [
        ["schema", "repeatable", "--min-length", "1", "--max-length", "3"]
      ]
    },
    {
      "kind": "string",
      "name": "report",
      "schema": [
        "schema",
        "string",
        "--uri",
        "--scheme",
        "https",
        "--secure",
        "--allow-query",
        "--allow-domains",
        "example.com",
        "--required"
      ]
    },
    {
      "kind": "string",
      "name": "accent",
      "schema": ["schema", "string", "--color", "--format", "hex", "--allow-alpha"]
    },
    {
      "kind": "string",
      "name": "target-arn",
      "schema": [
        "schema",
        "string",
        "--arn",
        "--allow-partition",
        "aws",
        "--allow-service",
        "s3",
        "--allow-service",
        "sns",
        "--allow-region",
        "us-east-2",
        "--allow-account-id",
        "123456789012",
        "--allow-resource",
        "example-sns-topic-name",
        "--required"
      ]
    },
    {
      "kind": "string",
      "name": "whitespace",
      "schema": ["schema", "string", "--whitespace", "--required"]
    },
    {
      "kind": "string",
      "name": "alphabetic",
      "schema": ["schema", "string", "--alphabetic", "--required"]
    },
    {
      "kind": "string",
      "name": "lowercase",
      "schema": ["schema", "string", "--lowercase", "--required"]
    },
    {
      "kind": "string",
      "name": "uppercase",
      "schema": ["schema", "string", "--uppercase", "--required"]
    },
    {
      "kind": "string",
      "name": "punctuation",
      "schema": ["schema", "string", "--punctuation", "--required"]
    },
    {
      "kind": "string",
      "name": "blank",
      "schema": ["schema", "string", "--blank", "--required"]
    },
    {
      "kind": "string",
      "name": "unicode-letter",
      "schema": ["schema", "string", "--unicode-letter", "--required"]
    },
    {
      "kind": "string",
      "name": "unicode-number",
      "schema": ["schema", "string", "--unicode-number", "--required"]
    },
    {
      "kind": "string",
      "name": "unicode-punctuation",
      "schema": ["schema", "string", "--unicode-punctuation", "--required"]
    },
    {
      "kind": "string",
      "name": "unicode-symbol",
      "schema": ["schema", "string", "--unicode-symbol", "--required"]
    },
    {
      "kind": "string",
      "name": "unicode-separator",
      "schema": ["schema", "string", "--unicode-separator", "--required"]
    },
    {
      "kind": "string",
      "name": "latin",
      "schema": ["schema", "string", "--latin", "--required"]
    },
    {
      "kind": "string",
      "name": "han",
      "schema": ["schema", "string", "--han", "--required"]
    },
    {
      "kind": "string",
      "name": "devanagari",
      "schema": ["schema", "string", "--devanagari", "--required"]
    },
    {
      "kind": "string",
      "name": "arabic",
      "schema": ["schema", "string", "--arabic", "--required"]
    },
    {
      "kind": "string",
      "name": "hiragana",
      "schema": ["schema", "string", "--hiragana", "--required"]
    },
    {
      "kind": "string",
      "name": "katakana",
      "schema": ["schema", "string", "--katakana", "--required"]
    },
    {
      "kind": "string",
      "name": "hangul",
      "schema": ["schema", "string", "--hangul", "--required"]
    },
    {
      "kind": "string",
      "name": "tamil",
      "schema": ["schema", "string", "--tamil", "--required"]
    },
    {
      "kind": "string",
      "name": "gujarati",
      "schema": ["schema", "string", "--gujarati", "--required"]
    },
    {
      "kind": "string",
      "name": "ethiopic",
      "schema": ["schema", "string", "--ethiopic", "--required"]
    },
    {
      "kind": "string",
      "name": "postal-code",
      "schema": ["custom", "postal-code", "--country", "US", "--required"]
    },
    {
      "kind": "string",
      "name": "contact-email",
      "schema": [
        "schema",
        "string",
        "--email",
        "--allow-domains",
        "example.com"
      ]
    },
    {
      "kind": "string",
      "name": "release-date-time",
      "schema": [
        "schema",
        "string",
        "--datetime",
        "--layout",
        "RFC3339",
        "--allow-timezone",
        "--location",
        "America/New_York"
      ]
    },
    {
      "kind": "string",
      "name": "retry-duration",
      "schema": [
        "schema",
        "string",
        "--duration",
        "--min-duration",
        "5m",
        "--max-duration",
        "2h"
      ]
    },
    {
      "kind": "string",
      "name": "checksum",
      "schema": ["schema", "string", "--base64"]
    },
    {
      "kind": "string",
      "name": "starts-with",
      "schema": ["schema", "string", "--starts-with", "wash-"]
    },
    {
      "kind": "string",
      "name": "codepoints",
      "schema": ["schema", "string", "--codepoint-range", "U+3040", "U+309F"]
    },
    {
      "kind": "string",
      "name": "hexa",
      "schema": ["schema", "string", "--hexa"]
    },
    {
      "kind": "string",
      "name": "temp",
      "schema": ["schema", "string", "--enum", "cold,warm,hot", "--required"]
    },
    {
      "kind": "string",
      "name": "accent-choice",
      "schema": [
        "schema",
        "string",
        "--enum",
        "red;green;blue",
        "--enum-separator",
        ";"
      ]
    },
    {
      "kind": "string",
      "name": "due-date",
      "schema": ["schema", "string", "--date", "--layout", "ISO8601", "--required"]
    },
    {
      "kind": "string",
      "name": "enabled",
      "schema": ["schema", "string", "--boolean"]
    },
    {
      "kind": "string",
      "name": "alarm-time",
      "schema": ["schema", "string", "--time", "--layout", "HHMMSS", "--required"]
    }
  ]
}
```

### 02 CLI Arguments

Command-line invocation examples.

#### CLI Command Examples

```sh
# Positional arguments
wash start heavy-duty

# String flags
wash start delicate --temp warm

# Integer flags
wash start normal --spin 1200
wash start normal --spin=1200

# Repeatable values
wash --add=bleach,softener,scent-beads
wash --add=bleach --add=softener

# Boolean flags
wash start bedding --extra-rinse

# Repeatable string flags
wash start whites --add bleach --add softener

# Comma-separated values with repeated option flags
wash start --options delicate,extra-rinse --options pre-wash
```

### 03 Admin and User API Types

TypeScript examples for builders and runtime parsing.

#### Args Type Model

```ts
import type {
  NumberValidation,
  StringValidation,
  ValidationError,
} from './common';

export type ArgsSchemaCommand = readonly string[];

// Persisted documents carry only admin-authored schema commands.
// Runtime schemas carry resolved validation objects for direct user input
// validation. Users only provide argv; they never register commands.
export interface ArgsCommandDocument {
  version: string;
  commandPath: readonly string[];
  flags: readonly ArgsDocumentFlag[];
  adminOnly: boolean;
}

// Tuple schema convention:
// - `schema` holds tuple-level shape and requiredness.
// - `schemas` holds per-slot validations and extra flag modifiers.
// - each slot command must include `--tuple <index>`.

export interface ArgsCommandSchema {
  commandPath: readonly string[];
  flags: readonly ArgsFlagSchema[];
  // Marks schemas that are controlled by the admin authoring flow.
  adminOnly: boolean;
}

export type ArgsDocumentFlag =
  | {
      kind: 'boolean';
      name: string;
      schema: ArgsSchemaCommand;
      schemas?: readonly ArgsSchemaCommand[];
    }
  | {
      kind: 'number';
      name: string;
      schema: ArgsSchemaCommand;
      schemas?: readonly ArgsSchemaCommand[];
    }
  | {
      kind: 'string';
      name: string;
      schema: ArgsSchemaCommand;
      schemas?: readonly ArgsSchemaCommand[];
    }
  | {
      kind: 'tuple';
      name: string;
      schema: ArgsSchemaCommand;
      schemas?: readonly ArgsSchemaCommand[];
    };

export interface ArgsParsedCommand {
  commandPath: readonly string[];
  flags: readonly ArgsParsedFlag[];
}

export type ArgsParsedFlag =
  | {
      kind: 'boolean';
      name: string;
      value: true;
    }
  | {
      kind: 'number';
      name: string;
      value: number;
    }
  | {
      kind: 'string';
      name: string;
      value: string;
    }
  | {
      kind: 'tuple';
      name: string;
      values: readonly string[];
    };

export type ArgsParseResult =
  | {
      error: null;
      value: ArgsParsedCommand;
    }
  | {
      error: ValidationError;
      value: null;
    };

export type ArgsFlagSchema =
  | {
      kind: 'boolean';
      name: string;
      schema: ArgsSchemaCommand;
      schemas?: readonly ArgsSchemaCommand[];
    }
  | {
      kind: 'number';
      name: string;
      schema: ArgsSchemaCommand;
      schemas?: readonly ArgsSchemaCommand[];
      validation: NumberValidation;
    }
  | {
      kind: 'string';
      name: string;
      schema: ArgsSchemaCommand;
      schemas?: readonly ArgsSchemaCommand[];
      validation: StringValidation;
    }
  | {
      kind: 'tuple';
      name: string;
      schema: ArgsSchemaCommand;
      schemas?: readonly ArgsSchemaCommand[];
      validations: readonly (StringValidation | NumberValidation)[];
    };

export interface AdminArgsFactory {
  command(commandPath: readonly string[]): AdminArgsCommandBuilder;
}

export interface AdminArgsCommandBuilder {
  adminOnly(): AdminArgsCommandBuilder;
  boolean(
    name: string,
    schemas?: readonly ArgsSchemaCommand[],
  ): AdminArgsCommandBuilder;
  number(
    name: string,
    validation: NumberValidation,
    schemas?: readonly ArgsSchemaCommand[],
  ): AdminArgsCommandBuilder;
  string(
    name: string,
    validation: StringValidation,
    schemas?: readonly ArgsSchemaCommand[],
  ): AdminArgsCommandBuilder;
  tuple(
    name: string,
    validations: readonly (StringValidation | NumberValidation)[],
    schemas?: readonly ArgsSchemaCommand[],
  ): AdminArgsCommandBuilder;
  build(): ArgsCommandSchema;
}

export interface UserArgsValidator {
  // `schema` must come from the trusted admin flow, never from user argv.
  parse(argv: readonly string[], schema: ArgsCommandSchema): ArgsParseResult;
  validate(
    argv: readonly string[],
    schema: ArgsCommandSchema,
  ): ValidationError | null;
}

export declare const adminArgs: AdminArgsFactory;
export declare const userArgs: UserArgsValidator;
```

#### Admin Args Model

```ts
import type { ArgsCommandSchema } from './args';
import { adminArgs } from './args';
import { numberValidations } from './number';
import { stringValidations } from './string';

export const schemaString: ArgsCommandSchema = adminArgs
  .command(['schema', 'string'])
  .adminOnly()
  .string('min-chars', stringValidations.minChars(10))
  .string('max-chars', stringValidations.maxChars(20))
  .string('enum', stringValidations.enum(['green', 'orange', 'red']), [
    ['schema', 'string', '--enum', 'green,orange,red', '--required'],
  ])
  .string(
    'enum-separator',
    stringValidations.enum(['green', 'orange', 'red'], { separator: ';' }),
    [
      [
        'schema',
        'string',
        '--enum',
        'green;orange;red',
        '--enum-separator',
        ';',
        '--required',
      ],
    ],
  )
  .build();

export const washStartSchema: ArgsCommandSchema = adminArgs
  .command(['wash', 'start'])
  .adminOnly()
  .string('mode', stringValidations.enum(['normal', 'delicate', 'whites']), [
    ['schema', 'string', '--enum', 'normal,delicate,whites', '--required'],
  ])
  .boolean('extra-rinse')
  .number('spin', numberValidations.int(), [
    ['schema', 'number', '--int', '--required'],
  ])
  .tuple(
    'range',
    [numberValidations.int(), numberValidations.int()],
    [
      ['schema', 'tuple', '--size', '2', '--required'],
      ['schema', 'number', '--tuple', '0', '--int'],
      ['schema', 'number', '--tuple', '1', '--int'],
    ],
  )
  .string('add', stringValidations.alphabetic(), [
    ['schema', 'repeatable', '--min-length', '1', '--max-length', '5'],
  ])
  .tuple(
    'pair',
    [stringValidations.alphabetic(), stringValidations.hexa()],
    [
      ['schema', 'tuple', '--size', '2'],
      ['schema', 'string', '--tuple', '0', '--alphabetic'],
      ['schema', 'string', '--tuple', '1', '--hexa'],
      ['schema', 'repeatable', '--min-length', '1', '--max-length', '5'],
    ],
  )
  .number('dose', numberValidations.int(), [
    ['schema', 'repeatable', '--min-length', '1', '--max-length', '3'],
  ])
  .string('temp', stringValidations.enum(['cold', 'warm', 'hot']), [
    ['schema', 'string', '--enum', 'cold,warm,hot', '--required'],
  ])
  .build();
```

#### User Args Model

```ts
import type { ArgsCommandSchema, ArgsParseResult } from './args';
import { userArgs } from './args';
import type { ValidationError } from './common';
import { numberValidations } from './number';
import { stringValidations } from './string';

export const washStartArgs = [
  'wash',
  'start',
  '--mode',
  'normal',
  '--spin',
  '1200',
  '--extra-rinse',
] as const;

export const washStartUserSchema: ArgsCommandSchema = {
  commandPath: ['wash', 'start'],
  adminOnly: true,
  flags: [
    {
      kind: 'string',
      name: 'mode',
      schema: [
        'schema',
        'string',
        '--enum',
        'normal,delicate,whites',
        '--required',
      ],
      validation: stringValidations.enum(['normal', 'delicate', 'whites']),
    },
    {
      kind: 'boolean',
      name: 'extra-rinse',
      schema: ['schema', 'boolean'],
    },
    {
      kind: 'number',
      name: 'spin',
      schema: ['schema', 'number', '--int', '--required'],
      validation: numberValidations.int(),
    },
  ],
};

export const washStartParseResult: ArgsParseResult = userArgs.parse(
  washStartArgs,
  washStartUserSchema,
);

export const washStartValidation: ValidationError | null = userArgs.validate(
  washStartArgs,
  washStartUserSchema,
);
```

### 04 Parser Architecture

Internal parser and compiler boundaries from admin schema commands to runtime validators.

#### Parser Architecture

```ts
import type { ValidationError } from './common';
import type {
  ValidationRegistry,
  ValidationSchemaCommand,
} from './validation-registry';

export type FieldSchemaKind = 'boolean' | 'number' | 'string' | 'tuple';

export type SchemaOperatorKind =
  | FieldSchemaKind
  | 'formatter'
  | 'conversion'
  | 'repeatable'
  | 'custom';

export type ParserStage =
  | 'schema-tokenize'
  | 'schema-compile'
  | 'command-register'
  | 'user-argv-parse'
  | 'runtime-validate';

export interface ParsedSchemaFlag {
  name: string;
  values: readonly string[];
}

export interface SchemaCommandAst {
  head: 'schema' | 'custom';
  operator: string;
  flags: readonly ParsedSchemaFlag[];
  raw: ValidationSchemaCommand;
  path: readonly (string | number)[];
  tupleIndex?: number;
}

export interface SchemaCompileContext {
  registry: ValidationRegistry;
  field: string;
  kind: FieldSchemaKind;
  tupleSize?: number;
}

export interface ValidatorSpec {
  operator: string;
  kind: SchemaOperatorKind;
  required: boolean;
  repeatable: boolean;
  tupleIndex?: number;
}

export interface CompiledFieldSchema {
  name: string;
  kind: FieldSchemaKind;
  primary: ValidatorSpec;
  modifiers: readonly ValidatorSpec[];
  tupleSlots: readonly ValidatorSpec[];
}

export type ParserResult<T> =
  | { ok: true; value: T }
  | { ok: false; errors: readonly ValidationError[] };

export interface SchemaCommandParser {
  parse(command: ValidationSchemaCommand): ParserResult<SchemaCommandAst>;
}

export interface SchemaCompiler {
  compile(
    ast: SchemaCommandAst,
    context: SchemaCompileContext,
  ): ParserResult<ValidatorSpec>;
}

export interface CommandSchemaCompiler {
  compileField(field: {
    name: string;
    kind: FieldSchemaKind;
    schema: ValidationSchemaCommand;
    schemas?: readonly ValidationSchemaCommand[];
  }): ParserResult<CompiledFieldSchema>;
}

export interface UserArgvParser {
  parse(
    argv: readonly string[],
    command: readonly CompiledFieldSchema[],
  ): ParserResult<Readonly<Record<string, unknown>>>;
}

export const parserPipeline = [
  'admin schema commands are parsed into SchemaCommandAst',
  'SchemaCommandAst is compiled with the admin-controlled registry',
  'compiled field schemas are registered for a command path',
  'user argv is parsed against compiled field schemas only',
  'runtime validators return typed values or validation errors',
] as const;

export const tupleCompilerExample = {
  field: 'range',
  kind: 'tuple',
  schema: ['schema', 'tuple', '--size', '2', '--required'],
  schemas: [
    ['schema', 'number', '--tuple', '0', '--int'],
    ['schema', 'number', '--tuple', '1', '--int'],
  ],
  compilesTo: {
    primary: 'TupleSize(required=true,size=2)',
    tupleSlots: ['NumberInt(tupleIndex=0)', 'NumberInt(tupleIndex=1)'],
  },
} as const;

export const runtimeBoundaryExample = {
  argv: ['wash', 'start', '--range', '10', '20'],
  uses: 'compiled range field schema',
  forbidden:
    'runtime argv must never compile schema commands or register operators',
} as const;
```

#### Parser Examples and Edge Cases

| case | error_id | expected | input | stage |
| --- | --- | --- | --- | --- |
| simple schema compile |  | Compile one required enum string validator | schema string --enum cold,warm,hot --required | schema-compile |
| happy tuple compile |  | Compile tuple size validator and two indexed slot validators | schema tuple --size 2 + schema number --tuple 0 --int + schema number --tuple 1 --int | schema-compile |
| happy repeatable compile |  | Compile primary string validator plus repeatable modifier | schema string --alphabetic + schema repeatable --min-length 1 --max-length 5 | schema-compile |
| happy enum separator |  | Split enum candidates with explicit separator | schema string --enum red;green;blue --enum-separator ; | schema-compile |
| unknown operator | schema.unknown_operator | Reject schema before command registration | schema frobnicate --required | schema-compile |
| unknown schema flag | schema.unknown_flag | Reject schema before command registration | schema string --bogus | schema-compile |
| missing schema flag value | schema.missing_value | Reject malformed schema tokens | schema string --enum | schema-tokenize |
| enum whitespace definition | schema.enum_whitespace | Reject admin enum candidates with leading or trailing whitespace | schema string --enum cold, warm,hot | schema-compile |
| tuple slot without index | schema.tuple_missing_index | Reject when command appears in tuple schemas without --tuple | schema number --int | schema-compile |
| tuple index out of range | schema.tuple_index_out_of_range | Reject zero-based tuple index outside tuple size | schema number --tuple 2 --int for tuple size 2 | schema-compile |
| duplicate tuple slot | schema.tuple_duplicate_slot | Reject ambiguous tuple slot ownership | schema number --tuple 0 --int + schema string --tuple 0 --hexa | schema-compile |
| missing tuple slot | schema.tuple_missing_slot | Reject tuple schema with uncovered slot | schema tuple --size 2 + schema number --tuple 0 --int | schema-compile |
| user unknown flag | validation.unexpected_flag | Reject unregistered runtime flag | wash start --unknown value | user-argv-parse |
| user schema command attempt | validation.schema_command_forbidden | Reject because users cannot define schemas or register commands | schema string --email | user-argv-parse |
| runtime invalid value | validation.invalid_type | Parse failure or validator failure with field path context | wash start --spin abc | runtime-validate |

## 03 Validation Domains

Core validation capabilities across string, number, tuple, and list domains.

### 01 Validation Capability Matrix

Feature inventory grouped by domain.

#### Validation Matrix

| api | domain | feature | purpose |
| --- | --- | --- | --- |
| MinChars(minChars) | string | min-chars | Require a minimum character count |
| MaxChars(maxChars) | string | max-chars | Limit a maximum character count |
| MinWords(minWords) | string | min-words | Require a minimum word count |
| MaxWords(maxWords) | string | max-words | Limit a maximum word count |
| Enum(allowedValues) | string | enum | Allow only values from a fixed set |
| EnumOptions(separator) | string | enum-separator | Split schema enum values with a custom separator |
| EnumOptions rejectWhitespacePaddedValues | string | enum-trim-check | Reject enum definitions with leading or trailing whitespace |
| EnumOptions rejectEmptyValues | string | enum-empty-check | Reject enum definitions with empty values after trimming |
| Digit | string | digit | Require decimal digits only |
| Whitespace | string | whitespace | Require whitespace characters only |
| Alphabetic | string | alphabetic | Require ASCII alphabetic characters only |
| Lowercase | string | lowercase | Require lowercase letters only |
| Uppercase | string | uppercase | Require uppercase letters only |
| Punctuation | string | punctuation | Require punctuation characters only |
| Base64 | string | base64 | Require base64 encoded text |
| Hexa | string | hexa | Require hexadecimal digits only |
| Blank | string | blank | Require spaces and tabs only |
| Latin | string | latin | Require Latin-script letters only |
| Han | string | han | Require Han characters only |
| Devanagari | string | devanagari | Require Devanagari script characters only |
| Arabic | string | arabic | Require Arabic-script characters only |
| Bengali | string | bengali | Require Bengali script characters only |
| Cyrillic | string | cyrillic | Require Cyrillic-script characters only |
| Hiragana | string | hiragana | Require Hiragana characters only |
| Katakana | string | katakana | Require Katakana characters only |
| Hangul | string | hangul | Require Hangul characters only |
| Telugu | string | telugu | Require Telugu script characters only |
| Gurmukhi | string | gurmukhi | Require Gurmukhi characters only |
| Tamil | string | tamil | Require Tamil script characters only |
| Thai | string | thai | Require Thai script characters only |
| Javanese | string | javanese | Require Javanese script characters only |
| Gujarati | string | gujarati | Require Gujarati script characters only |
| Ethiopic | string | ethiopic | Require Ethiopic-script characters only |
| Kannada | string | kannada | Require Kannada script characters only |
| CodepointRange(from,to) | string | codepoint-range | Restrict characters to a codepoint span |
| Bool | string | boolean | Accept only string booleans |
| Color | string | color | Accept only string colors |
| ColorOptions | string | color-options | Configure color format and alpha channel policy |
| DateString | string | date | Accept only string dates |
| DateOptions | string | date-options | Configure date-only layout policy |
| DateTimeString | string | datetime | Accept only string date-times |
| DateTimeOptions | string | datetime-options | Configure datetime layout timezone and location policy |
| DurationString | string | duration | Accept only string durations |
| DurationOptions | string | duration-options | Configure duration bounds and negative duration policy |
| UnicodeLetter | string | unicode-letter | Require Unicode letters only |
| UnicodeNumber | string | unicode-number | Require Unicode numbers only |
| UnicodePunctuation | string | unicode-punctuation | Require Unicode punctuation only |
| UnicodeSymbol | string | unicode-symbol | Require Unicode symbols only |
| UnicodeSeparator | string | unicode-separator | Require Unicode separators only |
| Uri | string | uri | Accept only string URIs |
| UriOptions | string | uri-options | Configure URL scheme component and host policy checks |
| Arn | string | arn | Accept only string ARNs |
| ArnOptions | string | arn-options | Configure ARN partition service region account and resource allow lists |
| Email | string | email | Accept only string email addresses |
| EmailOptions | string | email-options | Configure email domain allow lists |
| MatchesFormatter(formatter) | string | matches-formatter | Fail if formatting changes the value |
| NumberStringValidation(numberValidation) | string | number | Parse a string to number then validate it |
| StartsWith(prefix) | string | starts-with | Require a specific prefix |
| TimeString | string | time | Accept only string times |
| TimeOptions | string | time-options | Configure time-only layout fraction and timezone policy |
| Uuid | string | uuid | Accept only string UUIDs |
| stringValidations | string | factory | Create string validators from the factory API |
| StringValidationChain.pipe(...).build() | string | chain | Compose string validators and formatters |
| Trim | formatter | trim | Trim surrounding whitespace |
| Lowercase | formatter | lowercase | Convert text to lowercase |
| Uppercase | formatter | uppercase | Convert text to uppercase |
| stringFormatters | formatter | factory | Create string formatters from the factory API |
| StringFormatterChain.pipe(...).build() | formatter | chain | Compose string formatters |
| Min(min) | number | min | Require a minimum numeric value |
| Max(max) | number | max | Require a maximum numeric value |
| MultipleOf(factor) | number | multiple-of | Require a numeric multiple |
| Int | number | int | Require an integer value |
| Float | number | float | Require a floating-point value |
| ParseInt | number | parse-int | Convert string input to an integer |
| ParseFloat | number | parse-float | Convert string input to a float |
| schema number --tuple(index) | number | tuple | Apply a number schema command to one tuple slot |
| numberConversions | number | converter | Create number converters from the factory API |
| numberValidations | number | factory | Create number validators from the factory API |
| NumberValidationChain.pipe(...).build() | number | chain | Compose number validators |
| TupleOf(validations:StringValidation[]) | tuple | of | Validate tuple items with ordered string validations |
| tupleValidations | tuple | factory | Create tuple validators from the factory API |
| TupleValidationChain.pipe(...).build() | tuple | chain | Compose tuple validators |
| ListOf(itemValidation:StringValidation\|TupleValidation) | list | of | Validate every list item with one validation shape |
| MinLength(minLength) | list | min-length | Require a minimum list length |
| MaxLength(maxLength) | list | max-length | Require a maximum list length |
| listValidations | list | factory | Create list validators from the factory API |
| ListValidationChain.pipe(...).build() | list | chain | Compose list validators |
| ValidationRegistry.register(operator) | registry | register | Register built-in or custom validators and reject collisions |
| ValidationRegistry.resolve(domain,name) | registry | resolve | Look up a registered operator by domain and name |
| ValidationRegistry.has(domain,name) | registry | collision | Detect duplicate registrations before adding a validator |
| Schema modifier --required | registry | required | Mark value-bearing or custom flags as mandatory |
| Schema modifier `schema repeatable` | registry | repeatable | Mark flags as repeatable and constrain occurrence count |

### 02 String Validation Examples

String-focused validators and composition patterns.

#### String Validation API

```ts
import type {
  NumberValidation,
  StringFormatter,
  StringValidation,
  StringValidationChain,
  ValidationError,
  ValidatorOptions,
} from './common';

export interface StringValidationFactory {
  arn(options?: ArnOptions): StringValidation;
  base64(): StringValidation;
  alphabetic(): StringValidation;
  blank(): StringValidation;
  arabic(): StringValidation;
  bengali(): StringValidation;
  boolean(): StringValidation;
  cyrillic(): StringValidation;
  date(options?: DateOptions): StringValidation;
  datetime(options?: DateTimeOptions): StringValidation;
  duration(options?: DurationOptions): StringValidation;
  chain(): StringValidationChain;
  codepointRange(from: string, to: string): StringValidation;
  color(options?: ColorOptions): StringValidation;
  digit(): StringValidation;
  devanagari(): StringValidation;
  ethiopic(): StringValidation;
  email(options?: EmailOptions): StringValidation;
  enum(
    allowedValues: readonly string[],
    options?: EnumOptions,
  ): StringValidation;
  hexa(): StringValidation;
  gurmukhi(): StringValidation;
  han(): StringValidation;
  hangul(): StringValidation;
  hiragana(): StringValidation;
  javanese(): StringValidation;
  gujarati(): StringValidation;
  kannada(): StringValidation;
  latin(): StringValidation;
  katakana(): StringValidation;
  lowercase(): StringValidation;
  maxChars(maxChars: number): StringValidation;
  maxWords(maxWords: number): StringValidation;
  matchesFormatter(formatter: StringFormatter): StringValidation;
  minChars(minChars: number): StringValidation;
  minWords(minWords: number): StringValidation;
  number(numberValidation: NumberValidation): StringValidation;
  punctuation(): StringValidation;
  tamil(): StringValidation;
  telugu(): StringValidation;
  thai(): StringValidation;
  time(options?: TimeOptions): StringValidation;
  startsWith(prefix: string): StringValidation;
  unicodeLetter(): StringValidation;
  unicodeNumber(): StringValidation;
  unicodePunctuation(): StringValidation;
  unicodeSeparator(): StringValidation;
  unicodeSymbol(): StringValidation;
  uri(options?: UriOptions): StringValidation;
  uppercase(): StringValidation;
  whitespace(): StringValidation;
  uuid(): StringValidation;
}

export interface EnumOptions {
  separator?: string;
  rejectWhitespacePaddedValues?: true;
  rejectEmptyValues?: true;
}

export type DateLayout = 'ISO8601';

export interface DateOptions {
  layout?: DateLayout;
}

export type DateTimeLayout = 'RFC3339' | 'RFC1123Z' | 'Unix';

export interface DateTimeOptions {
  layout?: DateTimeLayout;
  allowTimezone?: boolean;
  location?: string;
}

export type TimeLayout = 'HHMMSS' | 'HHMM';

export interface TimeOptions {
  layout?: TimeLayout;
  allowFraction?: boolean;
  allowTimezone?: boolean;
}

export interface DurationOptions {
  minDuration?: string;
  maxDuration?: string;
  allowNegative?: boolean;
}

export interface EmailOptions {
  allowDomains?: readonly string[];
}

export type ColorFormat = 'hex';

export interface ColorOptions {
  format?: ColorFormat;
  allowAlpha?: boolean;
}

export interface UriOptions {
  scheme?: 'http' | 'https';
  secure?: boolean;
  allowPort?: boolean;
  allowUserInfo?: boolean;
  allowFragment?: boolean;
  allowQuery?: boolean;
  allowDomains?: readonly string[];
  allowIp?: boolean;
  allowOpaque?: boolean;
}

export interface ArnOptions {
  allowPartitions?: readonly string[];
  allowServices?: readonly string[];
  allowRegions?: readonly string[];
  allowAccountIds?: readonly string[];
  allowResources?: readonly string[];
}

export declare const stringValidations: StringValidationFactory;

export declare class MinChars implements StringValidation {
  readonly minChars: number;

  constructor(minChars: number);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class MaxChars implements StringValidation {
  readonly maxChars: number;

  constructor(maxChars: number);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class MinWords implements StringValidation {
  readonly minWords: number;

  constructor(minWords: number);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class MaxWords implements StringValidation {
  readonly maxWords: number;

  constructor(maxWords: number);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Enum implements StringValidation {
  readonly allowedValues: readonly string[];
  readonly separator?: string;

  constructor(allowedValues: readonly string[], options?: EnumOptions);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Digit implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Alphabetic implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Base64 implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Blank implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Arabic implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Bengali implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Hexa implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Cyrillic implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class CodepointRange implements StringValidation {
  readonly from: string;
  readonly to: string;

  constructor(from: string, to: string);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Bool implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Color implements StringValidation {
  readonly options?: ColorOptions;

  constructor(options?: ColorOptions);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Devanagari implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Ethiopic implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Lowercase implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Gurmukhi implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Han implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Hangul implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Hiragana implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Javanese implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Gujarati implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Kannada implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Latin implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Katakana implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class DateString implements StringValidation {
  readonly options?: DateOptions;

  constructor(options?: DateOptions);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class DateTimeString implements StringValidation {
  readonly options?: DateTimeOptions;

  constructor(options?: DateTimeOptions);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Uri implements StringValidation {
  readonly options?: UriOptions;

  constructor(options?: UriOptions);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Arn implements StringValidation {
  readonly options?: ArnOptions;

  constructor(options?: ArnOptions);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Email implements StringValidation {
  readonly options?: EmailOptions;

  constructor(options?: EmailOptions);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class StartsWith implements StringValidation {
  readonly prefix: string;

  constructor(prefix: string);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class MatchesFormatter implements StringValidation {
  readonly formatter: StringFormatter;

  constructor(formatter: StringFormatter);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class NumberStringValidation implements StringValidation {
  readonly numberValidation: NumberValidation;

  constructor(numberValidation: NumberValidation);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Punctuation implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Tamil implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Telugu implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Thai implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class TimeString implements StringValidation {
  readonly options?: TimeOptions;

  constructor(options?: TimeOptions);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class UnicodeLetter implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class UnicodeNumber implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class UnicodePunctuation implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class UnicodeSeparator implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class UnicodeSymbol implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Uppercase implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Whitespace implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class DurationString implements StringValidation {
  readonly options?: DurationOptions;

  constructor(options?: DurationOptions);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Uuid implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}
```

#### AWS ARN Validation Logic

| feature | flag | rule |
| --- | --- | --- |
| parsing |  | Parse ARNs with github.com/aws/aws-sdk-go-v2/aws/arn arn.Parse before applying policy checks |
| partition | --allow-partition | If any partition allow-list entries are present require arn.Partition to equal one of them |
| service | --allow-service | If any service allow-list entries are present require arn.Service to equal one of them |
| region | --allow-region | If any region allow-list entries are present require arn.Region to equal one of them |
| account-id | --allow-account-id | If any account ID allow-list entries are present require arn.AccountID to equal one of them |
| resource | --allow-resource | If any resource allow-list entries are present require arn.Resource to equal one of them |
| repeatable-allow-flags |  | Allow-list flags are repeatable so admins can authorize multiple values per ARN component |
| absent-allow-list |  | If no allow-list entries are provided for a component then all values for that component are allowed |

#### String Boolean and Color Validation

```ts
import type { StringValidation } from './common';
import { stringValidations } from './string';

export const booleanString: StringValidation = stringValidations.boolean();
export const colorString: StringValidation = stringValidations.color();
export const alphaHexColorString: StringValidation = stringValidations.color({
  format: 'hex',
  allowAlpha: true,
});
```

#### String Validation Chain

```ts
import type { StringValidation, StringValidationChain } from './common';
import { stringFormatters } from './formatters';
import { stringValidations } from './string';

export const normalizedEmailChain: StringValidationChain = stringValidations
  .chain()
  .pipe(stringValidations.matchesFormatter(stringFormatters.trim()))
  .pipe(stringValidations.minChars(10))
  .pipe(stringValidations.maxChars(40))
  .pipe(stringValidations.email());

export const normalizedEmail: StringValidation = normalizedEmailChain.build();
```

#### String Classes Reference

| api | class | description |
| --- | --- | --- |
| Digit | digit | Matches any digit (0-9) |
| Whitespace | whitespace | Matches any whitespace (space, tab, newline, etc.) |
| Alphabetic | alphabetic | Matches any alphabetic letter (A-Z, a-z) |
| Lowercase | lowercase | Matches any lowercase letter |
| Uppercase | uppercase | Matches any uppercase letter |
| Punctuation | punctuation | Matches punctuation characters |
| Hexa | hexa | Matches hexadecimal digits (0-9, A-F, a-f) |
| Blank | blank | Matches space and tab characters only |
| UnicodeLetter | unicode_letter | Matches any Unicode letter (any language) |
| UnicodeNumber | unicode_number | Matches any Unicode number |
| UnicodePunctuation | unicode_punctuation | Matches any Unicode punctuation |
| UnicodeSymbol | unicode_symbol | Matches any Unicode symbol |
| UnicodeSeparator | unicode_separator | Matches Unicode separator characters (spaces, etc.) |

#### String Classes API Example

```ts
import type { StringValidation } from './common';
import { stringValidations } from './string';

export const digitString: StringValidation = stringValidations.digit();
export const whitespaceString: StringValidation =
  stringValidations.whitespace();
export const alphabeticString: StringValidation =
  stringValidations.alphabetic();
export const lowercaseString: StringValidation = stringValidations.lowercase();
export const uppercaseString: StringValidation = stringValidations.uppercase();
export const punctuationString: StringValidation =
  stringValidations.punctuation();
export const hexaString: StringValidation = stringValidations.hexa();
export const blankString: StringValidation = stringValidations.blank();
export const unicodeLetterString: StringValidation =
  stringValidations.unicodeLetter();
export const unicodeNumberString: StringValidation =
  stringValidations.unicodeNumber();
export const unicodePunctuationString: StringValidation =
  stringValidations.unicodePunctuation();
export const unicodeSymbolString: StringValidation =
  stringValidations.unicodeSymbol();
export const unicodeSeparatorString: StringValidation =
  stringValidations.unicodeSeparator();
```

#### Color Validation Logic

| feature | flag | rule |
| --- | --- | --- |
| format | --format | Supported value is hex and it is the default color format |
| leading-hash |  | Require a leading # character |
| short-hex | hex | Accept #RGB |
| long-hex | hex | Accept #RRGGBB |
| alpha | --allow-alpha | Permit #RGBA and #RRGGBBAA forms only when this flag is present |
| case |  | Accept hexadecimal digits case-insensitively |
| runtime-trim |  | Do not implicitly trim runtime user input before color validation |
| unsupported-formats |  | Reject rgb hsl named colors and CSS color functions |

#### Date Validation Logic

| feature | flag | rule |
| --- | --- | --- |
| parsing |  | Parse date strings with Go standard library time package before applying policy checks |
| layout | --layout | Select a named date-only layout parser; supported value is ISO8601 |
| iso8601 | ISO8601 | Parse with Go layout 2006-01-02 |
| timezone |  | Reject timezone-bearing date input |
| location |  | Do not accept --location for date-only validation |
| date-only |  | Reject input containing time-of-day fields |

#### String Date and Time Validation

```ts
import type { StringValidation } from './common';
import { stringValidations } from './string';

export const dateString: StringValidation = stringValidations.date();
export const isoDateString: StringValidation = stringValidations.date({
  layout: 'ISO8601',
});
export const dateTimeString: StringValidation = stringValidations.datetime();
export const newYorkDateTimeString: StringValidation =
  stringValidations.datetime({
    layout: 'RFC3339',
    allowTimezone: true,
    location: 'America/New_York',
  });
export const timeString: StringValidation = stringValidations.time();
export const hourMinuteTimeString: StringValidation = stringValidations.time({
  layout: 'HHMM',
});
export const fractionalTimeString: StringValidation = stringValidations.time({
  layout: 'HHMMSS',
  allowFraction: true,
});
export const durationString: StringValidation = stringValidations.duration();
export const boundedDurationString: StringValidation =
  stringValidations.duration({
    minDuration: '5m',
    maxDuration: '2h',
  });
```

#### Datetime Validation Logic

| feature | flag | rule |
| --- | --- | --- |
| parsing |  | Parse datetime strings with Go standard library time package before applying policy checks |
| layout | --layout | Select a named layout parser; supported values are RFC3339 RFC1123Z and Unix |
| rfc3339 | RFC3339 | Parse with time.RFC3339 and require timezone-bearing input |
| rfc1123z | RFC1123Z | Parse with time.RFC1123Z and require timezone-bearing input |
| unix | Unix | Parse numeric Unix epoch seconds and do not consume a timezone from the input |
| timezone | --allow-timezone | Permit datetime input containing an explicit timezone or numeric offset |
| location | --location | Load the expected IANA timezone with time.LoadLocation such as America/New_York |
| location-without-timezone | --location | Use the loaded location when parsing timezone-less datetime input |
| location-with-timezone | --allow-timezone + --location | When input contains a timezone require it to be consistent with the configured location |

#### Duration Validation Logic

| feature | flag | rule |
| --- | --- | --- |
| parsing |  | Parse duration strings with Go standard library time.ParseDuration before applying policy checks |
| syntax |  | Accept Go duration syntax such as 300ms 1.5h and 2h45m |
| iso8601 |  | Reject ISO8601 duration syntax such as P1DT2H |
| minimum | --min-duration | Parse the bound with time.ParseDuration and require input duration to be greater than or equal to it |
| maximum | --max-duration | Parse the bound with time.ParseDuration and require input duration to be less than or equal to it |
| negative | --allow-negative | By default reject negative durations; permit them only when this flag is present |
| empty |  | Reject empty duration strings |

#### Email Validation Logic

| feature | flag | rule |
| --- | --- | --- |
| parsing |  | Parse email addresses with Go standard library net/mail ParseAddress before applying policy checks |
| single-address |  | Require exactly one email address |
| bare-address |  | Reject display names comments and angle-address forms; accept only the bare addr-spec |
| runtime-trim |  | Do not implicitly trim runtime user input before email parsing |
| domain-normalization |  | Lowercase the domain before allow-list comparison |
| domain-allow-list | --allow-domains | If provided require the email domain to equal an allowed domain or end with "." plus an allowed domain |
| repeatable-domains | --allow-domains | Allow this flag to repeat so admins can authorize multiple domains |
| absent-domain-list |  | If no allowed domains are provided then all syntactically valid domains are allowed |
| network-checks |  | Do not perform MX DNS lookup SMTP validation or other network checks |

#### Enum Validation Logic

| feature | flag | rule |
| --- | --- | --- |
| separator | --enum-separator | Split the --enum value with the configured separator or comma by default |
| trim-check |  | Trim each split enum candidate during admin schema parsing |
| whitespace-error |  | Reject the admin schema if any enum candidate has leading or trailing whitespace |
| empty-error |  | Reject the admin schema if any enum candidate is empty after trimming |
| runtime-values |  | Runtime user values are matched against the validated enum values without implicit trimming |

#### String Formatter Check

```ts
import type { StringFormatter, StringValidation } from './common';
import { stringFormatters } from './formatters';
import { stringValidations } from './string';

export const normalizedUserName: StringValidation =
  stringValidations.matchesFormatter(stringFormatters.trim());
export const normalizedUserNameFormatter: StringFormatter =
  stringFormatters.trim();
```

#### String Languages Reference

| api | class | description |
| --- | --- | --- |
| Latin | latin | Covers Latin-script languages like English and Spanish |
| Han | han | Covers Han characters used in Chinese and Japanese Kanji |
| Devanagari | devanagari | Covers Hindi and related Indo-Aryan languages |
| Arabic | arabic | Covers Arabic-script languages such as Arabic and Urdu |
| Bengali | bengali | Covers Bengali/Bangla script |
| Cyrillic | cyrillic | Covers Russian and other Cyrillic-script languages |
| Hiragana | hiragana | Covers Japanese Hiragana |
| Katakana | katakana | Covers Japanese Katakana |
| Hangul | hangul | Covers Korean Hangul |
| Telugu | telugu | Covers Telugu script |
| Gurmukhi | gurmukhi | Covers Punjabi when written in Gurmukhi |
| Tamil | tamil | Covers Tamil script |
| Thai | thai | Covers Thai script |
| Javanese | javanese | Covers traditional Javanese script |
| Gujarati | gujarati | Covers Gujarati script |
| Ethiopic | ethiopic | Covers Amharic and related Ethiopic-script languages |
| Kannada | kannada | Covers Kannada script |

#### String Languages API Example

```ts
import type { StringValidation } from './common';
import { stringValidations } from './string';

export const latinString: StringValidation = stringValidations.latin();
export const hanString: StringValidation = stringValidations.han();
export const devanagariString: StringValidation =
  stringValidations.devanagari();
export const arabicString: StringValidation = stringValidations.arabic();
export const bengaliString: StringValidation = stringValidations.bengali();
export const cyrillicString: StringValidation = stringValidations.cyrillic();
export const hiraganaString: StringValidation = stringValidations.hiragana();
export const katakanaString: StringValidation = stringValidations.katakana();
export const hangulString: StringValidation = stringValidations.hangul();
export const teluguString: StringValidation = stringValidations.telugu();
export const gurmukhiString: StringValidation = stringValidations.gurmukhi();
export const tamilString: StringValidation = stringValidations.tamil();
export const thaiString: StringValidation = stringValidations.thai();
export const javaneseString: StringValidation = stringValidations.javanese();
export const gujaratiString: StringValidation = stringValidations.gujarati();
export const ethiopicString: StringValidation = stringValidations.ethiopic();
export const kannadaString: StringValidation = stringValidations.kannada();
```

#### String Number Validation

```ts
import type { NumberValidation, StringValidation } from './common';
import { numberValidations } from './number';
import { stringValidations } from './string';

export const numericString: StringValidation = stringValidations.number(
  numberValidations.int(),
);
export const numericStringValidation: NumberValidation =
  numberValidations.int();
```

#### Time Validation Logic

| feature | flag | rule |
| --- | --- | --- |
| parsing |  | Parse time-only strings with Go standard library time package before applying policy checks |
| layout | --layout | Select a named time-only layout parser; supported values are HHMMSS and HHMM |
| hhmmss | HHMMSS | Parse with Go layout 15:04:05 and use this as the default layout |
| hhmm | HHMM | Parse with Go layout 15:04 |
| fraction | --allow-fraction | Permit fractional seconds only with HHMMSS layout |
| timezone | --allow-timezone | Permit time-only input containing an explicit timezone or numeric offset |
| location |  | Do not accept --location for time-only validation |
| date-fields |  | Reject input containing date fields |

#### URL Validation Logic

| feature | flag | rule |
| --- | --- | --- |
| parsing |  | Parse URLs with Go standard library net/url before applying policy checks |
| scheme | --scheme | Configured scheme value must be http or https; reject any other URL scheme |
| secure | --secure | Require scheme to be exactly https |
| port | --allow-port | Permit URL.Host to contain an explicit port |
| user-info | --allow-user-info | Permit URL.User authentication credentials |
| fragment | --allow-fragment | Permit URL.Fragment to be non-empty |
| query | --allow-query | Permit URL.RawQuery to be non-empty |
| domain-whitelist | --allow-domains | When provided require URL.Hostname() to equal an allowed domain or end with "." plus an allowed domain |
| ip-addresses | --allow-ip | By default reject IP host literals; when allowed require net.ParseIP(URL.Hostname()) to succeed for IPv4 or IPv6 hosts |
| opaque | --allow-opaque | By default reject opaque URLs; when allowed permit net/url parses Opaque instead of Host |

### 03 Number Validation Examples

Numeric validators, conversion, and chaining.

#### Number Validation API

```ts
import type {
  NumberConversion,
  NumberValidation,
  NumberValidationChain,
  ValidationError,
  ValidatorOptions,
} from './common';

export interface NumberConversionFactory {
  float(): NumberConversion;
  int(): NumberConversion;
}

export interface NumberValidationFactory {
  chain(): NumberValidationChain;
  float(): NumberValidation;
  int(): NumberValidation;
  max(max: number): NumberValidation;
  min(min: number): NumberValidation;
  multipleOf(factor: number): NumberValidation;
}

export declare const numberConversions: NumberConversionFactory;
export declare const numberValidations: NumberValidationFactory;

export declare class ParseInt implements NumberConversion {
  convert(input: string, opts: ValidatorOptions): number | null;
}

export declare class ParseFloat implements NumberConversion {
  convert(input: string, opts: ValidatorOptions): number | null;
}

export declare class Min implements NumberValidation {
  readonly min: number;

  constructor(min: number);

  validate(input: number, opts: ValidatorOptions): ValidationError | null;
}

export declare class Max implements NumberValidation {
  readonly max: number;

  constructor(max: number);

  validate(input: number, opts: ValidatorOptions): ValidationError | null;
}

export declare class MultipleOf implements NumberValidation {
  readonly factor: number;

  constructor(factor: number);

  validate(input: number, opts: ValidatorOptions): ValidationError | null;
}

export declare class Int implements NumberValidation {
  validate(input: number, opts: ValidatorOptions): ValidationError | null;
}

export declare class Float implements NumberValidation {
  validate(input: number, opts: ValidatorOptions): ValidationError | null;
}
```

#### Number Validation Chain

```ts
import type { NumberValidation, NumberValidationChain } from './common';
import { numberValidations } from './number';

export const boundedIntChain: NumberValidationChain = numberValidations
  .chain()
  .pipe(numberValidations.min(10))
  .pipe(numberValidations.max(20))
  .pipe(numberValidations.int());

export const boundedInt: NumberValidation = boundedIntChain.build();
```

#### Number Converter Examples

```ts
import type { NumberConversion } from './common';
import { numberConversions } from './number';

export const parseIntConversion: NumberConversion = numberConversions.int();
export const parseFloatConversion: NumberConversion = numberConversions.float();
```

### 04 Tuple and List Examples

Collection and positional validation patterns.

#### List Validation API

```ts
import type {
  ListItemValidation,
  ListValidation,
  ListValidationChain,
  ValidationError,
  ValidatorOptions,
} from './common';

export interface ListValidationFactory {
  chain(): ListValidationChain;
  of(itemValidation: ListItemValidation): ListValidation;
  minLength(minLength: number): ListValidation;
  maxLength(maxLength: number): ListValidation;
}

export declare const listValidations: ListValidationFactory;

export declare class ListOf implements ListValidation {
  readonly itemValidation: ListItemValidation;

  constructor(itemValidation: ListItemValidation);

  validate(
    input: readonly unknown[],
    opts: ValidatorOptions,
  ): ValidationError | null;
}

export declare class MinLength implements ListValidation {
  readonly minLength: number;

  constructor(minLength: number);

  validate(
    input: readonly unknown[],
    opts: ValidatorOptions,
  ): ValidationError | null;
}

export declare class MaxLength implements ListValidation {
  readonly maxLength: number;

  constructor(maxLength: number);

  validate(
    input: readonly unknown[],
    opts: ValidatorOptions,
  ): ValidationError | null;
}
```

#### List Validation Chain

```ts
import type { ListValidation, ListValidationChain } from './common';
import { listValidations } from './list';
import { stringValidations } from './string';

export const boundedListChain: ListValidationChain = listValidations
  .chain()
  .pipe(listValidations.of(stringValidations.email()))
  .pipe(listValidations.minLength(1))
  .pipe(listValidations.maxLength(5));

export const boundedList: ListValidation = boundedListChain.build();
```

#### Tuple Validation API

```ts
import type {
  StringValidation,
  TupleValidation,
  TupleValidationChain,
  ValidationError,
  ValidatorOptions,
} from './common';

export interface TupleValidationFactory {
  chain(): TupleValidationChain;
  of(validations: readonly StringValidation[]): TupleValidation;
}

export declare const tupleValidations: TupleValidationFactory;

export declare class TupleOf implements TupleValidation {
  readonly validations: readonly StringValidation[];

  constructor(validations: readonly StringValidation[]);

  validate(
    input: readonly unknown[],
    opts: ValidatorOptions,
  ): ValidationError | null;
}
```

#### Tuple Validation Chain

```ts
import type { TupleValidation, TupleValidationChain } from './common';
import { stringFormatters } from './formatters';
import { stringValidations } from './string';
import { tupleValidations } from './tuple';

export const constrainedTupleChain: TupleValidationChain = tupleValidations
  .chain()
  .pipe(
    tupleValidations.of([
      stringValidations.minChars(2),
      stringValidations.matchesFormatter(stringFormatters.trim()),
    ]),
  );

export const constrainedTuple: TupleValidation = constrainedTupleChain.build();
```

### 05 Formatter Examples

Formatter primitives and composition.

#### Formatter Chain

```ts
import type { StringFormatter, StringFormatterChain } from './common';
import { stringFormatters } from './formatters';

export const normalizedFormatterChain: StringFormatterChain = stringFormatters
  .chain()
  .pipe(stringFormatters.trim())
  .pipe(stringFormatters.lowercase());

export const normalizedFormatter: StringFormatter =
  normalizedFormatterChain.build();
```

#### Formatter Inventory

| Trim | string | trim |
| --- | --- | --- |
| Lowercase | string | lowercase |
| Uppercase | string | uppercase |
| stringFormatters | string | factory |
| StringFormatterChain.pipe(...).build() | string | formatter-chain |

#### Formatter API

```ts
import type { StringFormatter, StringFormatterChain } from './common';

export interface StringFormatterFactory {
  chain(): StringFormatterChain;
  lowercase(): StringFormatter;
  trim(): StringFormatter;
  uppercase(): StringFormatter;
}

export declare const stringFormatters: StringFormatterFactory;

export declare class Trim implements StringFormatter {
  format(input: string): string;
}

export declare class Lowercase implements StringFormatter {
  format(input: string): string;
}

export declare class Uppercase implements StringFormatter {
  format(input: string): string;
}
```

### 06 Registry and Extensibility

Operator registration, custom validators, and collision handling.

#### Error Strategy and IDs

| area | error_id | message |
| --- | --- | --- |
| strategy | error.id.stable | Error IDs are stable public API and must not change once released |
| strategy | error.message.template | Messages are templates rendered from params and should not be used for program control |
| strategy | error.collect.all | Collect all validation failures for a field when practical instead of failing at the first validator |
| strategy | error.admin.kind | Admin schema parsing and registration failures use kind schema |
| strategy | error.user.kind | Runtime argv parsing and value validation failures use kind validation |
| strategy | error.path | Use path for command path flag name tuple index or list index context |
| schema | schema.unknown_operator | Unknown schema operator |
| schema | schema.unknown_flag | Unknown schema flag for the selected operator |
| schema | schema.missing_value | Schema flag requires a following value |
| schema | schema.invalid_value | Schema flag value is malformed |
| schema | schema.invalid_combination | Schema flags cannot be used together |
| schema | schema.duplicate_registration | Validation operator is already registered |
| schema | schema.enum_whitespace | Enum value has leading or trailing whitespace |
| schema | schema.enum_empty | Enum value is empty after trimming |
| schema | schema.tuple_missing_index | Tuple slot schema must include --tuple index |
| schema | schema.tuple_index_out_of_range | Tuple slot index is outside tuple size |
| schema | schema.tuple_duplicate_slot | Tuple slot index is defined more than once |
| schema | schema.tuple_missing_slot | Tuple schema has at least one slot without validation |
| validation | validation.required | Required value is missing |
| validation | validation.unexpected_flag | User argv contains an unknown flag |
| validation | validation.schema_command_forbidden | User argv must not contain schema authoring commands |
| validation | validation.invalid_type | User argv value cannot be parsed as the expected type |
| validation | validation.string | Value failed string validation |
| validation | validation.number | Value failed number validation |
| validation | validation.tuple | Tuple value failed validation |
| validation | validation.list | List value failed validation |
| validation | validation.format | Value failed format validation |
| validation | validation.range | Value is outside the allowed range |

#### Validation Registry API

```ts
import type {
  NumberConversion,
  NumberValidation,
  StringFormatter,
  StringValidation,
  TupleValidation,
} from './common';

export type ValidationOperatorDomain =
  | 'string'
  | 'number'
  | 'formatter'
  | 'conversion'
  | 'tuple';

export type ValidationSchemaCommand = readonly string[];

export type ValidationOperator =
  | {
      domain: 'string';
      name: string;
      schema: ValidationSchemaCommand;
      validation: StringValidation;
    }
  | {
      domain: 'number';
      name: string;
      schema: ValidationSchemaCommand;
      validation: NumberValidation;
    }
  | {
      domain: 'formatter';
      name: string;
      schema: ValidationSchemaCommand;
      formatter: StringFormatter;
    }
  | {
      domain: 'conversion';
      name: string;
      schema: ValidationSchemaCommand;
      conversion: NumberConversion;
    }
  | {
      domain: 'tuple';
      name: string;
      schema: ValidationSchemaCommand;
      validation: TupleValidation;
    };

export interface ValidationRegistry {
  has(domain: ValidationOperatorDomain, name: string): boolean;
  register(operator: ValidationOperator): ValidationRegistry;
  resolve(
    domain: ValidationOperatorDomain,
    name: string,
  ): ValidationOperator | null;
}

export declare const validationRegistry: ValidationRegistry;
```

#### Validation Registry Examples

```ts
import { stringFormatters } from './formatters';
import { numberConversions, numberValidations } from './number';
import { stringValidations } from './string';
import { tupleValidations } from './tuple';
import type { ValidationRegistry } from './validation-registry';
import { validationRegistry } from './validation-registry';

export const stringEmailRegistration: ValidationRegistry =
  validationRegistry.register({
    domain: 'string',
    name: 'email',
    schema: ['schema', 'string', '--email', '--allow-domains', 'example.com'],
    validation: stringValidations.email({ allowDomains: ['example.com'] }),
  });

export const numberIntRegistration: ValidationRegistry =
  validationRegistry.register({
    domain: 'number',
    name: 'int',
    schema: ['schema', 'number', '--int'],
    validation: numberValidations.int(),
  });

export const formatterTrimRegistration: ValidationRegistry =
  validationRegistry.register({
    domain: 'formatter',
    name: 'trim',
    schema: ['schema', 'formatter', '--trim'],
    formatter: stringFormatters.trim(),
  });

export const conversionParseIntRegistration: ValidationRegistry =
  validationRegistry.register({
    domain: 'conversion',
    name: 'parse-int',
    schema: ['schema', 'conversion', '--parse-int'],
    conversion: numberConversions.int(),
  });

export const tuplePairRegistration: ValidationRegistry =
  validationRegistry.register({
    domain: 'tuple',
    name: 'pair',
    schema: ['schema', 'tuple', '--size', '2'],
    validation: tupleValidations.of([
      stringValidations.alphabetic(),
      stringValidations.hexa(),
    ]),
  });

export const postalCodeRegistration: ValidationRegistry =
  validationRegistry.register({
    domain: 'string',
    name: 'postal-code',
    schema: ['custom', 'postal-code', '--country', 'US'],
    validation: stringValidations.enum(['US', 'CA']),
  });
```

#### Validation Domain Reference

| api | domain | feature | purpose |
| --- | --- | --- | --- |
| MinChars(minChars) | string | min-chars | Require a minimum character count |
| MaxChars(maxChars) | string | max-chars | Limit a maximum character count |
| MinWords(minWords) | string | min-words | Require a minimum word count |
| MaxWords(maxWords) | string | max-words | Limit a maximum word count |
| Enum(allowedValues) | string | enum | Allow only values from a fixed set |
| EnumOptions(separator) | string | enum-separator | Split schema enum values with a custom separator |
| EnumOptions rejectWhitespacePaddedValues | string | enum-trim-check | Reject enum definitions with leading or trailing whitespace |
| EnumOptions rejectEmptyValues | string | enum-empty-check | Reject enum definitions with empty values after trimming |
| Digit | string | digit | Require decimal digits only |
| Whitespace | string | whitespace | Require whitespace characters only |
| Alphabetic | string | alphabetic | Require ASCII alphabetic characters only |
| Lowercase | string | lowercase | Require lowercase letters only |
| Uppercase | string | uppercase | Require uppercase letters only |
| Punctuation | string | punctuation | Require punctuation characters only |
| Base64 | string | base64 | Require base64 encoded text |
| Hexa | string | hexa | Require hexadecimal digits only |
| Blank | string | blank | Require spaces and tabs only |
| Latin | string | latin | Require Latin-script letters only |
| Han | string | han | Require Han characters only |
| Devanagari | string | devanagari | Require Devanagari script characters only |
| Arabic | string | arabic | Require Arabic-script characters only |
| Bengali | string | bengali | Require Bengali script characters only |
| Cyrillic | string | cyrillic | Require Cyrillic-script characters only |
| Hiragana | string | hiragana | Require Hiragana characters only |
| Katakana | string | katakana | Require Katakana characters only |
| Hangul | string | hangul | Require Hangul characters only |
| Telugu | string | telugu | Require Telugu script characters only |
| Gurmukhi | string | gurmukhi | Require Gurmukhi characters only |
| Tamil | string | tamil | Require Tamil script characters only |
| Thai | string | thai | Require Thai script characters only |
| Javanese | string | javanese | Require Javanese script characters only |
| Gujarati | string | gujarati | Require Gujarati script characters only |
| Ethiopic | string | ethiopic | Require Ethiopic-script characters only |
| Kannada | string | kannada | Require Kannada script characters only |
| CodepointRange(from,to) | string | codepoint-range | Restrict characters to a codepoint span |
| Bool | string | boolean | Accept only string booleans |
| Color | string | color | Accept only string colors |
| ColorOptions | string | color-options | Configure color format and alpha channel policy |
| DateString | string | date | Accept only string dates |
| DateOptions | string | date-options | Configure date-only layout policy |
| DateTimeString | string | datetime | Accept only string date-times |
| DateTimeOptions | string | datetime-options | Configure datetime layout timezone and location policy |
| DurationString | string | duration | Accept only string durations |
| DurationOptions | string | duration-options | Configure duration bounds and negative duration policy |
| UnicodeLetter | string | unicode-letter | Require Unicode letters only |
| UnicodeNumber | string | unicode-number | Require Unicode numbers only |
| UnicodePunctuation | string | unicode-punctuation | Require Unicode punctuation only |
| UnicodeSymbol | string | unicode-symbol | Require Unicode symbols only |
| UnicodeSeparator | string | unicode-separator | Require Unicode separators only |
| Uri | string | uri | Accept only string URIs |
| UriOptions | string | uri-options | Configure URL scheme component and host policy checks |
| Arn | string | arn | Accept only string ARNs |
| ArnOptions | string | arn-options | Configure ARN partition service region account and resource allow lists |
| Email | string | email | Accept only string email addresses |
| EmailOptions | string | email-options | Configure email domain allow lists |
| MatchesFormatter(formatter) | string | matches-formatter | Fail if formatting changes the value |
| NumberStringValidation(numberValidation) | string | number | Parse a string to number then validate it |
| StartsWith(prefix) | string | starts-with | Require a specific prefix |
| TimeString | string | time | Accept only string times |
| TimeOptions | string | time-options | Configure time-only layout fraction and timezone policy |
| Uuid | string | uuid | Accept only string UUIDs |
| stringValidations | string | factory | Create string validators from the factory API |
| StringValidationChain.pipe(...).build() | string | chain | Compose string validators and formatters |
| Trim | formatter | trim | Trim surrounding whitespace |
| Lowercase | formatter | lowercase | Convert text to lowercase |
| Uppercase | formatter | uppercase | Convert text to uppercase |
| stringFormatters | formatter | factory | Create string formatters from the factory API |
| StringFormatterChain.pipe(...).build() | formatter | chain | Compose string formatters |
| Min(min) | number | min | Require a minimum numeric value |
| Max(max) | number | max | Require a maximum numeric value |
| MultipleOf(factor) | number | multiple-of | Require a numeric multiple |
| Int | number | int | Require an integer value |
| Float | number | float | Require a floating-point value |
| ParseInt | number | parse-int | Convert string input to an integer |
| ParseFloat | number | parse-float | Convert string input to a float |
| schema number --tuple(index) | number | tuple | Apply a number schema command to one tuple slot |
| numberConversions | number | converter | Create number converters from the factory API |
| numberValidations | number | factory | Create number validators from the factory API |
| NumberValidationChain.pipe(...).build() | number | chain | Compose number validators |
| TupleOf(validations:StringValidation[]) | tuple | of | Validate tuple items with ordered string validations |
| tupleValidations | tuple | factory | Create tuple validators from the factory API |
| TupleValidationChain.pipe(...).build() | tuple | chain | Compose tuple validators |
| ListOf(itemValidation:StringValidation\|TupleValidation) | list | of | Validate every list item with one validation shape |
| MinLength(minLength) | list | min-length | Require a minimum list length |
| MaxLength(maxLength) | list | max-length | Require a maximum list length |
| listValidations | list | factory | Create list validators from the factory API |
| ListValidationChain.pipe(...).build() | list | chain | Compose list validators |
| ValidationRegistry.register(operator) | registry | register | Register built-in or custom validators and reject collisions |
| ValidationRegistry.resolve(domain,name) | registry | resolve | Look up a registered operator by domain and name |
| ValidationRegistry.has(domain,name) | registry | collision | Detect duplicate registrations before adding a validator |
| Schema modifier --required | registry | required | Mark value-bearing or custom flags as mandatory |
| Schema modifier `schema repeatable` | registry | repeatable | Mark flags as repeatable and constrain occurrence count |

## 04 Use Cases and Limits

Practical command-level scenarios and policy limits.

### 01 Use Case Table

Admin and user scenarios captured in CSV.

#### Use Cases

| description | example | mode |
| --- | --- | --- |
| Support only named values | wash start --mode heavy-duty | user |
| Support boolean flags | wash start --mode bedding --extra-rinse | user |
| Support single-parameter flags | wash start --mode normal --spin 1200 | user |
| Support n-tuple flags | wash start --mode normal --range 800 1200 | user |
| Support repeated string flags by repetition | wash start --mode whites --add bleach --add softener | user |
| Support repeated number flags by repetition | wash start --mode normal --spin 1200 --spin 1400 | user |
| Support string and enum types | wash start --mode delicate --temp warm | user |
| Support number types | wash start --mode normal --spin 1200 | user |
| Support URL or URI types | wash alarm --report https://website.com | user |
| Validate character counts | schema string --min-chars 10 --max-chars 20 | admin |
| Reserve a command for admins | schema admin report --admin-only | admin |
| Validate word counts | schema string --min-words 10 --max-words 20 | admin |
| Validate numeric values | schema number --min 10 --max 20 | admin |
| Expect a numeric multiple | schema number --multiple-of 10 | admin |
| Require a numeric value | schema number --int --required | admin |
| Expect a repeated numeric flag | schema number --int + schema repeatable | admin |
| Expect a tuple numeric value | schema tuple --size 2 --required + schema number --tuple 0 --int + schema number --tuple 1 --int | admin |
| Expect a prefixed string value | schema string --starts-with blue | admin |
| Expect an enum | schema string --enum green,orange,red | admin |
| Expect an enum with custom separator | schema string --enum green;orange;red --enum-separator ; | admin |
| Expect a secure URL | schema string --uri --scheme https --secure --allow-query --allow-domains example.com --required | admin |
| Expect an AWS ARN | schema string --arn --allow-partition aws --allow-service s3 --allow-service sns --allow-region us-east-2 --allow-account-id 123456789012 --allow-resource example-sns-topic-name --required | admin |
| Expect an email from allowed domains | schema string --email --allow-domains example.com --allow-domains example.org --required | admin |
| Expect a color like #F54927 | schema string --color --format hex --required | admin |
| Expect a postal code | custom postal-code --country US --required | admin |
| Expect a boolean-like string | schema string --boolean | admin |
| Expect a UUID string | schema string --uuid | admin |
| Expect a zoned datetime | schema string --datetime --layout RFC3339 --allow-timezone --location America/New_York --required | admin |
| Expect a date-only value | schema string --date --layout ISO8601 --required | admin |
| Expect a time-only value | schema string --time --layout HHMMSS --required | admin |
| Expect a bounded duration | schema string --duration --min-duration 5m --max-duration 2h | admin |
| Expect the first tuple element | schema string --tuple 0 --enum monday,tuesday | admin |
| Expect the second tuple element | schema string --tuple 1 --color | admin |
| Expect a repeated string flag with length bounds | schema string --alphabetic + schema repeatable --min-length 1 --max-length 5 | admin |
| Expect a repeated tuple flag with length bounds | schema tuple --size 2 + schema string --tuple 0 --enum monday,tuesday + schema string --tuple 1 --hexa + schema repeatable --min-length 1 --max-length 5 | admin |
| Expect a repeated numeric flag with length bounds | schema number --int + schema repeatable --min-length 1 --max-length 3 | admin |
| Expect a string with digits | schema string --digit --required | admin |
| Expect a string with whitespace | schema string --whitespace --required | admin |
| Expect a string with alphabetic characters | schema string --alphabetic --required | admin |
| Expect a lowercase string | schema string --lowercase --required | admin |
| Expect an uppercase string | schema string --uppercase --required | admin |
| Expect a string with punctuation | schema string --punctuation --required | admin |
| Expect a string with blank characters | schema string --blank --required | admin |
| Expect a Unicode letter string | schema string --unicode-letter --required | admin |
| Expect a Unicode number string | schema string --unicode-number --required | admin |
| Expect a Unicode punctuation string | schema string --unicode-punctuation --required | admin |
| Expect a Unicode symbol string | schema string --unicode-symbol --required | admin |
| Expect a Unicode separator string | schema string --unicode-separator --required | admin |
| Expect a Latin-script string | schema string --latin --required | admin |
| Expect a Han-script string | schema string --han --required | admin |
| Expect a Devanagari-script string | schema string --devanagari --required | admin |
| Expect an Arabic-script string | schema string --arabic --required | admin |
| Expect a Hiragana string | schema string --hiragana --required | admin |
| Expect a Katakana string | schema string --katakana --required | admin |
| Expect a Hangul string | schema string --hangul --required | admin |
| Expect a Tamil string | schema string --tamil --required | admin |
| Expect a Gujarati string | schema string --gujarati --required | admin |
| Expect an Ethiopic-script string | schema string --ethiopic --required | admin |
| Expect a string with hex digits | schema string --hexa --required | admin |
| Expect a string from the Japanese Hiragana alphabet | schema string --codepoint-range U+3040 U+309F --required | admin |

### 02 Cobra Limit Notes

Command-line framework constraints impacting schema design.

#### Cobra Limits

| note |
| --- |
| Cobra's CSV parsing is limited when the underlying data contains commas |

## 05 Validation Flows

Ordered flow specifications for admin authoring and user runtime validation.

### 01 Flow Scope

Flow intent and boundaries.

#### Flow Boundaries

Admin flow covers schema composition and registry registration.
User flow covers argv parsing, operator resolution, and value validation against a pre-defined admin-authored command schema.
Users must not submit schema commands or register new commands.

#### Flow-Specific Specification Intent

These flow specs make step order explicit for trusted admin schema authoring and untrusted user input validation.

### 02 Admin Authoring Flow

Trusted admin flow from use case to exported schema.

#### Capture Use Case and Constraints

Select expected command behavior, accepted value shapes, and repeatability constraints.

#### Compose Schema Commands

Author argv-style schema commands for each flag and tuple position.

#### Export JSON-Compatible Command Schema

Persist and share the command schema as JSON-friendly arrays of string tokens.

#### Register Built-in or Custom Operators

Register validator operators in the validation registry and reject collisions.

### 03 Admin Flow Graph

Transition graph for admin steps.

- <a id="graph-node-flow-admin-capture-usecase"></a> Capture Use Case and Constraints: Select expected command behavior, accepted value shapes, and repeatability constraints.
  - <a id="graph-node-flow-admin-compose-schema"></a> Compose Schema Commands: Author argv-style schema commands for each flag and tuple position.
    - <a id="graph-node-flow-admin-register-operator"></a> Register Built-in or Custom Operators: Register validator operators in the validation registry and reject collisions.
      - <a id="graph-node-flow-admin-export-schema"></a> Export JSON-Compatible Command Schema: Persist and share the command schema as JSON-friendly arrays of string tokens.

### 04 User Validation Flow

Untrusted user argv flow from parse to validation result.

#### Parse Argv to Typed Flags

Parse argv into boolean, string, number, and tuple flag values.

#### Receive User Argv Input

Collect raw argv tokens from untrusted runtime input.

#### Resolve Validators from Registry

Resolve each schema operator by domain and name from the registry.

#### Return Typed Result or Validation Error

Return a typed parsed command on success, otherwise a validation error payload.

#### Validate Values by Schema

Validate parsed values with string, number, tuple, list, and formatter-aware rules.

### 05 User Flow Graph

Transition graph for user validation steps.

- <a id="graph-node-flow-user-receive-argv"></a> Receive User Argv Input: Collect raw argv tokens from untrusted runtime input.
  - <a id="graph-node-flow-user-parse-argv"></a> Parse Argv to Typed Flags: Parse argv into boolean, string, number, and tuple flag values.
    - <a id="graph-node-flow-user-resolve-operators"></a> Resolve Validators from Registry: Resolve each schema operator by domain and name from the registry.
      - <a id="graph-node-flow-user-validate-values"></a> Validate Values by Schema: Validate parsed values with string, number, tuple, list, and formatter-aware rules.
        - <a id="graph-node-flow-user-return-result"></a> Return Typed Result or Validation Error: Return a typed parsed command on success, otherwise a validation error payload.

