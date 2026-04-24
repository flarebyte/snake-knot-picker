import type {
  ListValidation,
  NumberConversion,
  NumberValidation,
  StringFormatter,
  StringValidation,
  TupleValidation,
} from "./common";

export type ValidationOperatorDomain =
  | "string"
  | "number"
  | "formatter"
  | "conversion"
  | "tuple"
  | "list";

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
    }
  | {
      domain: "list";
      name: string;
      schema: ValidationSchemaCommand;
      validation: ListValidation;
    };

export interface ValidationRegistry {
  has(domain: ValidationOperatorDomain, name: string): boolean;
  register(operator: ValidationOperator): ValidationRegistry;
  resolve(domain: ValidationOperatorDomain, name: string): ValidationOperator | null;
}

export declare const validationRegistry: ValidationRegistry;

// Example registry shapes:
// validationRegistry.register({
//   domain: "string",
//   name: "postal-code",
//   validation: stringValidations.enum(["US", "CA"]),
//   schema: ["schema", "string", "--postal-code", "--country", "US"],
// })
//
// validationRegistry.register({
//   domain: "number",
//   name: "odometer",
//   validation: numberValidations.min(0),
//   schema: ["schema", "number", "--odometer", "--min", "0"],
// })
