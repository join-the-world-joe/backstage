package mysql

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"sync"
)

var g_db_map sync.Map // it holds all connected db

func store(unique string, db *gorm.DB) {
	g_db_map.Store(unique, db)
}

func load(unique string) (*gorm.DB, error) {
	value, ok := g_db_map.Load(unique)
	if ok {
		return value.(*gorm.DB), nil
	}
	return nil, errors.New(fmt.Sprintf("%s doesn't exist", unique))
}

func DumpMySQL() {
	a := g_db_map
	fmt.Println("a = ", a)
}
