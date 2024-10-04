package schemaextractor

import (
	"database/sql"
	"fmt"
)

type rawTableSchema struct {
	name      string `field:"name"`
	sqlType   string `field:"type"`
	tableName string `field:"tbl_name"`
	sql       string `field:"sql"`
	rootpage  int    `field:"rootpage"`
}

type TableSchema struct {
	Name      string
	Type      string
	TableName string
	SQL       string
	Columns   []ColumnSchema
	RootPage  int
}

func (t rawTableSchema) toTableSchema() TableSchema {
	return TableSchema{
		Name:      t.name,
		Type:      t.sqlType,
		TableName: t.tableName,
		SQL:       t.sql,
		RootPage:  t.rootpage,
		Columns:   make([]ColumnSchema, 0),
	}
}

func GetTableSchemas(db *sql.DB) []TableSchema {
	rows, err := db.Query("SELECT * FROM sqlite_master WHERE type='table';")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	tableSchemas := make([]TableSchema, 0)
	for rows.Next() {
		var tableData rawTableSchema
		err = rows.Scan(&tableData.sqlType, &tableData.name, &tableData.tableName, &tableData.rootpage, &tableData.sql)
		if err != nil {
			panic(err)
		}

		tableSchema := tableData.toTableSchema()
		tableSchema.Columns = getTableColumns(db, tableSchema.TableName)

		tableSchemas = append(tableSchemas, tableSchema)
	}

	return tableSchemas
}

type rawcolumnSchema struct {
	defaultValue *string `field:"dflt_value"`
	name         string  `field:"name"`
	columnType   string  `field:"type"`
	cid          int     `field:"cid"`
	notNull      int     `field:"notnull"`
	primaryKey   int     `field:"pk"`
}

type ColumnSchema struct {
	DefaultValue *string
	Name         string
	Type         string
	Cid          int
	Nullable     bool
	PrimaryKey   bool
}

func (c rawcolumnSchema) toColumnSchema() ColumnSchema {
	return ColumnSchema{
		DefaultValue: c.defaultValue,
		Name:         c.name,
		Type:         c.columnType,
		Cid:          c.cid,
		Nullable:     c.notNull == 0,
		PrimaryKey:   c.primaryKey == 1,
	}
}

func getTableColumns(db *sql.DB, tableName string) []ColumnSchema {
	rows, err := db.Query(fmt.Sprintf("PRAGMA table_info(%s);", tableName))
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	columnSchemas := make([]ColumnSchema, 0)
	for rows.Next() {
		var columnData rawcolumnSchema
		err = rows.Scan(&columnData.cid, &columnData.name, &columnData.columnType, &columnData.notNull, &columnData.defaultValue, &columnData.primaryKey)
		if err != nil {
			panic(err)
		}

		columnSchemas = append(columnSchemas, columnData.toColumnSchema())
	}

	return columnSchemas
}
