package translator

import (
	"moviescript/internal/ast"
)

func Translate(program *ast.Program) string {
	output := ""
	for _, statement := range program.Statements {
		output += statement.String()
	}

	return output
}
