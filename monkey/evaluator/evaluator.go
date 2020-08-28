// Package evaluator evaluator/evaluator.go
package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

var (
	// NULL var
	NULL = &object.Null{}
	// TRUE var
	TRUE = &object.Boolean{Value: true}
	// FALSE var
	FALSE = &object.Boolean{Value: false}
)

// Eval main
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// Statements
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
		// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	}
	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range stmts {
		result = Eval(statement)
	}
	return result
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}
