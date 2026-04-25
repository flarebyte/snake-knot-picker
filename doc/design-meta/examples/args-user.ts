import type { ArgsCommandSchema, ArgsParseResult } from './args';
import { userArgs } from './args';
import type { ValidationError } from './common';
import { numberValidations } from './number';
import { stringValidations } from './string';

export const washStartArgs = [
  'wash',
  'start',
  '--mode',
  'normal',
  '--spin',
  '1200',
  '--extra-rinse',
] as const;

export const washStartUserSchema: ArgsCommandSchema = {
  commandPath: ['wash', 'start'],
  adminOnly: true,
  flags: [
    {
      kind: 'string',
      name: 'mode',
      schema: [
        'schema',
        'string',
        '--enum',
        'normal,delicate,whites',
        '--required',
      ],
      validation: stringValidations.enum(['normal', 'delicate', 'whites']),
    },
    {
      kind: 'boolean',
      name: 'extra-rinse',
      schema: ['schema', 'boolean'],
    },
    {
      kind: 'number',
      name: 'spin',
      schema: ['schema', 'number', '--int', '--required'],
      validation: numberValidations.int(),
    },
  ],
};

export const washStartParseResult: ArgsParseResult = userArgs.parse(
  washStartArgs,
  washStartUserSchema,
);

export const washStartValidation: ValidationError | null = userArgs.validate(
  washStartArgs,
  washStartUserSchema,
);
