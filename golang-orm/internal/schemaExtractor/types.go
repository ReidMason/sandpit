package schemaextractor

import "strings"

type GoType string

func GetGoType(sqlType string) GoType {
	switch sqlType {
	case "INT":
	case "SERIAL":
		return "int"
	case "TEXT":
		return "string"
	case "BOOLEAN":
		return "bool"
	}

	if strings.Contains(sqlType, "VARCHAR") {
		return "string"
	}

	return ""
}
