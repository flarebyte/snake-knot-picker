package flyb

source: "snake-knot-picker-design"
name:   "snake-knot-picker-design"
modules: ["core"]

reports: [{
	title:       "Snake Knot Picker Design Specification"
	filepath:    "../design/snake-knot-picker.md"
	description: "Canonical specification generated from `doc/design-meta/examples` artefacts."
	sections: [{
		title:       "01 Overview"
		description: "Purpose, scope, and design goals."
		sections: [{
			title:       "01 Project Intent"
			description: "What the library optimizes for."
			notes:       ["snk.intent", "snk.goals"]
		}]
	}, {
		title:       "02 Schema Language"
		description: "Authoring and transport forms for schema commands."
		sections: [{
			title:       "01 Command JSON Shape"
			description: "JSON-compatible command representation."
			notes:       ["snk.schema.json"]
		}, {
			title:       "02 CLI Arguments"
			description: "Command-line invocation examples."
			notes:       ["snk.schema.arguments"]
		}, {
			title:       "03 Admin and User API Types"
			description: "TypeScript examples for builders and runtime parsing."
			notes:       ["snk.schema.args", "snk.schema.args.admin", "snk.schema.args.user"]
		}]
	}, {
		title:       "03 Validation Domains"
		description: "Core validation capabilities across string, number, tuple, and list domains."
		sections: [{
			title:       "01 Validation Capability Matrix"
			description: "Feature inventory grouped by domain."
			notes:       ["snk.validation.matrix"]
		}, {
			title:       "02 String Validation Examples"
			description: "String-focused validators and composition patterns."
			notes: [
				"snk.validation.string",
				"snk.validation.string.enum.csv",
				"snk.validation.string.chain",
				"snk.validation.string.number",
				"snk.validation.string.datetime",
				"snk.validation.string.date.csv",
				"snk.validation.string.datetime.csv",
				"snk.validation.string.time.csv",
				"snk.validation.string.duration.csv",
				"snk.validation.string.email.csv",
				"snk.validation.string.color.csv",
				"snk.validation.string.url.csv",
				"snk.validation.string.arn.csv",
				"snk.validation.string.formatter-check",
				"snk.validation.string.boolean-color",
				"snk.validation.string.classes.csv",
				"snk.validation.string.classes.ts",
				"snk.validation.string.languages.csv",
				"snk.validation.string.languages.ts",
			]
		}, {
			title:       "03 Number Validation Examples"
			description: "Numeric validators, conversion, and chaining."
			notes:       ["snk.validation.number", "snk.validation.number.chain", "snk.validation.number.converter"]
		}, {
			title:       "04 Tuple and List Examples"
			description: "Collection and positional validation patterns."
			notes:       ["snk.validation.tuple", "snk.validation.tuple.chain", "snk.validation.list", "snk.validation.list.chain"]
		}, {
			title:       "05 Formatter Examples"
			description: "Formatter primitives and composition."
			notes:       ["snk.formatters.csv", "snk.formatters.ts", "snk.formatters.chain"]
		}, {
			title:       "06 Registry and Extensibility"
			description: "Operator registration, custom validators, and collision handling."
			notes:       ["snk.registry.api", "snk.registry.examples", "snk.validation.csv"]
		}]
	}, {
		title:       "04 Use Cases and Limits"
		description: "Practical command-level scenarios and policy limits."
		sections: [{
			title:       "01 Use Case Table"
			description: "Admin and user scenarios captured in CSV."
			notes:       ["snk.usecases"]
		}, {
			title:       "02 Cobra Limit Notes"
			description: "Command-line framework constraints impacting schema design."
			notes:       ["snk.cobra.limits"]
		}]
	}, {
		title:       "05 Validation Flows"
		description: "Ordered flow specifications for admin authoring and user runtime validation."
		sections: [{
			title:       "01 Flow Scope"
			description: "Flow intent and boundaries."
			notes:       ["flow.intent", "flow.boundaries"]
		}, {
			title:       "02 Admin Authoring Flow"
			description: "Trusted admin flow from use case to exported schema."
			notes: [
				"flow.admin.capture-usecase",
				"flow.admin.compose-schema",
				"flow.admin.register-operator",
				"flow.admin.export-schema",
			]
		}, {
			title:       "03 Admin Flow Graph"
			description: "Transition graph for admin steps."
			arguments: [
				"graph-subject-label=admin-flow",
				"graph-edge-label=next_step",
				"graph-start-node=flow.admin.capture-usecase",
				"graph-renderer=markdown-text",
			]
		}, {
			title:       "04 User Validation Flow"
			description: "Untrusted user argv flow from parse to validation result."
			notes: [
				"flow.user.receive-argv",
				"flow.user.parse-argv",
				"flow.user.resolve-operators",
				"flow.user.validate-values",
				"flow.user.return-result",
			]
		}, {
			title:       "05 User Flow Graph"
			description: "Transition graph for user validation steps."
			arguments: [
				"graph-subject-label=user-flow",
				"graph-edge-label=next_step",
				"graph-start-node=flow.user.receive-argv",
				"graph-renderer=markdown-text",
			]
		}]
	}]
}]

