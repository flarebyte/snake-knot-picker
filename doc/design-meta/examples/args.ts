import type { ListValidation, StringValidation, ValidationError } from "./common";

export interface ArgsCommandSchema {
  commandPath: readonly string[];
  flags: readonly ArgsFlagSchema[];
  adminOnly: boolean;
}

export interface ArgsParsedCommand {
  commandPath: readonly string[];
  flags: readonly ArgsParsedFlag[];
}

export type ArgsParsedFlag =
  | {
      kind: "boolean";
      name: string;
      value: true;
    }
  | {
      kind: "list";
      name: string;
      values: readonly string[];
    }
  | {
      kind: "string";
      name: string;
      value: string;
    }
  | {
      kind: "tuple";
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
      kind: "boolean";
      name: string;
    }
  | {
      kind: "list";
      name: string;
      validation: ListValidation;
      repeatable: true;
    }
  | {
      kind: "string";
      name: string;
      validation: StringValidation;
    }
  | {
      kind: "tuple";
      name: string;
      validations: readonly StringValidation[];
    };

export interface AdminArgsFactory {
  command(commandPath: readonly string[]): AdminArgsCommandBuilder;
}

export interface AdminArgsCommandBuilder {
  adminOnly(): AdminArgsCommandBuilder;
  boolean(name: string): AdminArgsCommandBuilder;
  list(name: string, validation: ListValidation): AdminArgsCommandBuilder;
  string(name: string, validation: StringValidation): AdminArgsCommandBuilder;
  tuple(name: string, validations: readonly StringValidation[]): AdminArgsCommandBuilder;
  build(): ArgsCommandSchema;
}

export interface UserArgsValidator {
  parse(argv: readonly string[], schema: ArgsCommandSchema): ArgsParseResult;
  validate(argv: readonly string[], schema: ArgsCommandSchema): ValidationError | null;
}

export declare const adminArgs: AdminArgsFactory;
export declare const userArgs: UserArgsValidator;
