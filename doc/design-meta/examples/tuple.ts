import type {
  TupleValidation,
  TupleValidationChain,
  ValidationError,
  ValidatorOptions,
} from "./common";

export interface TupleValidationFactory {
  chain(): TupleValidationChain;
  length(length: number): TupleValidation;
}

export declare const tupleValidations: TupleValidationFactory;

export declare class Length implements TupleValidation {
  readonly length: number;

  constructor(length: number);

  validate(input: readonly unknown[], opts: ValidatorOptions): ValidationError | null;
}
