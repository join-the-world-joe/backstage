package mysql

import (
	"backstage/global/config"
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// for tables

func GetDB(which, dbName string) (*gorm.DB, error) {
	_, exist := config.MySQLConf().MySQL[which]
	if !exist {
		return nil, errors.New(fmt.Sprintf("cann't find db info of %s", which))
	}
	return _getDB(which, dbName)
}

func AutoMigrate(which, dbName, tblName string, model interface{}) error {
	_db, err := GetDB(which, dbName)
	if err != nil {
		return err
	}
	return _db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 COLLATE=utf8_bin").Table(tblName).AutoMigrate(model)
}

func DropTable(which, dbName, tblName string) error {
	_db, err := GetDB(which, dbName)
	if err != nil {
		return err
	}
	_sql := fmt.Sprintf("DROP TABLE IF EXISTS %s", tblName)
	_db.Exec(_sql)
	return nil
}

// for records
func Insert(which, dbName, tblName string, p interface{}) (interface{}, error) {
	_db, err := GetDB(which, dbName)
	if err != nil {
		return nil, err
	}
	return p, _db.Table(tblName).Create(p).Error
}

func Update(which, dbName, tblName string, data map[string]interface{}, where string) (int64, error) {
	_db, err := GetDB(which, dbName)
	if err != nil {
		return 0, err
	}
	result := _db.Table(tblName).Where(where).Updates(data)
	return result.RowsAffected, result.Error
}

func Query(which, dbName, sql string) (*sql.Rows, error) {
	_db, err := GetDB(which, dbName)
	if err != nil {
		return nil, err
	}
	return _db.Raw(sql).Rows()
}

func Delete(which, dbName, tblName string, sql string) error {
	_db, err := GetDB(which, dbName)
	if err != nil {
		return err
	}
	return _db.Table(tblName).Exec(sql).Error
}

// for raw sql
//func Exec(db *gorm.DB, tblName, sql string) error {
//	return db.Table(tblName).Exec(sql).Error
//}
