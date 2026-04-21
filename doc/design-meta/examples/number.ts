import type {
  NumberConversion,
  NumberValidation,
  NumberValidationChain,
  ValidationError,
  ValidatorOptions,
} from "./common";

export interface NumberConversionFactory {
  float(): NumberConversion;
  int(): NumberConversion;
}

export interface NumberValidationFactory {
  chain(): NumberValidationChain;
  float(): NumberValidation;
  int(): NumberValidation;
  max(max: number): NumberValidation;
  min(min: number): NumberValidation;
  multipleOf(factor: number): NumberValidation;
}

export declare const numberConversions: NumberConversionFactory;
export declare const numberValidations: NumberValidationFactory;

export declare class ParseInt implements NumberConversion {
  convert(input: string, opts: ValidatorOptions): number | null;
}

export declare class ParseFloat implements NumberConversion {
  convert(input: string, opts: ValidatorOptions): number | null;
}

export declare class Min implements NumberValidation {
  readonly min: number;

  constructor(min: number);

  validate(input: number, opts: ValidatorOptions): ValidationError | null;
}

export declare class Max implements NumberValidation {
  readonly max: number;

  constructor(max: number);

  validate(input: number, opts: ValidatorOptions): ValidationError | null;
}

export declare class MultipleOf implements NumberValidation {
  readonly factor: number;

  constructor(factor: number);

  validate(input: number, opts: ValidatorOptions): ValidationError | null;
}

export declare class Int implements NumberValidation {
  validate(input: number, opts: ValidatorOptions): ValidationError | null;
}

export declare class Float implements NumberValidation {
  validate(input: number, opts: ValidatorOptions): ValidationError | null;
}
