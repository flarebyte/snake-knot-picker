import type { StringValidation } from './common';
import { stringValidations } from './string';

export const digitString: StringValidation = stringValidations.digit();
export const whitespaceString: StringValidation =
  stringValidations.whitespace();
export const alphabeticString: StringValidation =
  stringValidations.alphabetic();
export const lowercaseString: StringValidation = stringValidations.lowercase();
export const uppercaseString: StringValidation = stringValidations.uppercase();
export const punctuationString: StringValidation =
  stringValidations.punctuation();
export const hexaString: StringValidation = stringValidations.hexa();
export const blankString: StringValidation = stringValidations.blank();
export const unicodeLetterString: StringValidation =
  stringValidations.unicodeLetter();
export const unicodeNumberString: StringValidation =
  stringValidations.unicodeNumber();
export const unicodePunctuationString: StringValidation =
  stringValidations.unicodePunctuation();
export const unicodeSymbolString: StringValidation =
  stringValidations.unicodeSymbol();
export const unicodeSeparatorString: StringValidation =
  stringValidations.unicodeSeparator();