notes: [
	{
		name:  "snk.intent"
		title: "Validation Library Purpose"
		markdown: """
`snake-knot-picker` provides strict validation with a compact schema language built from CLI-style argument tokens.

The same schema shape supports both built-in validators and custom registered operators.
"""
		labels: ["overview", "intent"]
	},
	{
		name:  "snk.goals"
		title: "Design Goals"
		markdown: """
1. Keep the schema surface compact.
2. Keep validation strict by default.
3. Avoid ambiguous positional parsing.
4. Keep schemas JSON-compatible for storage and interchange.
5. Support custom validators without a separate mini-language.
"""
		labels: ["overview", "goals"]
	},
	{
		name:     "snk.schema.args"
		title:    "Args Type Model"
		filepath: "examples/args.ts"
		labels:   ["schema", "typescript"]
	},
	{
		name:     "snk.schema.args.admin"
		title:    "Admin Args Model"
		filepath: "examples/args-admin.ts"
		labels:   ["schema", "typescript", "admin"]
	},
	{
		name:     "snk.schema.args.user"
		title:    "User Args Model"
		filepath: "examples/args-user.ts"
		labels:   ["schema", "typescript", "user"]
	},
	{
		name:     "snk.schema.json"
		title:    "Args Command JSON Example"
		filepath: "examples/args-command.json"
		labels:   ["schema", "json"]
	},
	{
		name:     "snk.schema.arguments"
		title:    "CLI Command Examples"
		markdown: """
```sh
# Positional arguments
wash start heavy-duty

# String flags
wash start delicate --temp warm

# Integer flags
wash start normal --spin 1200
wash start normal --spin=1200

# Repeatable values
wash --add=bleach,softener,scent-beads
wash --add=bleach --add=softener

# Boolean flags
wash start bedding --extra-rinse

# Repeatable string flags
wash start whites --add bleach --add softener

# Comma-separated values with repeated option flags
wash start --options delicate,extra-rinse --options pre-wash
```
"""
		labels:   ["schema", "cli", "examples"]
	},
	{
		name:      "snk.usecases"
		title:     "Use Cases"
		filepath:  "examples/args-usecase.csv"
		arguments: ["format-csv=table"]
		labels:    ["usecase", "csv"]
	},
	{
		name:      "snk.validation.matrix"
		title:     "Validation Matrix"
		filepath:  "examples/validation.csv"
		labels:    ["validation", "csv"]
	},
	{
		name:     "snk.validation.string"
		title:    "String Validation API"
		filepath: "examples/string.ts"
		labels:   ["validation", "string", "typescript"]
	},
	{
		name:      "snk.validation.string.enum.csv"
		title:     "Enum Validation Logic"
		filepath:  "examples/enum-validation.csv"
		arguments: ["format-csv=table"]
		labels:    ["validation", "string", "enum", "csv"]
	},
	{
		name:     "snk.validation.string.chain"
		title:    "String Validation Chain"
		filepath: "examples/string-chain.ts"
		labels:   ["validation", "string", "chain"]
	},
	{
		name:     "snk.validation.string.number"
		title:    "String Number Validation"
		filepath: "examples/string-number-check.ts"
		labels:   ["validation", "string", "number"]
	},
	{
		name:     "snk.validation.string.datetime"
		title:    "String Date and Time Validation"
		filepath: "examples/string-date-time.ts"
		labels:   ["validation", "string", "datetime"]
	},
	{
		name:      "snk.validation.string.date.csv"
		title:     "Date Validation Logic"
		filepath:  "examples/date-validation.csv"
		arguments: ["format-csv=table"]
		labels:    ["validation", "string", "date", "csv"]
	},
	{
		name:      "snk.validation.string.datetime.csv"
		title:     "Datetime Validation Logic"
		filepath:  "examples/datetime-validation.csv"
		arguments: ["format-csv=table"]
		labels:    ["validation", "string", "datetime", "csv"]
	},
	{
		name:      "snk.validation.string.time.csv"
		title:     "Time Validation Logic"
		filepath:  "examples/time-validation.csv"
		arguments: ["format-csv=table"]
		labels:    ["validation", "string", "time", "csv"]
	},
	{
		name:      "snk.validation.string.duration.csv"
		title:     "Duration Validation Logic"
		filepath:  "examples/duration-validation.csv"
		arguments: ["format-csv=table"]
		labels:    ["validation", "string", "duration", "csv"]
	},
	{
		name:      "snk.validation.string.email.csv"
		title:     "Email Validation Logic"
		filepath:  "examples/email-validation.csv"
		arguments: ["format-csv=table"]
		labels:    ["validation", "string", "email", "csv"]
	},
	{
		name:      "snk.validation.string.color.csv"
		title:     "Color Validation Logic"
		filepath:  "examples/color-validation.csv"
		arguments: ["format-csv=table"]
		labels:    ["validation", "string", "color", "csv"]
	},
	{
		name:      "snk.validation.string.url.csv"
		title:     "URL Validation Logic"
		filepath:  "examples/url-validation.csv"
		arguments: ["format-csv=table"]
		labels:    ["validation", "string", "url", "csv"]
	},
	{
		name:      "snk.validation.string.arn.csv"
		title:     "AWS ARN Validation Logic"
		filepath:  "examples/arn-validation.csv"
		arguments: ["format-csv=table"]
		labels:    ["validation", "string", "aws", "arn", "csv"]
	},
	{
		name:     "snk.validation.string.formatter-check"
		title:    "String Formatter Check"
		filepath: "examples/string-formatter-check.ts"
		labels:   ["validation", "string", "formatter"]
	},
	{
		name:     "snk.validation.string.boolean-color"
		title:    "String Boolean and Color Validation"
		filepath: "examples/string-boolean-color.ts"
		labels:   ["validation", "string", "boolean", "color"]
	},
	{
		name:      "snk.validation.string.classes.csv"
		title:     "String Classes Reference"
		filepath:  "examples/string-classes.csv"
		arguments: ["format-csv=table"]
		labels:    ["validation", "string", "csv", "classes"]
	},
	{
		name:     "snk.validation.string.classes.ts"
		title:    "String Classes API Example"
		filepath: "examples/string-classes.ts"
		labels:   ["validation", "string", "typescript", "classes"]
	},
	{
		name:      "snk.validation.string.languages.csv"
		title:     "String Languages Reference"
		filepath:  "examples/string-languages.csv"
		arguments: ["format-csv=table"]
		labels:    ["validation", "string", "csv", "languages"]
	},
	{
		name:     "snk.validation.string.languages.ts"
		title:    "String Languages API Example"
		filepath: "examples/string-languages.ts"
		labels:   ["validation", "string", "typescript", "languages"]
	},
	{
		name:     "snk.validation.number"
		title:    "Number Validation API"
		filepath: "examples/number.ts"
		labels:   ["validation", "number", "typescript"]
	},
	{
		name:     "snk.validation.number.chain"
		title:    "Number Validation Chain"
		filepath: "examples/number-chain.ts"
		labels:   ["validation", "number", "chain"]
	},
	{
		name:     "snk.validation.number.converter"
		title:    "Number Converter Examples"
		filepath: "examples/number-converter.ts"
		labels:   ["validation", "number", "conversion"]
	},
	{
		name:     "snk.validation.tuple"
		title:    "Tuple Validation API"
		filepath: "examples/tuple.ts"
		labels:   ["validation", "tuple", "typescript"]
	},
	{
		name:     "snk.validation.tuple.chain"
		title:    "Tuple Validation Chain"
		filepath: "examples/tuple-chain.ts"
		labels:   ["validation", "tuple", "chain"]
	},
	{
		name:     "snk.validation.list"
		title:    "List Validation API"
		filepath: "examples/list.ts"
		labels:   ["validation", "list", "typescript"]
	},
	{
		name:     "snk.validation.list.chain"
		title:    "List Validation Chain"
		filepath: "examples/list-chain.ts"
		labels:   ["validation", "list", "chain"]
	},
	{
		name:      "snk.formatters.csv"
		title:     "Formatter Inventory"
		filepath:  "examples/formatters.csv"
		arguments: ["format-csv=table"]
		labels:    ["formatter", "csv"]
	},
	{
		name:     "snk.formatters.ts"
		title:    "Formatter API"
		filepath: "examples/formatters.ts"
		labels:   ["formatter", "typescript"]
	},
	{
		name:     "snk.formatters.chain"
		title:    "Formatter Chain"
		filepath: "examples/formatter-chain.ts"
		labels:   ["formatter", "chain"]
	},
	{
		name:     "snk.registry.api"
		title:    "Validation Registry API"
		filepath: "examples/validation-registry.ts"
		labels:   ["registry", "typescript"]
	},
	{
		name:     "snk.registry.examples"
		title:    "Validation Registry Examples"
		filepath: "examples/validation-registry-examples.ts"
		labels:   ["registry", "examples", "typescript"]
	},
	{
		name:      "snk.validation.csv"
		title:     "Validation Domain Reference"
		filepath:  "examples/validation.csv"
		labels:    ["validation", "domain", "csv"]
	},
	{
		name:      "snk.cobra.limits"
		title:     "Cobra Limits"
		filepath:  "examples/cobra-limits.csv"
		arguments: ["format-csv=table"]
		labels:    ["constraints", "cobra", "csv"]
	},
	{
		name:  "flow.intent"
		title: "Flow-Specific Specification Intent"
		markdown: """
These flow specs make step order explicit for trusted admin schema authoring and untrusted user input validation.
"""
		labels: ["overview", "flow"]
	},
	{
		name:  "flow.boundaries"
		title: "Flow Boundaries"
		markdown: """
Admin flow covers schema composition and registry registration.
User flow covers argv parsing, operator resolution, and value validation against a pre-defined admin-authored command schema.
Users must not submit schema commands or register new commands.
"""
		labels: ["overview", "flow", "boundary"]
	},
	{
		name:     "flow.admin.capture-usecase"
		title:    "Capture Use Case and Constraints"
		markdown: "Select expected command behavior, accepted value shapes, and repeatability constraints."
		labels:   ["admin-flow"]
	},
	{
		name:     "flow.admin.compose-schema"
		title:    "Compose Schema Commands"
		markdown: "Author argv-style schema commands for each flag and tuple position."
		labels:   ["admin-flow"]
	},
	{
		name:     "flow.admin.register-operator"
		title:    "Register Built-in or Custom Operators"
		markdown: "Register validator operators in the validation registry and reject collisions."
		labels:   ["admin-flow"]
	},
	{
		name:     "flow.admin.export-schema"
		title:    "Export JSON-Compatible Command Schema"
		markdown: "Persist and share the command schema as JSON-friendly arrays of string tokens."
		labels:   ["admin-flow"]
	},
	{
		name:     "flow.user.receive-argv"
		title:    "Receive User Argv Input"
		markdown: "Collect raw argv tokens from untrusted runtime input."
		labels:   ["user-flow"]
	},
	{
		name:     "flow.user.parse-argv"
		title:    "Parse Argv to Typed Flags"
		markdown: "Parse argv into boolean, string, number, and tuple flag values."
		labels:   ["user-flow"]
	},
	{
		name:     "flow.user.resolve-operators"
		title:    "Resolve Validators from Registry"
		markdown: "Resolve each schema operator by domain and name from the registry."
		labels:   ["user-flow"]
	},
	{
		name:     "flow.user.validate-values"
		title:    "Validate Values by Schema"
		markdown: "Validate parsed values with string, number, tuple, list, and formatter-aware rules."
		labels:   ["user-flow"]
	},
	{
		name:     "flow.user.return-result"
		title:    "Return Typed Result or Validation Error"
		markdown: "Return a typed parsed command on success, otherwise a validation error payload."
		labels:   ["user-flow"]
	},
]

