import type { TupleValidation, TupleValidationChain } from "./common";
import type { ExampleBlock } from "./example";

export declare const constrainedTuple: TupleValidation;
export declare const constrainedTupleChain: TupleValidationChain;

export const tupleChainExamples: readonly ExampleBlock[] = [
  {
    name: "Tuple validation composition",
    code: [
      "tupleValidations.chain()",
      "  .pipe(tupleValidations.of([",
      "    stringValidations.minChars(2),",
      "    stringValidations.matchesFormatter(stringFormatters.trim()),",
      "  ]))",
      "  .build();",
    ],
  },
];
