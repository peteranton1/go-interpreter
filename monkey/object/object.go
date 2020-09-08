// Package object object/object.go
package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"monkey/ast"
	"monkey/code"
	"strings"
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
	// RETURN_VALUE_OBJ const
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	// ERROR_OBJ const
	ERROR_OBJ = "ERROR"
	// FUNCTION_OBJ const
	FUNCTION_OBJ = "FUNCTION"
	// STRING_OBJ const
	STRING_OBJ = "STRING"
	// BUILTIN_OBJ const
	BUILTIN_OBJ = "BUILTIN"
	// ARRAY_OBJ const
	ARRAY_OBJ = "ARRAY"
	// HASH_OBJ const
	HASH_OBJ = "HASH"
	// COMPILED_FUNCTION_OBJ const
	COMPILED_FUNCTION_OBJ = "COMPILED_FUNCTION_OBJ"
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

// ReturnValue struct
type ReturnValue struct {
	Value Object
}

// Type interface method
func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

// Inspect interface method
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

// Error struct
type Error struct {
	Message string
}

// Type interface method
func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}

// Inspect interface method
func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

// Function struct
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type interface method
func (f *Function) Type() ObjectType {
	return FUNCTION_OBJ
}

// Inspect interface method
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}

	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

// String struct
type String struct {
	Value string
}

// Type interface method
func (s *String) Type() ObjectType {
	return STRING_OBJ
}

// Inspect interface method
func (s *String) Inspect() string {
	return s.Value
}

// BuiltinFunction type
type BuiltinFunction func(args ...Object) Object

// Builtin struct
type Builtin struct {
	Fn BuiltinFunction
}

// Type interface method
func (b *Builtin) Type() ObjectType {
	return BUILTIN_OBJ
}

// Inspect interface method
func (b *Builtin) Inspect() string {
	return "builtin function"
}

// Array struct
type Array struct {
	Elements []Object
}

// Type interface method
func (ao *Array) Type() ObjectType {
	return ARRAY_OBJ
}

// Inspect interface method
func (ao *Array) Inspect() string {
	var out bytes.Buffer
	elements := []string{}

	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// Hashable interface
type Hashable interface {
	HashKey() HashKey
}

// HashKey struct
type HashKey struct {
	Type  ObjectType
	Value uint64
}

// HashPair struct
type HashPair struct {
	Key   Object
	Value Object
}

// Hash struct
type Hash struct {
	Pairs map[HashKey]HashPair
}

// Type interface method
func (h *Hash) Type() ObjectType {
	return HASH_OBJ
}

// Inspect interface method
func (h *Hash) Inspect() string {
	var out bytes.Buffer
	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}

// HashKey interface method
func (b *Boolean) HashKey() HashKey {
	var value uint64
	if b.Value {
		value = 1
	} else {
		value = 0
	}
	return HashKey{Type: b.Type(), Value: value}
}

// HashKey interface method
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

// HashKey interface method
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

// CompiledFunction struct
type CompiledFunction struct {
	Instructions code.Instructions
}

// Type func
func (cf *CompiledFunction) Type() ObjectType {
	return COMPILED_FUNCTION_OBJ
}

// Inspect func
func (cf *CompiledFunction) Inspect() string {
	return fmt.Sprintf("CompiledFunction[%p]", cf)
}
