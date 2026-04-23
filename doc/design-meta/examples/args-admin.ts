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
//   .string("mode", stringValidations.enum(["normal", "delicate", "whites"]))
//   .boolean("extra-rinse")
//   .number("spin", numberValidations.int())
//   .tuple("range", [
//     numberValidations.int(),
//     numberValidations.int(),
//   ])
//   .string("add", stringValidations.alpha(), { repeatable: true })
//   .tuple(
//     "pair",
//     [stringValidations.alpha(), stringValidations.hex()],
//     { repeatable: true },
//   )
//   .string("temp", stringValidations.enum(["cold", "warm", "hot"]))
//   .build();
