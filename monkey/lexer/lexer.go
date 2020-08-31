// Package lexer = lexer/lexer.go
package lexer

import (
	"monkey/token"
)

// Lexer is a data structure representing a lexer
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// New Creates a new Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// NextToken returns the next token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	tok.Col = l.position
	l.skipWhitespace()

	switch l.ch {
	case '+':
		tok = newToken(token.PLUS, l.ch, l.position)
	case '-':
		tok = newToken(token.MINUS, l.ch, l.position)
	case '/':
		tok = newToken(token.SLASH, l.ch, l.position)
	case '*':
		tok = newToken(token.ASTERISK, l.ch, l.position)

	case '=':
		tok = l.newTwoByteToken()
	case '!':
		tok = l.newTwoByteToken()

	case ',':
		tok = newToken(token.COMMA, l.ch, l.position)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch, l.position)
	case '(':
		tok = newToken(token.LPAREN, l.ch, l.position)
	case ')':
		tok = newToken(token.RPAREN, l.ch, l.position)
	case '{':
		tok = newToken(token.LBRACE, l.ch, l.position)
	case '}':
		tok = newToken(token.RBRACE, l.ch, l.position)
	case '[':
		tok = newToken(token.LBRACKET, l.ch, l.position)
	case ']':
		tok = newToken(token.RBRACKET, l.ch, l.position)
	case '<':
		tok = newToken(token.LT, l.ch, l.position)
	case '>':
		tok = newToken(token.GT, l.ch, l.position)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			tok.Col = l.position
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			tok.Col = l.position
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch, l.position)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' ||
		l.ch == '\t' ||
		l.ch == '\n' ||
		l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		} else if l.ch == '\\' && l.peekChar() == '"' {
			l.readChar()
		}
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' ||
		'A' <= ch && ch <= 'Z' ||
		ch == '_'
}

func (l *Lexer) newTwoByteToken() token.Token {
	var tokenType token.TokenType = token.ILLEGAL
	ch := l.ch
	if l.peekChar() == '=' {
		if ch == '!' {
			tokenType = token.NOT_EQ
		} else if ch == '=' {
			tokenType = token.EQ
		}
		l.readChar()
		literal := string(ch) + string(l.ch)
		return token.Token{Type: tokenType, Literal: literal}
	}
	if ch == '!' {
		tokenType = token.BANG
	} else if ch == '=' {
		tokenType = token.ASSIGN
	}
	return newToken(tokenType, l.ch, l.position)
}

func newToken(tokenType token.TokenType, ch byte, col int) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch), Col: col}
}
