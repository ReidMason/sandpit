package lexer

import (
	"moviescript/internal/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `theres this movie called myObject`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.THERES, "theres"},
		{token.THIS, "this"},
		{token.MOVIE, "movie"},
		{token.CALLED, "called"},
		{token.IDENT, "myObject"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		t.Log(tok, i, tt)
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. Expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. Expected=%q, got =%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
