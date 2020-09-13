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
	// FreeScope const
	FreeScope SymbolScope = "FREE"
	// FunctionScope const
	FunctionScope SymbolScope = "FUNCTION"
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
	FreeSymbols    []Symbol
}

// NewSymbolTable func
func NewSymbolTable() *SymbolTable {
	s := make(map[string]Symbol)
	free := []Symbol{}
	return &SymbolTable{store: s, FreeSymbols: free}
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
		if !ok {
			return obj, ok
		}
		if obj.Scope == GlobalScope || obj.Scope == BuiltinScope {
			return obj, ok
		}
		free := s.defineFree(obj)
		return free, true
	}
	return obj, ok
}

// DefineBuiltin func
func (s *SymbolTable) DefineBuiltin(index int, name string) Symbol {
	symbol := Symbol{Name: name, Index: index, Scope: BuiltinScope}
	s.store[name] = symbol
	return symbol
}

// defineFree func
func (s *SymbolTable) defineFree(original Symbol) Symbol {
	s.FreeSymbols = append(s.FreeSymbols, original)
	symbol := Symbol{
		Name:  original.Name,
		Index: len(s.FreeSymbols) - 1,
		Scope: FreeScope,
	}
	s.store[original.Name] = symbol
	return symbol
}

// DefineFunctionName func
func (s *SymbolTable) DefineFunctionName(name string) Symbol {
	symbol := Symbol{Name: name, Index: 0, Scope: FunctionScope}
	s.store[name] = symbol
	return symbol
}
