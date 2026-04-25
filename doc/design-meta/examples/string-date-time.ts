import type { StringValidation } from "./common";
import type { ExampleBlock } from "./example";

export declare const dateString: StringValidation;
export declare const dateTimeString: StringValidation;
export declare const timeString: StringValidation;
export declare const durationString: StringValidation;

export const stringDateTimeExamples: readonly ExampleBlock[] = [
  {
    name: "Temporal string validations",
    code: [
      "stringValidations.date()",
      "stringValidations.datetime()",
      "stringValidations.time()",
      "stringValidations.duration()",
    ],
  },
];
