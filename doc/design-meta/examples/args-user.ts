import type { ArgsCommandSchema, ArgsParseResult } from "./args";
import type { ValidationError } from "./common";

export declare const washStartArgs: readonly string[];
export declare const washStartParseResult: ArgsParseResult;
export declare const washStartValidation: ValidationError | null;
export declare const washStartUserSchema: ArgsCommandSchema;

// Example user validation shape:
// userArgs.parse(
//   ["wash", "start", "normal", "--spin", "1200", "--extra-rinse"],
//   washStartUserSchema,
// )
// userArgs.validate(
//   ["wash", "start", "normal", "--spin", "1200", "--extra-rinse"],
//   washStartUserSchema,
// )
