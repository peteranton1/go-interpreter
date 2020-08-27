// Package ast ast/ast.go
package ast

import (
	"bytes"
	"monkey/token"
)

// Node base type
type Node interface {
	TokenLiteral() string
	String() string
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

// String interface method
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
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

// String interface method
func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
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

// String interface method
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// ExpressionStatement struct
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral interface method
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// String interface method
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// Identifier struct
type Identifier struct {
	Token token.Token
	Value string
}

// expressionNode interface method
func (i *Identifier) expressionNode() {}

// TokenLiteral interface method
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// String interface method
func (i *Identifier) String() string {
	return i.Value
}

// IntegerLiteral struct
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

// expressionNode interface method
func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral interface method
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

// String interface method
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

// PrefixExpression struct
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral interface method
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

// String interface method
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("!")
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}
