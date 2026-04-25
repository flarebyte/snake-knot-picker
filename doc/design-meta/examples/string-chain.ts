import type { StringValidation, StringValidationChain } from './common';
import { stringFormatters } from './formatters';
import { numberValidations } from './number';
import { stringValidations } from './string';

export const normalizedEmailChain: StringValidationChain = stringValidations
  .chain()
  .pipe(stringValidations.matchesFormatter(stringFormatters.trim()))
  .pipe(stringValidations.number(numberValidations.int()))
  .pipe(stringValidations.minChars(10))
  .pipe(stringValidations.maxChars(40))
  .pipe(stringValidations.email());

export const normalizedEmail: StringValidation = normalizedEmailChain.build();
