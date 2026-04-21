import type { ArgsCommandSchema } from "./args";
import type { ValidationError } from "./common";

export declare const washStartArgs: readonly string[];
export declare const washStartValidation: ValidationError | null;
export declare const washStartUserSchema: ArgsCommandSchema;

// Example user validation shape:
// userArgs.validate(
//   ["wash", "start", "normal", "--spin", "1200", "--extra-rinse"],
//   washStartUserSchema,
// )
