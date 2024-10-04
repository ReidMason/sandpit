package main

import (
	schemaextractor "ReidMason/golang-orm/internal/schemaExtractor"
	"database/sql"
	"fmt"

	"github.com/dave/jennifer/jen"
	_ "github.com/mattn/go-sqlite3"
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
		generateStructs(table)
	}

	defer db.Close()
}

func generateStructs(table schemaextractor.TableSchema) {
	f := jen.NewFile("output")

	fields := make([]jen.Code, 0)
	for _, column := range table.Columns {
		goType := schemaextractor.GetGoType(column.Type)
		field := jen.Id(column.Name).Add(getFieldType(goType))
		fields = append(fields, field)
	}

	f.Type().Id(table.Name).Struct(
		fields...,
	)

	fmt.Printf("%#v", f)
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
