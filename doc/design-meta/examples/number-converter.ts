import { numberConversions } from "./number";
import type { NumberConversion } from "./common";

export const parseIntConversion: NumberConversion = numberConversions.int();
export const parseFloatConversion: NumberConversion = numberConversions.float();
