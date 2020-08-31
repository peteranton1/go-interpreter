// Package token token/token.go
package token

// TokenType type
type TokenType string

// Token struct
type Token struct {
	Type    TokenType
	Literal string
	Col     int
}

const (
	// ILLEGAL Control chars
	ILLEGAL = "ILLEGAL"
	// EOF Control chars
	EOF = "EOF"

	// IDENT Identifiers
	IDENT = "IDENT"
	// INT datatype
	INT = "INT"
	// STRING datatype
	STRING = "STRING"

	// ASSIGN Operators
	ASSIGN = "="
	// PLUS Operator
	PLUS = "+"
	// MINUS Operator
	MINUS = "-"
	// SLASH Operator
	SLASH = "/"
	// ASTERISK Operator
	ASTERISK = "*"

	// BANG Delimiters
	BANG = "!"
	// COMMA Delimiters
	COMMA = ","
	// SEMICOLON Delimiter
	SEMICOLON = ";"
	// LPAREN Delimiter
	LPAREN = "("
	// RPAREN Delimiter
	RPAREN = ")"
	// LBRACE Delimiter
	LBRACE = "{"
	// RBRACE Delimiter
	RBRACE = "}"
	// LBRACKET Delimiter
	LBRACKET = "["
	// RBRACKET Delimiter
	RBRACKET = "]"
	// LT Delimiter
	LT = "<"
	// GT Delimiter
	GT = ">"
	// EQ Delimiter
	EQ = "=="
	// NOT_EQ Delimiter
	NOT_EQ = "!="

	// FUNCTION Keyword
	FUNCTION = "FUNCTION"
	// LET Keyword
	LET = "LET"
	// IF Keyword
	IF = "IF"
	// ELSE Keyword
	ELSE = "ELSE"
	// RETURN Keyword
	RETURN = "RETURN"
	// TRUE Keyword
	TRUE = "TRUE"
	// FALSE Keyword
	FALSE = "FALSE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

// LookupIdent returns Ident or Keyword
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
