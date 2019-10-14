package model

import (
	"fmt"
	"regexp"
	"strings"
	"github.com/jschneider98/jgoweb"
	"github.com/jschneider98/jgoweb/db/psql"
)

type Model struct {
	Schema string `json:"schema"`
	Table string `json:"table"`
	Fields []psql.Field `json:"-"`
	Ctx jgoweb.ContextInterface `json:"-"`
}

