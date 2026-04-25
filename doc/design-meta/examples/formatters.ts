import type { StringFormatter, StringFormatterChain } from './common';

export interface StringFormatterFactory {
  chain(): StringFormatterChain;
  lowercase(): StringFormatter;
  trim(): StringFormatter;
  uppercase(): StringFormatter;
}

export declare const stringFormatters: StringFormatterFactory;

export declare class Trim implements StringFormatter {
  format(input: string): string;
}

export declare class Lowercase implements StringFormatter {
  format(input: string): string;
}

export declare class Uppercase implements StringFormatter {
  format(input: string): string;
}
