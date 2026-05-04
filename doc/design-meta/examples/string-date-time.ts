import type { StringValidation } from './common';
import { stringValidations } from './string';

export const dateString: StringValidation = stringValidations.date();
export const dateTimeString: StringValidation = stringValidations.datetime();
export const newYorkDateTimeString: StringValidation =
  stringValidations.datetime({
    layout: 'RFC3339',
    allowTimezone: true,
    location: 'America/New_York',
  });
export const timeString: StringValidation = stringValidations.time();
export const durationString: StringValidation = stringValidations.duration();
