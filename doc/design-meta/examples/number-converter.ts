import type { NumberConversion } from "./common";
import type { ExampleBlock } from "./example";

export declare const parseIntConversion: NumberConversion;
export declare const parseFloatConversion: NumberConversion;

export const numberConversionExamples: readonly ExampleBlock[] = [
  {
    name: "String to number conversions",
    code: [
      "numberConversions.int()",
      "numberConversions.float()",
    ],
  },
];
