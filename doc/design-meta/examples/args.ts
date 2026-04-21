import type { ListValidation, StringValidation, TupleValidation, ValidationError } from "./common";

export interface ArgsCommandSchema {
  commandPath: readonly string[];
  flags: readonly ArgsFlagSchema[];
  positionals: readonly StringValidation[];
}

export type ArgsFlagSchema =
  | {
      kind: "boolean";
      name: string;
    }
  | {
      kind: "list";
      name: string;
      validation: ListValidation;
    }
  | {
      kind: "string";
      name: string;
      validation: StringValidation;
    }
  | {
      kind: "tuple";
      name: string;
      validation: TupleValidation;
    };

export interface AdminArgsFactory {
  command(commandPath: readonly string[]): AdminArgsCommandBuilder;
}

export interface AdminArgsCommandBuilder {
  boolean(name: string): AdminArgsCommandBuilder;
  list(name: string, validation: ListValidation): AdminArgsCommandBuilder;
  positional(validation: StringValidation): AdminArgsCommandBuilder;
  string(name: string, validation: StringValidation): AdminArgsCommandBuilder;
  tuple(name: string, validation: TupleValidation): AdminArgsCommandBuilder;
  build(): ArgsCommandSchema;
}

export interface UserArgsValidator {
  validate(argv: readonly string[], schema: ArgsCommandSchema): ValidationError | null;
}

export declare const adminArgs: AdminArgsFactory;
export declare const userArgs: UserArgsValidator;
