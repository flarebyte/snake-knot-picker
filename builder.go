// purpose: Provide a small fluent builder for creating command documents programmatically with predictable defaults.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package picker

// CommandBuilder builds CommandDocument values with fluent chaining.
type CommandBuilder struct {
	doc CommandDocument
}

// NewCommandBuilder creates a builder initialized with version "1" and command path.
func NewCommandBuilder(commandPath ...string) *CommandBuilder {
	return &CommandBuilder{
		doc: CommandDocument{
			Version:     "1",
			CommandPath: append([]string(nil), commandPath...),
		},
	}
}

// SetAdminOnly marks whether the resulting document is admin-only.
func (b *CommandBuilder) SetAdminOnly(adminOnly bool) *CommandBuilder {
	b.doc.AdminOnly = adminOnly
	return b
}

// AddFlag appends one flag definition to the document.
func (b *CommandBuilder) AddFlag(flag CommandFlagDef) *CommandBuilder {
	b.doc.Flags = append(b.doc.Flags, flag)
	return b
}

// Build returns a detached CommandDocument snapshot from the builder.
func (b *CommandBuilder) Build() CommandDocument {
	out := b.doc
	out.CommandPath = append([]string(nil), b.doc.CommandPath...)
	out.Flags = append([]CommandFlagDef(nil), b.doc.Flags...)
	return out
}
