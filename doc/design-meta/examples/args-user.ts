import type { ArgsCommandSchema, ArgsParseResult } from "./args";
import type { ValidationError } from "./common";
import type { ExampleBlock } from "./example";

export declare const washStartArgs: readonly string[];
export declare const washStartParseResult: ArgsParseResult;
export declare const washStartValidation: ValidationError | null;
export declare const washStartUserSchema: ArgsCommandSchema;

export const argsUserExamples: readonly ExampleBlock[] = [
  {
    name: "Parse user args",
    code: [
      'userArgs.parse([',
      '  "wash", "start", "--mode", "normal", "--spin", "1200", "--extra-rinse",',
      "], washStartUserSchema);",
    ],
  },
  {
    name: "Validate user args",
    code: [
      'userArgs.validate([',
      '  "wash", "start", "--mode", "normal", "--spin", "1200", "--extra-rinse",',
      "], washStartUserSchema);",
    ],
  },
];
