import { stringValidations } from "./string";
import type { StringValidation } from "./common";

export const dateString: StringValidation = stringValidations.date();
export const dateTimeString: StringValidation = stringValidations.datetime();
export const timeString: StringValidation = stringValidations.time();
export const durationString: StringValidation = stringValidations.duration();
