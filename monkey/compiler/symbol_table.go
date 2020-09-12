// Package compiler compiler/symbol_table.go
package compiler

// SymbolScope string
type SymbolScope string

const (
	// GlobalScope const
	GlobalScope SymbolScope = "GLOBAL"
	// LocalScope const
	LocalScope SymbolScope = "LOCAL"
	// BuiltinScope const
	BuiltinScope SymbolScope = "BUILTIN"
)

// Symbol struct
type Symbol struct {
	Name  string
	Scope SymbolScope
	Index int
}

// SymbolTable struct
type SymbolTable struct {
	Outer          *SymbolTable
	store          map[string]Symbol
	numDefinitions int
}

// NewSymbolTable func
func NewSymbolTable() *SymbolTable {
	s := make(map[string]Symbol)
	return &SymbolTable{store: s}
}

// NewEnclosedSymbolTable func
func NewEnclosedSymbolTable(outer *SymbolTable) *SymbolTable {
	s := NewSymbolTable()
	s.Outer = outer
	return s
}

// Define func
func (s *SymbolTable) Define(name string) Symbol {
	symbol := Symbol{Name: name, Index: s.numDefinitions}
	if s.Outer == nil {
		symbol.Scope = GlobalScope
	} else {
		symbol.Scope = LocalScope
	}
	s.store[name] = symbol
	s.numDefinitions++
	return symbol
}

// Resolve func
func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := s.store[name]
	if !ok && s.Outer != nil {
		obj, ok = s.Outer.Resolve(name)
		return obj, ok
	}
	return obj, ok
}

// DefineBuiltin func
func (s *SymbolTable) DefineBuiltin(index int, name string) Symbol {
	symbol := Symbol{Name: name, Index: index, Scope: BuiltinScope}
	s.store[name] = symbol
	return symbol
}
