import type { ValidationRegistry } from "./validation-registry";

export declare const registryExample: ValidationRegistry;

// Example registrations:
// validationRegistry.register({
//   domain: "string",
//   name: "email",
//   schema: ["schema", "string", "--email"],
//   validation: stringValidations.email(),
// })
//
// validationRegistry.register({
//   domain: "number",
//   name: "int",
//   schema: ["schema", "number", "--int"],
//   validation: numberValidations.int(),
// })
//
// validationRegistry.register({
//   domain: "formatter",
//   name: "trim",
//   schema: ["schema", "formatter", "--trim"],
//   formatter: stringFormatters.trim(),
// })
//
// validationRegistry.register({
//   domain: "conversion",
//   name: "parse-int",
//   schema: ["schema", "conversion", "--parse-int"],
//   conversion: numberConversions.int(),
// })
//
// validationRegistry.register({
//   domain: "tuple",
//   name: "pair",
//   schema: ["schema", "tuple", "--pair"],
//   validation: tupleValidations.of([
//     stringValidations.alpha(),
//     stringValidations.hex(),
//   ]),
// })
//
// validationRegistry.register({
//   domain: "string",
//   name: "postal-code",
//   schema: ["custom", "postal-code", "--country", "US"],
//   validation: stringValidations.enum(["US", "CA"]),
// })
