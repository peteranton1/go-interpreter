// Package object object/environment.go
package object

// NewEnvironment method
func NewEnvironment() *Environment {
	s := make(map[string]Object)

	return &Environment{store: s, outer: nil}
}

// NewEnclosedEnvironment method
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer

	return env
}

// Environment struct
type Environment struct {
	store map[string]Object
	outer *Environment
}

// Get object method
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set object method
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
