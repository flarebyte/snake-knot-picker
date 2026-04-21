import type { NumberFunction, NumberFunctionChain, ValidationError, ValidatorOptions } from "./common";

export interface NumberFunctionFactory {
  chain(): NumberFunctionChain;
  float(): NumberFunction;
  int(): NumberFunction;
  max(max: number): NumberFunction;
  min(min: number): NumberFunction;
  multipleOf(factor: number): NumberFunction;
}

export declare const numberFunctions: NumberFunctionFactory;

export declare class Min implements NumberFunction {
  readonly min: number;

  constructor(min: number);

  validate(input: number, opts: ValidatorOptions): ValidationError | null;
}

export declare class Max implements NumberFunction {
  readonly max: number;

  constructor(max: number);

  validate(input: number, opts: ValidatorOptions): ValidationError | null;
}

export declare class MultipleOf implements NumberFunction {
  readonly factor: number;

  constructor(factor: number);

  validate(input: number, opts: ValidatorOptions): ValidationError | null;
}

export declare class Int implements NumberFunction {
  validate(input: number, opts: ValidatorOptions): ValidationError | null;
}

export declare class Float implements NumberFunction {
  validate(input: number, opts: ValidatorOptions): ValidationError | null;
}
