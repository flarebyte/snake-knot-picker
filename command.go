// purpose: Define the public command, flag, and parse-result data contracts shared by compile and runtime paths.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package picker

// CommandDocument is the persisted/admin-authored command schema document.
type CommandDocument struct {
	Version     string           `json:"version"`
	CommandPath []string         `json:"commandPath"`
	AdminOnly   bool             `json:"adminOnly"`
	Flags       []CommandFlagDef `json:"flags"`
}

// CommandFlagDef defines one flag in a command document.
type CommandFlagDef struct {
	Kind    string     `json:"kind"`
	Name    string     `json:"name"`
	Schema  []string   `json:"schema"`
	Schemas [][]string `json:"schemas,omitempty"`
}

// CompiledCommand is the runtime-ready immutable representation of a command.
type CompiledCommand struct {
	CommandPath []string
	AdminOnly   bool
	Flags       []CompiledFlag
}

// CompiledFlag is the runtime-ready immutable representation of one flag.
type CompiledFlag struct {
	Kind       string
	Name       string
	TupleSize  int
	Repeatable bool
}

// ParseResult is the typed output returned after successful runtime parsing.
type ParseResult struct {
	CommandPath []string
	Values      map[string]Value
}

// Value stores one parsed typed value and supports scalar, tuple, and list shapes.
type Value struct {
	Bool   *bool
	String *string
	Number *float64
	List   []Value
	Tuple  []Value
}
