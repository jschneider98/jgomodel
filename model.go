package jgomodel

import (
	"fmt"
	"strings"
	"github.com/jschneider98/jgoweb"
	"github.com/jschneider98/jgoweb/db/psql"
)

type Model struct {
	Schema string `json:"schema"`
	Table string `json:"table"`
	FullTableName string `json:"full_table_name"`
	Fields []psql.Field `json:"-"`
	Ctx jgoweb.ContextInterface `json:"-"`
}

//
func NewModel(ctx jgoweb.ContextInterface, schema string, table string) (*Model, error) {
	var err error

	m := &Model{ Ctx: ctx }
	err = m.SetMetaData(schema, table)

	if err != nil {
		return nil, err
	}

	return m, nil
}

//
func (m *Model) SetMetaData(schema string, table string) error {
	var err error

	m.Schema = schema
	m.Table = table

	m.FullTableName = m.GetFullTableName()
	m.Fields, err = psql.GetFields(m.Ctx, schema, table)


	return err
}

//
func (m *Model) GetInsertQuery() string {
	var dbCols []string
	var colList string
	var placeHolders []string
	var colCount int

	// 
	for key := range m.Fields {
		if (m.Fields[key].DbFieldName != "id" && m.Fields[key].DbFieldName != "created_at" && m.Fields[key].DbFieldName != "updated_at") {
			colCount++
			// (account_id, units, ...)
			dbCols = append(dbCols, m.Fields[key].DbFieldName)
			// ($1, $2, ...)
			placeHolders = append(placeHolders, fmt.Sprintf("$%d", colCount))
		}
	}

	colList = strings.Join(dbCols, ",\n\t")

	query := fmt.Sprintf("INSERT INTO\n%s (%s)\nVALUES (%s)\nRETURNING id\n", m.FullTableName, colList, strings.Join(placeHolders, ","))

	return query
}

//
func (m *Model) GetFullTableName() string {
	return fmt.Sprintf("%s.%s", m.Schema, m.Table)
}