// Package ast ast/ast_test.go
package ast

import (
	"monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	expectedStmt := "let myVar = anotherVar;"

	if program.String() != expectedStmt {
		t.Errorf("program.String() wrong, "+
			"\n\texpected=%q"+
			"\n\t     got=%q",
			expectedStmt, program.String())
	}
}
