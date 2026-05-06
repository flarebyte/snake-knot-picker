package picker

func makeInvalidTupleMissingSlotDoc() CommandDocument {
	return CommandDocument{
		Version:     "1",
		CommandPath: []string{"wash", "start"},
		Flags: []CommandFlagDef{
			{
				Kind:   "tuple",
				Name:   "range",
				Schema: []string{"schema", "tuple", "--size", "2"},
				Schemas: [][]string{
					{"schema", "number", "--tuple", "0", "--int"},
				},
			},
		},
	}
}
