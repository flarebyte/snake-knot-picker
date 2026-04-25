import { stringFormatters } from "./formatters";
import type { StringFormatter, StringFormatterChain } from "./common";

export const normalizedFormatterChain: StringFormatterChain = stringFormatters
  .chain()
  .pipe(stringFormatters.trim())
  .pipe(stringFormatters.lowercase());

export const normalizedFormatter: StringFormatter = normalizedFormatterChain.build();
