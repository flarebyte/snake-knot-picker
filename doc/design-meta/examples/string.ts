import type { StringFunction, ValidationError, ValidatorOptions } from "./common";

export interface StringFunctionChain {
  build(): StringFunction;
  pipe(next: StringFunction): StringFunctionChain;
}

export interface StringFunctionFactory {
  alpha(): StringFunction;
  arn(): StringFunction;
  base64(): StringFunction;
  codepointRange(from: string, to: string): StringFunction;
  digit(): StringFunction;
  email(): StringFunction;
  enum(allowedValues: readonly string[]): StringFunction;
  chain(): StringFunctionChain;
  lowercase(): StringFunction;
  maxChars(maxChars: number): StringFunction;
  maxWords(maxWords: number): StringFunction;
  minChars(minChars: number): StringFunction;
  minWords(minWords: number): StringFunction;
  startsWith(prefix: string): StringFunction;
  trim(): StringFunction;
  uppercase(): StringFunction;
  uri(): StringFunction;
  uuid(): StringFunction;
}

export declare const stringFunctions: StringFunctionFactory;

export declare class MinChars implements StringFunction {
  readonly minChars: number;

  constructor(minChars: number);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class MaxChars implements StringFunction {
  readonly maxChars: number;

  constructor(maxChars: number);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class MinWords implements StringFunction {
  readonly minWords: number;

  constructor(minWords: number);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class MaxWords implements StringFunction {
  readonly maxWords: number;

  constructor(maxWords: number);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Enum implements StringFunction {
  readonly allowedValues: readonly string[];

  constructor(allowedValues: readonly string[]);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Digit implements StringFunction {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Alpha implements StringFunction {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Base64 implements StringFunction {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class CodepointRange implements StringFunction {
  readonly from: string;
  readonly to: string;

  constructor(from: string, to: string);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Lowercase implements StringFunction {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Uppercase implements StringFunction {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Trim implements StringFunction {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Uri implements StringFunction {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Arn implements StringFunction {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Email implements StringFunction {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class StartsWith implements StringFunction {
  readonly prefix: string;

  constructor(prefix: string);

  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export declare class Uuid implements StringFunction {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}
