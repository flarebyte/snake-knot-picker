export interface ValidationError {
  errorMessageIds: string[];
  field?: string;
}

export interface ValidatorOptions {
  field?: string;
}

export interface StringValidation {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export interface StringFormatter {
  format(input: string): string;
}

export interface StringFormatterChain {
  build(): StringFormatter;
  pipe(next: StringFormatter): StringFormatterChain;
}

export interface StringValidationChain {
  build(): StringValidation;
  pipe(next: StringValidation | StringFormatter): StringValidationChain;
}

export interface NumberValidation {
  validate(input: number, opts: ValidatorOptions): ValidationError | null;
}

export interface NumberValidationChain {
  build(): NumberValidation;
  pipe(next: NumberValidation): NumberValidationChain;
}

export interface NumberConversion {
  convert(input: string, opts: ValidatorOptions): number | null;
}

export interface TupleValidation {
  validate(
    input: readonly unknown[],
    opts: ValidatorOptions,
  ): ValidationError | null;
}

export interface TupleValidationChain {
  build(): TupleValidation;
  pipe(next: TupleValidation): TupleValidationChain;
}

export interface ListValidation {
  validate(
    input: readonly unknown[],
    opts: ValidatorOptions,
  ): ValidationError | null;
}

export interface ListValidationChain {
  build(): ListValidation;
  pipe(next: ListValidation): ListValidationChain;
}

export type ListItemValidation = StringValidation | TupleValidation;

export interface NumberValidator extends NumberValidation {}
