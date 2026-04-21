import type { StringValidation, StringValidationChain } from "./common";

export declare const normalizedEmail: StringValidation;
export declare const normalizedEmailChain: StringValidationChain;

// Example chain shape:
// stringValidations.chain()
//   .pipe(stringValidations.matchesFormatter(stringFormatters.trim()))
//   .pipe(stringValidations.number(numberValidations.int()))
//   .pipe(stringValidations.minChars(10))
//   .pipe(stringValidations.maxChars(40))
//   .pipe(stringValidations.email())
//   .build();
