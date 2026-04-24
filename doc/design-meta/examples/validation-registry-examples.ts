import type {
  ValidationOperatorDomain,
  ValidationSchemaCommand,
} from "./validation-registry";

export interface ValidationRegistryExample {
  domain: ValidationOperatorDomain;
  name: string;
  schema: ValidationSchemaCommand;
  purpose: string;
}

export const registryExamples: readonly ValidationRegistryExample[] = [
  {
    domain: "string",
    name: "email",
    schema: ["schema", "string", "--email"],
    purpose: "Built-in string validation",
  },
  {
    domain: "number",
    name: "int",
    schema: ["schema", "number", "--int"],
    purpose: "Built-in numeric validation",
  },
  {
    domain: "formatter",
    name: "trim",
    schema: ["schema", "formatter", "--trim"],
    purpose: "Built-in formatter registration",
  },
  {
    domain: "conversion",
    name: "parse-int",
    schema: ["schema", "conversion", "--parse-int"],
    purpose: "Built-in string-to-number conversion",
  },
  {
    domain: "tuple",
    name: "pair",
    schema: ["schema", "tuple", "--pair"],
    purpose: "Built-in tuple validation",
  },
  {
    domain: "string",
    name: "postal-code",
    schema: ["custom", "postal-code", "--country", "US"],
    purpose: "Custom validator registered in Go",
  },
];
