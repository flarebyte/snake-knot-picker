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
