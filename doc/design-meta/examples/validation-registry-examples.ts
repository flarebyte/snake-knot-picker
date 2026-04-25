import { stringFormatters } from './formatters';
import { numberConversions, numberValidations } from './number';
import { stringValidations } from './string';
import { tupleValidations } from './tuple';
import type { ValidationRegistry } from './validation-registry';
import { validationRegistry } from './validation-registry';

export const stringEmailRegistration: ValidationRegistry =
  validationRegistry.register({
    domain: 'string',
    name: 'email',
    schema: ['schema', 'string', '--email'],
    validation: stringValidations.email(),
  });

export const numberIntRegistration: ValidationRegistry =
  validationRegistry.register({
    domain: 'number',
    name: 'int',
    schema: ['schema', 'number', '--int'],
    validation: numberValidations.int(),
  });

export const formatterTrimRegistration: ValidationRegistry =
  validationRegistry.register({
    domain: 'formatter',
    name: 'trim',
    schema: ['schema', 'formatter', '--trim'],
    formatter: stringFormatters.trim(),
  });

export const conversionParseIntRegistration: ValidationRegistry =
  validationRegistry.register({
    domain: 'conversion',
    name: 'parse-int',
    schema: ['schema', 'conversion', '--parse-int'],
    conversion: numberConversions.int(),
  });

export const tuplePairRegistration: ValidationRegistry =
  validationRegistry.register({
    domain: 'tuple',
    name: 'pair',
    schema: ['schema', 'tuple', '--pair'],
    validation: tupleValidations.of([
      stringValidations.alpha(),
      stringValidations.hex(),
    ]),
  });

export const postalCodeRegistration: ValidationRegistry =
  validationRegistry.register({
    domain: 'string',
    name: 'postal-code',
    schema: ['custom', 'postal-code', '--country', 'US'],
    validation: stringValidations.enum(['US', 'CA']),
  });
