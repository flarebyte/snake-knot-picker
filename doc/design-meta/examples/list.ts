import type { ListFunction, ValidationError, ValidatorOptions } from "./common";

export declare class MinLength implements ListFunction {
  readonly minLength: number;

  constructor(minLength: number);

  validate(input: readonly unknown[], opts: ValidatorOptions): ValidationError | null;
}

export declare class MaxLength implements ListFunction {
  readonly maxLength: number;

  constructor(maxLength: number);

  validate(input: readonly unknown[], opts: ValidatorOptions): ValidationError | null;
}
