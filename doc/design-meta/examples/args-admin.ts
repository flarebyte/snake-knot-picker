import { adminArgs } from "./args";
import { numberValidations } from "./number";
import { stringValidations } from "./string";
import type { ArgsCommandSchema } from "./args";

export const schemaString: ArgsCommandSchema = adminArgs
  .command(["schema", "string"])
  .adminOnly()
  .string("min-chars", stringValidations.minChars(10))
  .string("max-chars", stringValidations.maxChars(20))
  .string("enum", stringValidations.enum(["green", "orange", "red"]))
  .build();

export const washStartSchema: ArgsCommandSchema = adminArgs
  .command(["wash", "start"])
  .adminOnly()
  .string("mode", stringValidations.enum(["normal", "delicate", "whites"]))
  .boolean("extra-rinse")
  .number("spin", numberValidations.int())
  .tuple("range", [numberValidations.int(), numberValidations.int()])
  .string("add", stringValidations.alpha(), [
    ["schema", "repeatable", "--min-length", "1", "--max-length", "5"],
  ])
  .tuple("pair", [stringValidations.alpha(), stringValidations.hex()], [
    ["schema", "repeatable", "--min-length", "1", "--max-length", "5"],
  ])
  .number("dose", numberValidations.int(), [
    ["schema", "repeatable", "--min-length", "1", "--max-length", "3"],
  ])
  .string("temp", stringValidations.enum(["cold", "warm", "hot"]))
  .build();
