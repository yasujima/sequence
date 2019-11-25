package lexer

import "seqa/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	isNewLine    bool
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	l.isNewLine = true
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
	l.isNewLine = false
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '-':
		newline := l.isNewLine
		for l.peekChar() == '-' {
			l.readChar()
		}
		if newline {
			tok = newToken(token.BULLET, l.ch)
		} else {
			tok = newToken(token.LINE, l.ch)
		}
	case '>':
		tok = newToken(token.GT, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case ':':
		fallthrough
	case '#':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '\r':
		fallthrough
	case '\n':
		tok = newToken(token.CRLF, l.ch)
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.IDENT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	if tok.Type == token.CRLF {
		l.isNewLine = true
	}
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' { //|| l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tt token.TokenType, ch byte) token.Token {
	return token.Token{Type: tt, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readString() string {
	l.readChar()
	l.skipWhitespace()
	position := l.position
	for {
		l.readChar()
		if l.peekChar() == 0 || l.peekChar() == '\r' || l.peekChar() == '\n' {
			break
		}
	}
	return l.input[position:l.readPosition]
}
