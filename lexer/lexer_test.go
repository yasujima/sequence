package lexer

import (
	"seqa/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	//	input := `=+(){},;`
	input := `aa->bb
bb < - cc
- cc-->dd`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "aa"},
		{token.LINE, "-"},
		{token.GT, ">"},
		{token.IDENT, "bb"},
		{token.CRLF, "\n"},
		{token.IDENT, "bb"},
		{token.LT, "<"},
		{token.LINE, "-"},
		{token.IDENT, "cc"},
		{token.CRLF, "\n"},
		{token.BULLET, "-"},
		{token.IDENT, "cc"},
		{token.LINE, "-"},
		{token.GT, ">"},
		{token.IDENT, "dd"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] - leteral wrong, expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
		t.Logf("test[%d] token %q: %q", i, tok.Type, tok.Literal)
	}
}
