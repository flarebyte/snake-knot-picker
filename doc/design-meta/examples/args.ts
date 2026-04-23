import type { NumberValidation, StringValidation, ValidationError } from "./common";

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
      kind: "number";
      name: string;
      value: number;
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
      kind: "number";
      name: string;
      validation: NumberValidation;
      repeatable?: true;
    }
  | {
      kind: "string";
      name: string;
      validation: StringValidation;
      repeatable?: true;
    }
  | {
      kind: "tuple";
      name: string;
      validations: readonly StringValidation[];
      repeatable?: true;
    };

export interface ArgsRepeatableOptions {
  repeatable?: true;
}

export interface AdminArgsFactory {
  command(commandPath: readonly string[]): AdminArgsCommandBuilder;
}

export interface AdminArgsCommandBuilder {
  adminOnly(): AdminArgsCommandBuilder;
  boolean(name: string): AdminArgsCommandBuilder;
  number(name: string, validation: NumberValidation, options?: ArgsRepeatableOptions): AdminArgsCommandBuilder;
  string(name: string, validation: StringValidation, options?: ArgsRepeatableOptions): AdminArgsCommandBuilder;
  tuple(name: string, validations: readonly StringValidation[], options?: ArgsRepeatableOptions): AdminArgsCommandBuilder;
  build(): ArgsCommandSchema;
}

export interface UserArgsValidator {
  parse(argv: readonly string[], schema: ArgsCommandSchema): ArgsParseResult;
  validate(argv: readonly string[], schema: ArgsCommandSchema): ValidationError | null;
}

export declare const adminArgs: AdminArgsFactory;
export declare const userArgs: UserArgsValidator;
