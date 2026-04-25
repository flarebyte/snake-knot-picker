import type { StringValidation } from './common';
import { stringValidations } from './string';

export const dateString: StringValidation = stringValidations.date();
export const dateTimeString: StringValidation = stringValidations.datetime();
export const timeString: StringValidation = stringValidations.time();
export const durationString: StringValidation = stringValidations.duration();
