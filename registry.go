package picker

type ValidatorFactory interface {
	Name() string
}

type Registry interface {
	Register(factory ValidatorFactory) error
	Lookup(operator string) (ValidatorFactory, bool)
}

