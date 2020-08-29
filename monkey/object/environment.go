// Package object object/environment.go
package object

// NewEnvironment method
func NewEnvironment() *Environment {
	s := make(map[string]Object)

	return &Environment{store: s}
}

// Environment struct
type Environment struct {
	store map[string]Object
}

// Get object method
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

// Set object method
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
