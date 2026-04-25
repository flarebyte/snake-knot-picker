import type { StringValidation } from './common';
import { stringValidations } from './string';

export const booleanString: StringValidation = stringValidations.boolean();
export const colorString: StringValidation = stringValidations.color();
