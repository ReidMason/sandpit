package parser

import (
	"fmt"
	"moviescript/internal/ast"
	"moviescript/internal/lexer"
	"moviescript/internal/token"
)

type Parser struct {
	l *lexer.Lexer

	errors []string

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		statement := p.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}

		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.THERES:
		return p.parseTheresSatement()
	default:
		return nil
	}
}

func (p *Parser) parseTheresSatement() *ast.LetStatement {
	statement := &ast.LetStatement{Token: token.Token{Type: "LET", Literal: "let"}}

	path := [4]token.TokenType{token.THIS, token.MOVIE, token.CALLED, token.IDENT}
	for _, tokenType := range path {
		if !p.expectPeek(tokenType) {
			return nil
		}
	}

	statement.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	return statement
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}

	p.peekError(t)
	return false
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}
