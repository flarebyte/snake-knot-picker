import type { ListValidation, ListValidationChain } from "./common";

export declare const boundedList: ListValidation;
export declare const boundedListChain: ListValidationChain;

// Example chain shape:
// listValidations.chain()
//   .pipe(listValidations.minLength(1))
//   .pipe(listValidations.maxLength(5))
//   .build();
