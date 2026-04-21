import type { StringFunction } from "./common";
import type { StringFunctionChain } from "./common";

export declare const normalizedEmail: StringFunction;
export declare const normalizedEmailChain: StringFunctionChain;

// Example chain shape:
// stringFunctions.chain()
//   .pipe(stringFunctions.minChars(10))
//   .pipe(stringFunctions.maxChars(40))
//   .pipe(stringFunctions.email())
//   .build();
