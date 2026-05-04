import type { StringValidation } from './common';
import { stringValidations } from './string';

export const dateString: StringValidation = stringValidations.date();
export const isoDateString: StringValidation = stringValidations.date({
  layout: 'ISO8601',
});
export const dateTimeString: StringValidation = stringValidations.datetime();
export const newYorkDateTimeString: StringValidation =
  stringValidations.datetime({
    layout: 'RFC3339',
    allowTimezone: true,
    location: 'America/New_York',
  });
export const timeString: StringValidation = stringValidations.time();
export const hourMinuteTimeString: StringValidation = stringValidations.time({
  layout: 'HHMM',
});
export const fractionalTimeString: StringValidation = stringValidations.time({
  layout: 'HHMMSS',
  allowFraction: true,
});
export const durationString: StringValidation = stringValidations.duration();
export const boundedDurationString: StringValidation =
  stringValidations.duration({
    minDuration: '5m',
    maxDuration: '2h',
  });
