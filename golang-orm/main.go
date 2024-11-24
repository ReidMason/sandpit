package main

import (
	schemaextractor "ReidMason/golang-orm/internal/schemaExtractor"
	"database/sql"
	"fmt"

	"github.com/dave/jennifer/jen"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	const file = "data/data.db"

	fmt.Println("Connecting to database...")
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic(err)
	}

	tables := schemaextractor.GetTableSchemas(db)
	for _, table := range tables {
		generateStruct(table)
	}

	defer db.Close()
}

func generateStruct(table schemaextractor.TableSchema) {
	fields := make([]jen.Code, 0)
	for _, column := range table.Columns {
		goType := schemaextractor.GetGoType(column.Type)
		field := jen.Id(title(column.Name))
		if column.Nullable {
			field = field.Op("*")
		}
		field = field.Add(getFieldType(goType))
		fields = append(fields, field)
	}

	f := jen.Type().Id(title(table.Name)).Struct(
		fields...,
	)

	fmt.Printf("%#v", f)
}

func title(s string) string {
	caser := cases.Title(language.English)
	return caser.String(s)
}

func getFieldType(gotype schemaextractor.GoType) jen.Code {
	switch gotype {
	case "int":
		return jen.Int()
	case "string":
		return jen.String()
	case "bool":
		return jen.Bool()
	}

	return jen.Empty()
}
