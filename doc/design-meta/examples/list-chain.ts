import type { ListValidation, ListValidationChain } from "./common";
import type { ExampleBlock } from "./example";

export declare const boundedList: ListValidation;
export declare const boundedListChain: ListValidationChain;

export const listChainExamples: readonly ExampleBlock[] = [
  {
    name: "Repeatable list validation",
    code: [
      "listValidations.chain()",
      "  .pipe(listValidations.of(stringValidations.email()))",
      "  .pipe(listValidations.minLength(1))",
      "  .pipe(listValidations.maxLength(5))",
      "  .build();",
    ],
  },
];
