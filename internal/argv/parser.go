package argv

import "github.com/flarebyte/snake-knot-picker"

type Parser struct{}

func NewParser() *Parser { return &Parser{} }

func (p *Parser) Parse(command picker.CompiledCommand, argv []string) (*picker.ParseResult, error) {
	return picker.Parse(command, argv)
}
