package picker

type CommandBuilder struct {
	doc CommandDocument
}

func NewCommandBuilder(commandPath ...string) *CommandBuilder {
	return &CommandBuilder{
		doc: CommandDocument{
			Version:     "1",
			CommandPath: append([]string(nil), commandPath...),
		},
	}
}

func (b *CommandBuilder) SetAdminOnly(adminOnly bool) *CommandBuilder {
	b.doc.AdminOnly = adminOnly
	return b
}

func (b *CommandBuilder) AddFlag(flag CommandFlagDef) *CommandBuilder {
	b.doc.Flags = append(b.doc.Flags, flag)
	return b
}

func (b *CommandBuilder) Build() CommandDocument {
	out := b.doc
	out.CommandPath = append([]string(nil), b.doc.CommandPath...)
	out.Flags = append([]CommandFlagDef(nil), b.doc.Flags...)
	return out
}
