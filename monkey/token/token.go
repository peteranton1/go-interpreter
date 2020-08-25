// Package token token/token.go
package token

// TokenType type
type TokenType string

// Token struct
type Token struct {
	Type           TokenType
	Literal        string
	SourcePosition int
}

const (
	// ILLEGAL Control chars
	ILLEGAL = "ILLEGAL"
	// EOF Control chars
	EOF = "EOF"

	// IDENT Identifiers + Literals
	IDENT = "IDENT"
	// INT Integer
	INT = "INT"

	// ASSIGN Operators
	ASSIGN = "="
	// PLUS Operator
	PLUS = "+"

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

	// FUNCTION Keywords
	FUNCTION = "FUNCTION"
	// LET Keywords
	LET = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent returns Ident or Keyword
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
