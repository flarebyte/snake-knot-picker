import { stringValidations } from "./string";
import type { StringValidation } from "./common";

export const booleanString: StringValidation = stringValidations.boolean();
export const colorString: StringValidation = stringValidations.color();
