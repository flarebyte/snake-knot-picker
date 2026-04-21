import type { ArgsCommandSchema } from "./args";

export declare const schemaString: ArgsCommandSchema;
export declare const washStartSchema: ArgsCommandSchema;

// Example trusted schema shapes:
// adminArgs.command(["schema", "string"])
//   .adminOnly()
//   .string("min-chars", stringValidations.minChars(10))
//   .string("max-chars", stringValidations.maxChars(20))
//   .string("enum", stringValidations.enum(["green", "orange", "red"]))
//   .build();
//
// adminArgs.command(["wash", "start"])
//   .adminOnly()
//   .positional(stringValidations.enum(["normal", "delicate", "whites"]))
//   .boolean("extra-rinse")
//   .string("temp", stringValidations.enum(["cold", "warm", "hot"]))
//   .tuple("range", tupleValidations.of([
//     stringValidations.number(numberValidations.int()),
//     stringValidations.number(numberValidations.int()),
//   ]))
//   .list("add", listValidations.of(stringValidations.alpha()))
//   .build();
