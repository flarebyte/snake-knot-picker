import type { NumberFunction } from "./common";
import type { NumberFunctionChain } from "./common";

export declare const boundedInt: NumberFunction;
export declare const boundedIntChain: NumberFunctionChain;

// Example chain shape:
// numberFunctions.chain()
//   .pipe(numberFunctions.min(10))
//   .pipe(numberFunctions.max(20))
//   .pipe(numberFunctions.int())
//   .build();
