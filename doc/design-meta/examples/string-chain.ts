import type { StringFunction } from "./common";
import type { StringFunctionChain } from "./string";

export declare const normalizedEmail: StringFunction;
export declare const normalizedEmailChain: StringFunctionChain;

// Example chain shape:
// stringFunctions.chain()
//   .pipe(stringFunctions.trim())
//   .pipe(stringFunctions.lowercase())
//   .pipe(stringFunctions.email())
//   .build();
