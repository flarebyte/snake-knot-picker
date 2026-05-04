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
  arabic(): StringValidation;
  bengali(): StringValidation;
  boolean(): StringValidation;
  cyrillic(): StringValidation;
  date(): StringValidation;
  datetime(): StringValidation;
  duration(): StringValidation;
  chain(): StringValidationChain;
  codepointRange(from: string, to: string): StringValidation;
  color(): StringValidation;
  digit(): StringValidation;
  devanagari(): StringValidation;
  ethiopic(): StringValidation;
  email(): StringValidation;
  enum(
    allowedValues: readonly string[],
    options?: EnumOptions,
  ): StringValidation;
  hexa(): StringValidation;
  gurmukhi(): StringValidation;
  han(): StringValidation;
  hangul(): StringValidation;
  hiragana(): StringValidation;
  javanese(): StringValidation;
  gujarati(): StringValidation;
  kannada(): StringValidation;
  latin(): StringValidation;
  katakana(): StringValidation;
  lowercase(): StringValidation;
  maxChars(maxChars: number): StringValidation;
  maxWords(maxWords: number): StringValidation;
  matchesFormatter(formatter: StringFormatter): StringValidation;
  minChars(minChars: number): StringValidation;
  minWords(minWords: number): StringValidation;
  number(numberValidation: NumberValidation): StringValidation;
  punctuation(): StringValidation;
  tamil(): StringValidation;
  telugu(): StringValidation;
  thai(): StringValidation;
  time(): StringValidation;
  startsWith(prefix: string): StringValidation;
  unicodeLetter(): StringValidation;
  unicodeNumber(): StringValidation;
  unicodePunctuation(): StringValidation;
  unicodeSeparator(): StringValidation;
  unicodeSymbol(): StringValidation;
  uri(options?: UriOptions): StringValidation;
  uppercase(): StringValidation;
  whitespace(): StringValidation;
  uuid(): StringValidation;
}

export interface EnumOptions {
  separator?: string;
}

export interface UriOptions {
  scheme?: 'http' | 'https';
  secure?: boolean;
  allowPort?: boolean;
  allowUserInfo?: boolean;
  allowFragment?: boolean;
  allowQuery?: boolean;
  allowDomains?: readonly string[];
  allowIp?: boolean;
  allowOpaque?: boolean;
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
  readonly separator?: string;

  constructor(allowedValues: readonly string[], options?: EnumOptions);

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

export declare class Arabic implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Bengali implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Hexa implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Cyrillic implements StringValidation {
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

export declare class Devanagari implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Ethiopic implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Lowercase implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Gurmukhi implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Han implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Hangul implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Hiragana implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Javanese implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Gujarati implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Kannada implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Latin implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Katakana implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class DateString implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class DateTimeString implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Uri implements StringValidation {
  readonly options?: UriOptions;

  constructor(options?: UriOptions);

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

export declare class Tamil implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Telugu implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Thai implements StringValidation {
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
