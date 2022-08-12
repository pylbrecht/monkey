package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatal("tests[", i, "] - tokentype wrong. expected=", tt.expectedType, ", got=", tok.Type, "")
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatal("tests[", i, "] - tokentype wrong. expected=", tt.expectedLiteral, ", got=", tok.Literal, "")
		}
	}
}
