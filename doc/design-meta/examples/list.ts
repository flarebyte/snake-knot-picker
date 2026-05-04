import type {
  ListItemValidation,
  ListValidation,
  ListValidationChain,
  ValidationError,
  ValidatorOptions,
} from './common';

export interface ListValidationFactory {
  chain(): ListValidationChain;
  of(itemValidation: ListItemValidation): ListValidation;
  minLength(minLength: number): ListValidation;
  maxLength(maxLength: number): ListValidation;
}

export declare const listValidations: ListValidationFactory;

export declare class ListOf implements ListValidation {
  readonly itemValidation: ListItemValidation;

  constructor(itemValidation: ListItemValidation);

  validate(
    input: readonly unknown[],
    opts: ValidatorOptions,
  ): ValidationError | null;
}

export declare class MinLength implements ListValidation {
  readonly minLength: number;

  constructor(minLength: number);

  validate(
    input: readonly unknown[],
    opts: ValidatorOptions,
  ): ValidationError | null;
}

export declare class MaxLength implements ListValidation {
  readonly maxLength: number;

  constructor(maxLength: number);

  validate(
    input: readonly unknown[],
    opts: ValidatorOptions,
  ): ValidationError | null;
}
