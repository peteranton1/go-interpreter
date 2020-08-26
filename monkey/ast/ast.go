// Package ast ast/ast.go
package ast

import (
	"monkey/token"
)

// Node base type
type Node interface {
	TokenLiteral() string
}

// Statement subtype
type Statement interface {
	Node
	statementNode()
}

// Expression subtype
type Expression interface {
	Node
	expressionNode()
}

// Program struct
type Program struct {
	Statements []Statement
}

// TokenLiteral interface method
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement struct
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// statementNode interface method
func (ls *LetStatement) statementNode() {}

// TokenLiteral interface method
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier struct
type Identifier struct {
	Token token.Token
	Value string
}

// statementNode interface method
func (i *Identifier) expressionNode() {}

// TokenLiteral interface method
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// ReturnStatement struct
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral interface method
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
