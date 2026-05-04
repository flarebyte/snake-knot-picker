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
