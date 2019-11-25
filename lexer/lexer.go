package lexer

import "seqa/token"

type Lexer struct {
	input string
	position int
	readPosition int
	istop bool
	ch byte
}

func New(input string) *Lexer {
	l := &Lexer{input:input}
	l.readChar()
	l.istop = true
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
	l.istop = false
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '-':
		if l.istop {
			tok = newToken(token.BULLET, l.ch)
		} else {
			for l.peekChar() == '-' {
				l.readChar()
			}
			tok = newToken(token.LINE, l.ch)
		}
	case '>':
		tok = newToken(token.GT, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '\r':
		tok = newToken(token.CRLF, l.ch)
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
		l.istop = true
	}
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' {//|| l.ch == '\n' || l.ch == '\r' {
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
