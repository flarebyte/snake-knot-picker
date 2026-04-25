import type { StringFormatter, StringFormatterChain } from "./common";
import type { ExampleBlock } from "./example";

export declare const normalizedFormatter: StringFormatter;
export declare const normalizedFormatterChain: StringFormatterChain;

export const formatterChainExamples: readonly ExampleBlock[] = [
  {
    name: "Formatter normalization",
    code: [
      "stringFormatters.chain()",
      "  .pipe(stringFormatters.trim())",
      "  .pipe(stringFormatters.lowercase())",
      "  .build();",
    ],
  },
];
