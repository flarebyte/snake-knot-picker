import { numberValidations } from "./number";
import type { NumberValidation, NumberValidationChain } from "./common";

export const boundedIntChain: NumberValidationChain = numberValidations
  .chain()
  .pipe(numberValidations.min(10))
  .pipe(numberValidations.max(20))
  .pipe(numberValidations.int());

export const boundedInt: NumberValidation = boundedIntChain.build();
