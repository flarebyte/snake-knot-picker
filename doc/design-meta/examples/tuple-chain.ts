import type { TupleValidation, TupleValidationChain } from "./common";

export declare const constrainedTuple: TupleValidation;
export declare const constrainedTupleChain: TupleValidationChain;

// Example chain shape:
// tupleValidations.chain()
//   .pipe(tupleValidations.length(3))
//   .build();
