package dao

import (
	"Homework/week2/orm"
	"errors"
)

func getUserBySql(sql string) map[string]string {
	ormHandler := orm.New("article", "test", "article")
	if row, err := ormHandler.Query("select title from article where id = 1"); errors.Is(err, ) {
		return make(map[string]string)
	} else {
		return row
	}
}


