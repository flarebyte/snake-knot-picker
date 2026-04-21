import type { TupleFunction, ValidationError, ValidatorOptions } from "./common";

export declare class Length implements TupleFunction {
  readonly length: number;

  constructor(length: number);

  validate(input: readonly unknown[], opts: ValidatorOptions): ValidationError | null;
}
