package parser

import (
	"moviescript/internal/ast"
	"moviescript/internal/lexer"
	"moviescript/internal/token"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `theres this movie called myObject which is 5`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	// checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier ast.LetStatement
	}{
		{
			ast.LetStatement{
				Token: token.Token{Type: "LET", Literal: "let"},
				Name:  &ast.Identifier{Token: token.Token{Type: token.IDENT, Literal: "myObject"}, Value: "myObject"},
				Value: &ast.AssignmentStatement{
					Token: token.Token{
						Type:    "EQUALS",
						Literal: "equals",
					},
					Name: &ast.Identifier{
						Token: token.Token{Type: token.INT, Literal: "int"},
						Value: "5"},
				},
			},
		},
	}

	for i, tt := range tests {
		statement := program.Statements[i]
		if !testLetStatement(t, statement, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, statement ast.LetStatement) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStatement, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStatement.Name.Value != statement.Name.Value {
		t.Errorf("letStatement.Name.Value not '%s'. got=%s", statement.Name.Value, letStatement.Name.Value)
		return false
	}

	if letStatement.Name.TokenLiteral() != statement.Name.TokenLiteral() {
		t.Errorf("s.Name not '%s'. got='%s'", statement.Name.TokenLiteral(), letStatement.Name.TokenLiteral())
		return false
	}

	if letStatement.Value.TokenLiteral() != statement.Value.TokenLiteral() {
		t.Errorf("s.Value.TokenLitera() not '%s'. got='%s'", statement.Value.TokenLiteral(), letStatement.Value.TokenLiteral())
		return false
	}

	return true
}
