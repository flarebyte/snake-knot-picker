// purpose: Provide a small fluent builder for creating command documents programmatically with predictable defaults.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
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
