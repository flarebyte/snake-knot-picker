import type { StringValidation, StringValidationChain } from "./common";
import type { ExampleBlock } from "./example";

export declare const normalizedEmail: StringValidation;
export declare const normalizedEmailChain: StringValidationChain;

export const stringChainExamples: readonly ExampleBlock[] = [
  {
    name: "String normalization and validation",
    code: [
      "stringValidations.chain()",
      "  .pipe(stringValidations.matchesFormatter(stringFormatters.trim()))",
      "  .pipe(stringValidations.number(numberValidations.int()))",
      "  .pipe(stringValidations.minChars(10))",
      "  .pipe(stringValidations.maxChars(40))",
      "  .pipe(stringValidations.email())",
      "  .build();",
    ],
  },
];
