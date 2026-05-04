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
	Kind string
	Name string
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

