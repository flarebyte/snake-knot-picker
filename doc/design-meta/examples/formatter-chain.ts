import type { StringFormatter, StringFormatterChain } from "./common";

export declare const normalizedFormatter: StringFormatter;
export declare const normalizedFormatterChain: StringFormatterChain;

// Example chain shape:
// stringFormatters.chain()
//   .pipe(stringFormatters.trim())
//   .pipe(stringFormatters.lowercase())
//   .build();
