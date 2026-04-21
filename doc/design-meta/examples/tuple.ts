import type { TupleFunction, TupleFunctionChain, ValidationError, ValidatorOptions } from "./common";

export interface TupleFunctionFactory {
  chain(): TupleFunctionChain;
  length(length: number): TupleFunction;
}

export declare const tupleFunctions: TupleFunctionFactory;

export declare class Length implements TupleFunction {
  readonly length: number;

  constructor(length: number);

  validate(input: readonly unknown[], opts: ValidatorOptions): ValidationError | null;
}
