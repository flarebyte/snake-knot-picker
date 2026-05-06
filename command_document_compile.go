package picker

func CompileCommandDocument(doc CommandDocument) (CompiledCommand, error) {
	if err := validateCommandDocument(doc); err != nil {
		return CompiledCommand{}, err
	}

	out := CompiledCommand{
		CommandPath: append([]string(nil), doc.CommandPath...),
		AdminOnly:   doc.AdminOnly,
		Flags:       make([]CompiledFlag, 0, len(doc.Flags)),
	}
	for _, flag := range doc.Flags {
		cf, err := compileFlag(flag)
		if err != nil {
			return CompiledCommand{}, err
		}
		out.Flags = append(out.Flags, cf)
	}
	return out, nil
}

func compileFlag(flag CommandFlagDef) (CompiledFlag, error) {
	shape, err := parsePrimarySchema(flag.Schema)
	if err != nil {
		return CompiledFlag{}, err
	}
	out := CompiledFlag{
		Kind:      flag.Kind,
		Name:      flag.Name,
		TupleSize: shape.TupleSize,
	}
	for _, child := range flag.Schemas {
		if len(child) >= 2 && child[0] == "schema" && child[1] == "repeatable" {
			out.Repeatable = true
			break
		}
	}
	return out, nil
}
