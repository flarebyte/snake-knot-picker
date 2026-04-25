import type { ArgsCommandSchema } from './args';
import { adminArgs } from './args';
import { numberValidations } from './number';
import { stringValidations } from './string';

export const schemaString: ArgsCommandSchema = adminArgs
  .command(['schema', 'string'])
  .adminOnly()
  .string('min-chars', stringValidations.minChars(10))
  .string('max-chars', stringValidations.maxChars(20))
  .string('enum', stringValidations.enum(['green', 'orange', 'red']), [
    ['schema', 'string', '--enum', 'green,orange,red', '--required'],
  ])
  .build();

export const washStartSchema: ArgsCommandSchema = adminArgs
  .command(['wash', 'start'])
  .adminOnly()
  .string('mode', stringValidations.enum(['normal', 'delicate', 'whites']), [
    ['schema', 'string', '--enum', 'normal,delicate,whites', '--required'],
  ])
  .boolean('extra-rinse')
  .number('spin', numberValidations.int(), [
    ['schema', 'number', '--int', '--required'],
  ])
  .tuple(
    'range',
    [numberValidations.int(), numberValidations.int()],
    [
      [
        'schema',
        'number',
        '--tuple',
        '0',
        '--int',
        '--tuple',
        '1',
        '--int',
        '--required',
      ],
    ],
  )
  .string('add', stringValidations.alphabetic(), [
    ['schema', 'repeatable', '--min-length', '1', '--max-length', '5'],
  ])
  .tuple(
    'pair',
    [stringValidations.alphabetic(), stringValidations.hexa()],
    [['schema', 'repeatable', '--min-length', '1', '--max-length', '5']],
  )
  .number('dose', numberValidations.int(), [
    ['schema', 'repeatable', '--min-length', '1', '--max-length', '3'],
  ])
  .string('temp', stringValidations.enum(['cold', 'warm', 'hot']), [
    ['schema', 'string', '--enum', 'cold,warm,hot', '--required'],
  ])
  .build();
