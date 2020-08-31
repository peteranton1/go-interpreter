// Package ast ast/ast.go
package ast

import (
	"bytes"
	"monkey/token"
	"strings"
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
	}
	return ""
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

// StringLiteral struct
type StringLiteral struct {
	Token token.Token
	Value string
}

// expressionNode interface method
func (sl *StringLiteral) expressionNode() {}

// TokenLiteral interface method
func (sl *StringLiteral) TokenLiteral() string {
	return sl.Token.Literal
}

// String interface method
func (sl *StringLiteral) String() string {
	return sl.Token.Literal
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
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

// InfixExpression struct
type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}

// TokenLiteral interface method
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// String interface method
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}

// Boolean struct
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

// TokenLiteral interface method
func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

// String interface method
func (b *Boolean) String() string {
	return b.Token.Literal
}

// IfExpression struct
type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

// TokenLiteral interface method
func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// String interface method
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(" ")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}

// BlockStatement struct
type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

// TokenLiteral interface method
func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

// String interface method
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// FunctionLiteral struct
type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}

// TokenLiteral interface method
func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

// String interface method
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())
	return out.String()
}

// CallExpression struct
type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}

// TokenLiteral interface method
func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

// String interface method
func (ce *CallExpression) String() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}

// ArrayLiteral struct
type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode() {}

// TokenLiteral interface method
func (al *ArrayLiteral) TokenLiteral() string {
	return al.Token.Literal
}

// String interface method
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer
	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

// IndexExpression struct
type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode() {}

// TokenLiteral interface method
func (ie *IndexExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// String interface method
func (ie *IndexExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")
	return out.String()
}
