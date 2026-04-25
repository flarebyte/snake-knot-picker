import type { NumberConversion } from './common';
import { numberConversions } from './number';

export const parseIntConversion: NumberConversion = numberConversions.int();
export const parseFloatConversion: NumberConversion = numberConversions.float();
