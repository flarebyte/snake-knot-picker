import type { TupleFunction } from "./common";
import type { TupleFunctionChain } from "./common";

export declare const constrainedTuple: TupleFunction;
export declare const constrainedTupleChain: TupleFunctionChain;

// Example chain shape:
// tupleFunctions.chain()
//   .pipe(tupleFunctions.length(3))
//   .build();
