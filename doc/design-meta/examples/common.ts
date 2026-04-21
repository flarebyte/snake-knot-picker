export interface ValidationError {
  errorMessageIds: string[];
  field?: string;
}

export interface ValidatorOptions {
  field?: string;
}

export interface StringFunction {
  validate(input: string, opts: ValidatorOptions): ValidationError | null;
}

export interface NumberFunction {
  validate(input: number, opts: ValidatorOptions): ValidationError | null;
}

export interface NumberValidator extends NumberFunction {}
