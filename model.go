package jgomodel

import (
	"fmt"
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

	m := &Model{
		Ctx: ctx,
		Schema: schema,
		Table: table,
	}

	m.FullTableName = m.GetFullTableName()
	m.Fields, err = psql.GetFields(ctx, schema, table)

	if err != nil {
		return nil, err
	}

	return m, nil
}

// Validate the model
func (m *Model) isValid() error {
	return m.Ctx.GetValidator().Struct(m)
}

//
func (m *Model) GetFullTableName() string {
	return fmt.Sprintf(`"%s"."%s"`, m.Schema, m.Table)
}