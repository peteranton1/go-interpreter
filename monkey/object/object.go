// Package object object/object.go
package object

import (
	"fmt"
)

// ObjectType type
type ObjectType string

const (
	// INTEGER_OBJ const
	INTEGER_OBJ = "INTEGER"
	// BOOLEAN_OBJ const
	BOOLEAN_OBJ = "BOOLEAN"
	// NULL_OBJ const
	NULL_OBJ = "NULL"
)

// Object interface
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer struct
type Integer struct {
	Value int64
}

// Type interface method
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

// Inspect interface method
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// Boolean struct
type Boolean struct {
	Value bool
}

// Type interface method
func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

// Inspect interface method
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// Null struct
type Null struct{}

// Type interface method
func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

// Inspect interface method
func (n *Null) Inspect() string {
	return "null"
}
