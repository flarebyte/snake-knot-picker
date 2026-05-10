// purpose: Define the public command, flag, and parse-result data contracts shared by compile and runtime paths.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package picker

type CommandDocument struct {
	Version     string           `json:"version"`
	CommandPath []string         `json:"commandPath"`
	AdminOnly   bool             `json:"adminOnly"`
	Flags       []CommandFlagDef `json:"flags"`
}

type CommandFlagDef struct {
	Kind    string     `json:"kind"`
	Name    string     `json:"name"`
	Schema  []string   `json:"schema"`
	Schemas [][]string `json:"schemas,omitempty"`
}

type CompiledCommand struct {
	CommandPath []string
	AdminOnly   bool
	Flags       []CompiledFlag
}

type CompiledFlag struct {
	Kind       string
	Name       string
	TupleSize  int
	Repeatable bool
}

type ParseResult struct {
	CommandPath []string
	Values      map[string]Value
}

type Value struct {
	Bool   *bool
	String *string
	Number *float64
	List   []Value
	Tuple  []Value
}
