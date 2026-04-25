import type {
  NumberValidation,
  StringFormatter,
  StringValidation,
  StringValidationChain,
  ValidationError,
  ValidatorOptions,
} from './common';

export interface StringValidationFactory {
  arn(): StringValidation;
  base64(): StringValidation;
  alphabetic(): StringValidation;
  blank(): StringValidation;
  boolean(): StringValidation;
  date(): StringValidation;
  datetime(): StringValidation;
  duration(): StringValidation;
  chain(): StringValidationChain;
  codepointRange(from: string, to: string): StringValidation;
  color(): StringValidation;
  digit(): StringValidation;
  email(): StringValidation;
  enum(allowedValues: readonly string[]): StringValidation;
  hexadecimalDigit(): StringValidation;
  lowercase(): StringValidation;
  maxChars(maxChars: number): StringValidation;
  maxWords(maxWords: number): StringValidation;
  matchesFormatter(formatter: StringFormatter): StringValidation;
  minChars(minChars: number): StringValidation;
  minWords(minWords: number): StringValidation;
  number(numberValidation: NumberValidation): StringValidation;
  punctuation(): StringValidation;
  time(): StringValidation;
  startsWith(prefix: string): StringValidation;
  unicodeLetter(): StringValidation;
  unicodeNumber(): StringValidation;
  unicodePunctuation(): StringValidation;
  unicodeSeparator(): StringValidation;
  unicodeSymbol(): StringValidation;
  uri(): StringValidation;
  uppercase(): StringValidation;
  whitespace(): StringValidation;
  uuid(): StringValidation;
}

export declare const stringValidations: StringValidationFactory;

export declare class MinChars implements StringValidation {
  readonly minChars: number;

  constructor(minChars: number);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class MaxChars implements StringValidation {
  readonly maxChars: number;

  constructor(maxChars: number);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class MinWords implements StringValidation {
  readonly minWords: number;

  constructor(minWords: number);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class MaxWords implements StringValidation {
  readonly maxWords: number;

  constructor(maxWords: number);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Enum implements StringValidation {
  readonly allowedValues: readonly string[];

  constructor(allowedValues: readonly string[]);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Digit implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Alphabetic implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Base64 implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Blank implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class HexadecimalDigit implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class CodepointRange implements StringValidation {
  readonly from: string;
  readonly to: string;

  constructor(from: string, to: string);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Bool implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Color implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Lowercase implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class DateString implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class DateTimeString implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Uri implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Arn implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Email implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class StartsWith implements StringValidation {
  readonly prefix: string;

  constructor(prefix: string);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class MatchesFormatter implements StringValidation {
  readonly formatter: StringFormatter;

  constructor(formatter: StringFormatter);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class NumberStringValidation implements StringValidation {
  readonly numberValidation: NumberValidation;

  constructor(numberValidation: NumberValidation);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Punctuation implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class TimeString implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class UnicodeLetter implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class UnicodeNumber implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class UnicodePunctuation implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class UnicodeSeparator implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class UnicodeSymbol implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Uppercase implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Whitespace implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class DurationString implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Uuid implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}
