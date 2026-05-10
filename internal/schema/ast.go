// purpose: Define schema-token AST types used between schema parsing and compilation stages.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package schema

type ParsedSchemaFlag struct {
	Name   string
	Values []string
}

type CommandAST struct {
	Head       string
	Operator   string
	Flags      []ParsedSchemaFlag
	Raw        []string
	Path       []string
	TupleIndex *int
}
