import type {
  ListValidation,
  ListValidationChain,
  ValidationError,
  ValidatorOptions,
} from "./common";

export interface ListValidationFactory {
  chain(): ListValidationChain;
  maxLength(maxLength: number): ListValidation;
  minLength(minLength: number): ListValidation;
}

export declare const listValidations: ListValidationFactory;

export declare class MinLength implements ListValidation {
  readonly minLength: number;

  constructor(minLength: number);

  validate(input: readonly unknown[], opts: ValidatorOptions): ValidationError | null;
}

export declare class MaxLength implements ListValidation {
  readonly maxLength: number;

  constructor(maxLength: number);

  validate(input: readonly unknown[], opts: ValidatorOptions): ValidationError | null;
}
