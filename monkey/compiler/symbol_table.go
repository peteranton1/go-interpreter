// Package compiler compiler/symbol_table.go
package compiler

// SymbolScope string
type SymbolScope string

const (
	// GlobalScope const
	GlobalScope SymbolScope = "GLOBAL"
)

// Symbol struct
type Symbol struct {
	Name  string
	Scope SymbolScope
	Index int
}

// SymbolTable struct
type SymbolTable struct {
	store          map[string]Symbol
	numDefinitions int
}

// NewSymbolTable func
func NewSymbolTable() *SymbolTable {
	s := make(map[string]Symbol)
	return &SymbolTable{store: s}
}

// Define func
func (s *SymbolTable) Define(name string) Symbol {
	symbol := Symbol{Name: name, Index: s.numDefinitions, Scope: GlobalScope}
	s.store[name] = symbol
	s.numDefinitions++
	return symbol
}

// Resolve func
func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := s.store[name]
	return obj, ok
}
