package translator

import (
	"moviescript/internal/lexer"
	"moviescript/internal/parser"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `theres this movie called myObject`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()

	output := Translate(program)
	expected := "let myObject"

	if output != expected {
		t.Fatalf("Output mismatch. Expected '%s' got '%s'", expected, output)
	}
}
