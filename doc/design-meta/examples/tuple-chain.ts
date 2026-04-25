import { stringFormatters } from "./formatters";
import { stringValidations } from "./string";
import { tupleValidations } from "./tuple";
import type { TupleValidation, TupleValidationChain } from "./common";

export const constrainedTupleChain: TupleValidationChain = tupleValidations
  .chain()
  .pipe(
    tupleValidations.of([
      stringValidations.minChars(2),
      stringValidations.matchesFormatter(stringFormatters.trim()),
    ]),
  );

export const constrainedTuple: TupleValidation = constrainedTupleChain.build();
