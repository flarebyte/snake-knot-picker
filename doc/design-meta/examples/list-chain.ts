import type { ListFunction } from "./common";
import type { ListFunctionChain } from "./common";

export declare const boundedList: ListFunction;
export declare const boundedListChain: ListFunctionChain;

// Example chain shape:
// listFunctions.chain()
//   .pipe(listFunctions.minLength(1))
//   .pipe(listFunctions.maxLength(5))
//   .build();
