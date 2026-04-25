import { numberValidations } from "./number";
import { stringFormatters } from "./formatters";
import { stringValidations } from "./string";
import type { StringValidation, StringValidationChain } from "./common";

export const normalizedEmailChain: StringValidationChain = stringValidations
  .chain()
  .pipe(stringValidations.matchesFormatter(stringFormatters.trim()))
  .pipe(stringValidations.number(numberValidations.int()))
  .pipe(stringValidations.minChars(10))
  .pipe(stringValidations.maxChars(40))
  .pipe(stringValidations.email());

export const normalizedEmail: StringValidation = normalizedEmailChain.build();
