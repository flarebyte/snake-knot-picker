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
