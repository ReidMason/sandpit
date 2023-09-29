package main

import (
	"log"
	"moviescript/internal/lexer"
	"moviescript/internal/parser"
	"moviescript/internal/translator"
	"os"
)

func main() {
	input := "theres this movie called myObject which is 5000"

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	output := translator.Translate(program)

	output += "\nconsole.log(myObject)"

	data := []byte(output)
	err := os.WriteFile("output.js", data, 0644)
	if err != nil {
		log.Fatal("Failed to write output to file")
	}
}
