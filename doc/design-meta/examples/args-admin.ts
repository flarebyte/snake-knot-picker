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
//   .string("add", stringValidations.alpha(), { minLength: 1, maxLength: 5 })
//   .tuple(
//     "pair",
//     [stringValidations.alpha(), stringValidations.hex()],
//     { minLength: 1, maxLength: 5 },
//   )
//   .number(
//     "dose",
//     numberValidations.int(),
//     { minLength: 1, maxLength: 3 },
//   )
//   .string("temp", stringValidations.enum(["cold", "warm", "hot"]))
//   .build();
