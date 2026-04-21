import type { NumberValidation, NumberValidationChain } from "./common";

export declare const boundedInt: NumberValidation;
export declare const boundedIntChain: NumberValidationChain;

// Example chain shape:
// numberValidations.chain()
//   .pipe(numberValidations.min(10))
//   .pipe(numberValidations.max(20))
//   .pipe(numberValidations.int())
//   .build();
