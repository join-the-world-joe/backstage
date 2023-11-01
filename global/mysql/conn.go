package mysql

import (
	"backstage/global/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"sync"
)

var g_lock sync.Mutex

func _getDB(which, dbName string) (*gorm.DB, error) {
	fullName := which + "-" + dbName

	_db, err := load(fullName)
	if err == nil { // created
		return _db, nil
	}

	g_lock.Lock()
	defer g_lock.Unlock()

	_db, err = load(fullName)
	if err == nil { // created
		return _db, nil
	}

	_db, err = connectToDB(
		config.MySQLConf().MySQL[which].Master.Host,
		config.MySQLConf().MySQL[which].Master.Port,
		config.MySQLConf().MySQL[which].Master.User,
		config.MySQLConf().MySQL[which].Master.Password,
		dbName,
	)
	if err != nil {
		return nil, err
	}
	store(fullName, _db)

	if len(config.MySQLConf().MySQL[which].Sources) > 0 || len(config.MySQLConf().MySQL[which].Replicas) > 0 {
		var sources []gorm.Dialector
		var replicas []gorm.Dialector
		resolver := dbresolver.Config{
			Policy: dbresolver.RandomPolicy{},
		}
		for _, c := range config.MySQLConf().MySQL[which].Sources {
			sources = append(sources, mysql.Open(fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&allowNativePasswords=true",
				c.User, c.Password, c.Host, c.Port, dbName)))
		}
		for _, c := range config.MySQLConf().MySQL[which].Replicas {
			replicas = append(replicas, mysql.Open(fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&allowNativePasswords=true",
				c.User, c.Password, c.Host, c.Port, dbName)))
		}
		if len(sources) != 0 {
			resolver.Sources = sources
		}
		if len(replicas) != 0 {
			resolver.Replicas = replicas
		}
		err = _db.Use(dbresolver.Register(resolver))
		if err != nil {
			return nil, err
		}
	}

	return _db, nil
}

func connectToDB(host, port, user, passwd, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&allowNativePasswords=true",
		user, passwd, host, port, dbName)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info), AllowGlobalUpdate: true})
}
