// purpose: Provide operator-factory registration and lookup so schema operators can be validated against known capabilities.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: Built-in operators are pre-registered at construction time and duplicate registration is explicitly rejected.
package picker

import "sort"

type ValidatorFactory interface {
	Name() string
}

type Registry interface {
	Register(factory ValidatorFactory) error
	Lookup(operator string) (ValidatorFactory, bool)
}

type registry struct {
	factories map[string]ValidatorFactory
}

func NewRegistry() Registry {
	r := &registry{factories: map[string]ValidatorFactory{}}
	for _, name := range builtInOperators() {
		_ = r.Register(staticFactory{name: name})
	}
	return r
}

func (r *registry) Register(factory ValidatorFactory) error {
	if factory == nil || factory.Name() == "" {
		return NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": "operator"})
	}
	name := factory.Name()
	if _, exists := r.factories[name]; exists {
		return NewSchemaError(ErrorIDSchemaDuplicateRegistration, map[string]string{"operator": name})
	}
	r.factories[name] = factory
	return nil
}

func (r *registry) Lookup(operator string) (ValidatorFactory, bool) {
	f, ok := r.factories[operator]
	return f, ok
}

func (r *registry) Operators() []string {
	out := make([]string, 0, len(r.factories))
	for k := range r.factories {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

type staticFactory struct {
	name string
}

func (f staticFactory) Name() string { return f.name }

func NewStaticFactory(name string) ValidatorFactory {
	return staticFactory{name: name}
}

func builtInOperators() []string {
	return []string{
		"boolean",
		"number",
		"string",
		"tuple",
		"repeatable",
		"postal-code",
	}
}
