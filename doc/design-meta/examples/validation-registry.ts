import type {
  NumberConversion,
  NumberValidation,
  StringFormatter,
  StringValidation,
  TupleValidation,
} from "./common";

export type ValidationOperatorDomain = "string" | "number" | "formatter" | "conversion" | "tuple";

export type ValidationSchemaCommand = readonly string[];

export type ValidationOperator =
  | {
      domain: "string";
      name: string;
      schema: ValidationSchemaCommand;
      validation: StringValidation;
    }
  | {
      domain: "number";
      name: string;
      schema: ValidationSchemaCommand;
      validation: NumberValidation;
    }
  | {
      domain: "formatter";
      name: string;
      schema: ValidationSchemaCommand;
      formatter: StringFormatter;
    }
  | {
      domain: "conversion";
      name: string;
      schema: ValidationSchemaCommand;
      conversion: NumberConversion;
    }
  | {
      domain: "tuple";
      name: string;
      schema: ValidationSchemaCommand;
      validation: TupleValidation;
    };

export interface ValidationRegistry {
  has(domain: ValidationOperatorDomain, name: string): boolean;
  register(operator: ValidationOperator): ValidationRegistry;
  resolve(domain: ValidationOperatorDomain, name: string): ValidationOperator | null;
}

export declare const validationRegistry: ValidationRegistry;
