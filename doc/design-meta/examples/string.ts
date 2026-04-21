import type {
  StringFormatter,
  StringValidation,
  StringValidationChain,
  NumberValidation,
  ValidationError,
  ValidatorOptions,
} from "./common";

export interface StringValidationFactory {
  alpha(): StringValidation;
  arn(): StringValidation;
  base64(): StringValidation;
  chain(): StringValidationChain;
  codepointRange(from: string, to: string): StringValidation;
  color(): StringValidation;
  digit(): StringValidation;
  email(): StringValidation;
  enum(allowedValues: readonly string[]): StringValidation;
  boolean(): StringValidation;
  maxChars(maxChars: number): StringValidation;
  maxWords(maxWords: number): StringValidation;
  matchesFormatter(formatter: StringFormatter): StringValidation;
  minChars(minChars: number): StringValidation;
  minWords(minWords: number): StringValidation;
  number(numberValidation: NumberValidation): StringValidation;
  startsWith(prefix: string): StringValidation;
  uri(): StringValidation;
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

export declare class Alpha implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Base64 implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class CodepointRange implements StringValidation {
  readonly from: string;
  readonly to: string;

  constructor(from: string, to: string);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Boolean implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Color implements StringValidation {
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

export declare class Uuid implements StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}
