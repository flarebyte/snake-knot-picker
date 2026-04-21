import type { NumberValidation, StringValidation } from "./common";

export declare const numericString: StringValidation;
export declare const numericStringValidation: NumberValidation;

// Example validation shape:
// stringValidations.number(numberValidations.int())
// stringValidations.number(numberValidations.float())
