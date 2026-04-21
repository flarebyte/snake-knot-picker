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

export interface StringFormatter {
  format(input: string): string;
}

export interface StringFormatterChain {
  build(): StringFormatter;
  pipe(next: StringFormatter): StringFormatterChain;
}

export interface StringFunctionChain {
  build(): StringFunction;
  pipe(next: StringFunction | StringFormatter): StringFunctionChain;
}

export interface NumberFunction {
  validate(input: number, opts: ValidatorOptions): ValidationError | null;
}

export interface NumberFunctionChain {
  build(): NumberFunction;
  pipe(next: NumberFunction): NumberFunctionChain;
}

export interface TupleFunction {
  validate(input: readonly unknown[], opts: ValidatorOptions): ValidationError | null;
}

export interface TupleFunctionChain {
  build(): TupleFunction;
  pipe(next: TupleFunction): TupleFunctionChain;
}

export interface ListFunction {
  validate(input: readonly unknown[], opts: ValidatorOptions): ValidationError | null;
}

export interface ListFunctionChain {
  build(): ListFunction;
  pipe(next: ListFunction): ListFunctionChain;
}

export interface NumberValidator extends NumberFunction {}
