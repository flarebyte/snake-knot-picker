import type { StringFormatter, StringValidation } from './common';
import { stringFormatters } from './formatters';
import { stringValidations } from './string';

export const normalizedUserName: StringValidation =
  stringValidations.matchesFormatter(stringFormatters.trim());
export const normalizedUserNameFormatter: StringFormatter =
  stringFormatters.trim();
