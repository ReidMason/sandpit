package ast

import (
	"bytes"
	"moviescript/internal/token"
)

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Identifier) String() string {
	return i.Value
}

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
}

type Expression interface {
	Node
}

type Program struct {
	Statements []Statement
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	return out.String()
}

type AssignmentStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (as *AssignmentStatement) TokenLiteral() string {
	return as.Token.Literal
}
func (as *AssignmentStatement) String() string {
	var out bytes.Buffer

	out.WriteString(" = ")
	out.WriteString(as.Name.Value)

	return out.String()
}
