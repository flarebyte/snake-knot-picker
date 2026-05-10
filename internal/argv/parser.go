// purpose: Provide a thin argv parser adapter that delegates to the core picker runtime parser.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package argv

import "github.com/flarebyte/snake-knot-picker"

// Parser adapts the top-level picker parser for internal argv workflows.
type Parser struct{}

// NewParser creates a new argv parser adapter.
func NewParser() *Parser { return &Parser{} }

// Parse validates argv against a compiled command.
func (p *Parser) Parse(command picker.CompiledCommand, argv []string) (*picker.ParseResult, error) {
	return picker.Parse(command, argv)
}
