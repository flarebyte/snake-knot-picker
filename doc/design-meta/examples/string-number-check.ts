import type { NumberValidation, StringValidation } from './common';
import { numberValidations } from './number';
import { stringValidations } from './string';

export const numericString: StringValidation = stringValidations.number(
  numberValidations.int(),
);
export const numericStringValidation: NumberValidation =
  numberValidations.int();
