import type { TupleValidation, TupleValidationChain } from "./common";

export declare const constrainedTuple: TupleValidation;
export declare const constrainedTupleChain: TupleValidationChain;

// Example chain shape:
// tupleValidations.chain()
//   .pipe(tupleValidations.of([
//     stringValidations.minChars(2),
//     stringValidations.matchesFormatter(stringFormatters.trim()),
//   ]))
//   .build();