relationships: [
	{
		from:   "snk.schema.args"
		to:     "snk.schema.json"
		label:  "explains_shape_of"
		labels: ["explains_shape_of"]
	},
	{
		from:   "snk.schema.args.admin"
		to:     "snk.usecases"
		label:  "supports"
		labels: ["supports", "admin"]
	},
	{
		from:   "snk.schema.args.user"
		to:     "snk.usecases"
		label:  "supports"
		labels: ["supports", "user"]
	},
	{
		from:   "snk.validation.string"
		to:     "snk.validation.matrix"
		label:  "implements"
		labels: ["implements", "string"]
	},
	{
		from:   "snk.validation.number"
		to:     "snk.validation.matrix"
		label:  "implements"
		labels: ["implements", "number"]
	},
	{
		from:   "snk.validation.tuple"
		to:     "snk.validation.matrix"
		label:  "implements"
		labels: ["implements", "tuple"]
	},
	{
		from:   "snk.validation.list"
		to:     "snk.validation.matrix"
		label:  "implements"
		labels: ["implements", "list"]
	},
	{
		from:   "snk.formatters.ts"
		to:     "snk.formatters.csv"
		label:  "documents"
		labels: ["documents", "formatter"]
	},
	{
		from:   "snk.registry.api"
		to:     "snk.registry.examples"
		label:  "demonstrated_by"
		labels: ["demonstrated_by"]
	},
	{
		from:   "snk.registry.api"
		to:     "snk.validation.matrix"
		label:  "extends"
		labels: ["extends", "custom"]
	},
	{
		from:   "flow.admin.capture-usecase"
		to:     "flow.admin.compose-schema"
		label:  "next_step"
		labels: ["next_step", "admin-flow"]
	},
	{
		from:   "flow.admin.compose-schema"
		to:     "flow.admin.register-operator"
		label:  "next_step"
		labels: ["next_step", "admin-flow"]
	},
	{
		from:   "flow.admin.register-operator"
		to:     "flow.admin.export-schema"
		label:  "next_step"
		labels: ["next_step", "admin-flow"]
	},
	{
		from:   "flow.user.receive-argv"
		to:     "flow.user.parse-argv"
		label:  "next_step"
		labels: ["next_step", "user-flow"]
	},
	{
		from:   "flow.user.parse-argv"
		to:     "flow.user.resolve-operators"
		label:  "next_step"
		labels: ["next_step", "user-flow"]
	},
	{
		from:   "flow.user.resolve-operators"
		to:     "flow.user.validate-values"
		label:  "next_step"
		labels: ["next_step", "user-flow"]
	},
	{
		from:   "flow.user.validate-values"
		to:     "flow.user.return-result"
		label:  "next_step"
		labels: ["next_step", "user-flow"]
	},
]

argumentRegistry: {
	version: "1"
	arguments: [
		{
			name:      "format-csv"
			valueType: "enum"
			scopes:    ["note"]
			allowedValues: [
				"table",
			]
			defaultValue: "table"
		},
		{
			name:      "graph-subject-label"
			valueType: "string"
			scopes:    ["h3-section"]
		},
		{
			name:      "graph-edge-label"
			valueType: "string"
			scopes:    ["h3-section"]
		},
		{
			name:      "graph-start-node"
			valueType: "string"
			scopes:    ["h3-section"]
		},
		{
			name:          "graph-renderer"
			valueType:     "enum"
			scopes:        ["h3-section", "note"]
			allowedValues: ["markdown-text", "mermaid"]
			defaultValue:  "markdown-text"
		},
	]
}
