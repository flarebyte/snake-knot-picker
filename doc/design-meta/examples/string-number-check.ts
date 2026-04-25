import { numberValidations } from "./number";
import { stringValidations } from "./string";
import type { NumberValidation, StringValidation } from "./common";

export const numericString: StringValidation = stringValidations.number(numberValidations.int());
export const numericStringValidation: NumberValidation = numberValidations.int();
