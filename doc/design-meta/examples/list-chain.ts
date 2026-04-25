import { listValidations } from "./list";
import { stringValidations } from "./string";
import type { ListValidation, ListValidationChain } from "./common";

export const boundedListChain: ListValidationChain = listValidations
  .chain()
  .pipe(listValidations.of(stringValidations.email()))
  .pipe(listValidations.minLength(1))
  .pipe(listValidations.maxLength(5));

export const boundedList: ListValidation = boundedListChain.build();
