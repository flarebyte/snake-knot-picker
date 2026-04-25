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
    ['schema', 'required'],
  ])
  .build();

export const washStartSchema: ArgsCommandSchema = adminArgs
  .command(['wash', 'start'])
  .adminOnly()
  .string('mode', stringValidations.enum(['normal', 'delicate', 'whites']), [
    ['schema', 'required'],
  ])
  .boolean('extra-rinse')
  .number('spin', numberValidations.int(), [['schema', 'required']])
  .tuple(
    'range',
    [numberValidations.int(), numberValidations.int()],
    [['schema', 'required']],
  )
  .string('add', stringValidations.alpha(), [
    ['schema', 'repeatable', '--min-length', '1', '--max-length', '5'],
  ])
  .tuple(
    'pair',
    [stringValidations.alpha(), stringValidations.hex()],
    [['schema', 'repeatable', '--min-length', '1', '--max-length', '5']],
  )
  .number('dose', numberValidations.int(), [
    ['schema', 'repeatable', '--min-length', '1', '--max-length', '3'],
  ])
  .string('temp', stringValidations.enum(['cold', 'warm', 'hot']), [
    ['schema', 'required'],
  ])
  .build();
