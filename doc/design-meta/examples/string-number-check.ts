import type { NumberValidation, StringValidation } from "./common";
import type { ExampleBlock } from "./example";

export declare const numericString: StringValidation;
export declare const numericStringValidation: NumberValidation;

export const stringNumberExamples: readonly ExampleBlock[] = [
  {
    name: "String validated as number",
    code: [
      "stringValidations.number(numberValidations.int())",
      "stringValidations.number(numberValidations.float())",
    ],
  },
];
