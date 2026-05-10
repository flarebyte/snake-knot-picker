// purpose: Compile parsed schema AST into compiled specs after registry and rule validation.
// responsibilities: Validate compile preconditions, collect flags, run compiler rules, and build compiled specs.
// architecture notes: Operator resolution is registry-driven so custom operators can be introduced without changing parser structure.
package schema

import "github.com/flarebyte/snake-knot-picker"

// Compiler compiles parsed schema AST nodes into compiled specs.
type Compiler struct {
	registry picker.Registry
}

// NewCompiler creates a schema compiler bound to an operator registry.
func NewCompiler(registry picker.Registry) *Compiler {
	return &Compiler{registry: registry}
}

// Compile validates compile preconditions and returns a compiled schema spec.
func (c *Compiler) Compile(ast *CommandAST) (*CompiledSpec, error) {
	if err := c.validateCompileInputs(ast); err != nil {
		return nil, err
	}
	flagSet, err := collectFlags(ast)
	if err != nil {
		return nil, err
	}
	if err := validateCompilerRules(flagSet); err != nil {
		return nil, err
	}
	return buildCompiledSpec(ast, flagSet), nil
}

func (c *Compiler) validateCompileInputs(ast *CommandAST) error {
	if ast == nil {
		return picker.NewSchemaError(picker.ErrorIDSchemaInvalidValue, map[string]string{"field": "ast"})
	}
	if c.registry == nil {
		return picker.NewSchemaError(picker.ErrorIDSchemaInvalidValue, map[string]string{"field": "registry"})
	}
	if _, ok := c.registry.Lookup(ast.Operator); !ok {
		return picker.NewSchemaError(picker.ErrorIDSchemaUnknownOperator, map[string]string{"operator": ast.Operator})
	}
	return nil
}
