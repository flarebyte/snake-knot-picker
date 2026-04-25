import type { NumberValidation, NumberValidationChain } from "./common";
import type { ExampleBlock } from "./example";

export declare const boundedInt: NumberValidation;
export declare const boundedIntChain: NumberValidationChain;

export const numberChainExamples: readonly ExampleBlock[] = [
  {
    name: "Bounded integer validation",
    code: [
      "numberValidations.chain()",
      "  .pipe(numberValidations.min(10))",
      "  .pipe(numberValidations.max(20))",
      "  .pipe(numberValidations.int())",
      "  .build();",
    ],
  },
];
