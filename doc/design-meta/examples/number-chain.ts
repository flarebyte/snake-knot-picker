import type { NumberValidation, NumberValidationChain } from './common';
import { numberValidations } from './number';

export const boundedIntChain: NumberValidationChain = numberValidations
  .chain()
  .pipe(numberValidations.min(10))
  .pipe(numberValidations.max(20))
  .pipe(numberValidations.int());

export const boundedInt: NumberValidation = boundedIntChain.build();
