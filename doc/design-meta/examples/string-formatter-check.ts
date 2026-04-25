import type { StringFormatter, StringValidation } from "./common";
import type { ExampleBlock } from "./example";

export declare const normalizedUserName: StringValidation;
export declare const normalizedUserNameFormatter: StringFormatter;

export const stringFormatterExamples: readonly ExampleBlock[] = [
  {
    name: "Match a formatter without changing input",
    code: [
      "stringValidations.matchesFormatter(stringFormatters.trim())",
    ],
  },
];
