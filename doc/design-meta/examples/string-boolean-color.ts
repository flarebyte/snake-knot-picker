import type { StringValidation } from "./common";
import type { ExampleBlock } from "./example";

export declare const booleanString: StringValidation;
export declare const colorString: StringValidation;

export const stringBooleanColorExamples: readonly ExampleBlock[] = [
  {
    name: "Boolean and color string validations",
    code: [
      "stringValidations.boolean()",
      "stringValidations.color()",
    ],
  },
];
