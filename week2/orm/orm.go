package orm

import (
	"Homework/week2/errors"
)

type OrmHandler struct {
	//假装是个查询实例
	Name string
	Db string
	Table string
}

func (orm *OrmHandler) Query (sql string) (map[string]string, error) {
	//假装查询
	return make(map[string]string), errors.EmptyRow("test", "is a test")
}

func New(handlerName string, dbName string, tableName string) *OrmHandler {
	return &OrmHandler{
		Name: handlerName,
		Db: dbName,
		Table: tableName,
	}
}