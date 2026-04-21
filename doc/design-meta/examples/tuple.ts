import type {
  StringValidation,
  TupleValidation,
  TupleValidationChain,
  ValidationError,
  ValidatorOptions,
} from "./common";

export interface TupleValidationFactory {
  chain(): TupleValidationChain;
  of(validations: readonly StringValidation[]): TupleValidation;
}

export declare const tupleValidations: TupleValidationFactory;

export declare class TupleOf implements TupleValidation {
  readonly validations: readonly StringValidation[];

  constructor(validations: readonly StringValidation[]);

  validate(input: readonly unknown[], opts: ValidatorOptions): ValidationError | null;
}
