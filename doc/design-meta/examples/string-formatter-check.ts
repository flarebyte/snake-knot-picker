import { stringFormatters } from "./formatters";
import { stringValidations } from "./string";
import type { StringFormatter, StringValidation } from "./common";

export const normalizedUserName: StringValidation = stringValidations.matchesFormatter(
  stringFormatters.trim(),
);
export const normalizedUserNameFormatter: StringFormatter = stringFormatters.trim();
