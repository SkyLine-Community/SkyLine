package SkyLine

func NewEnvironment() Environment {
	return &Environment_of_environment{
		Store: make(map[string]Object),
		Outer: nil,
	}
}

func (e *Environment_of_environment) Get(name string) (Object, bool) {
	obj, exists := e.Store[name]
	if !exists && e.Outer != nil {
		obj, exists = e.Outer.Get(name)
	}
	return obj, exists
}

func (e *Environment_of_environment) Set(name string, val Object) Object {
	e.Store[name] = val
	return val
}

func NewEnclosedEnvironment(outer Environment) Environment {
	return &Environment_of_environment{
		Store: make(map[string]Object),
		Outer: outer,
	}
}
