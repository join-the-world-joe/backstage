package mysql

import (
	"testing"
)

var mysql_multiple_master_conf = `
[MySQL.HongKong.Master]
	Host = "192.168.130.128"
	Port = "23305"
	User = "root"
	Password = "123456"

[MySQL.HongKong1.Master]
	Host = "192.168.130.128"
	Port = "23305"
	User = "root"
	Password = "123456"

[MySQL.HongKong2.Master]
	Host = "192.168.130.128"
	Port = "23306"
	User = "root"
	Password = "123456"

[MySQL.HongKong3.Master]
	Host = "192.168.130.128"
	Port = "23307"
	User = "root"
	Password = "123456"
`

var mysql_master_slave_conf = `
# MySQL Master Slave Replication
[MySQL.Singapore.Master]
	Host = "192.168.130.128"
	Port = "13306"
	User = "root"
	Password = "123456"
[[MySQL.Singapore.Replicas]]
	Host = "192.168.130.128"
	Port = "13307"
	User = "root"
	Password = "123456"
`

func TestConnection1(t *testing.T) {
	host := "192.168.130.128"
	port := "23305"
	user := "root"
	password := "123456"
	dbName := "test"
	_, err := connectToDB(host, port, user, password, dbName)
	if err != nil {
		t.Fatal(err)
	}
}

//
//func TestConnection2(t *testing.T) {
//	which := "HongKong"
//	dbName := "test"
//	cf := &conf.MySQLConf{}
//	if err := toml.Unmarshal([]byte(mysql_multiple_master_conf), &cf); err != nil {
//		t.Fatal(err)
//	}
//	if _, err := GetDB(cf, which, dbName); err != nil {
//		t.Fatal(err)
//	}
//}
//
//func TestConnection3(t *testing.T) {
//	which := "Singapore"
//	dbName := "test"
//	cf := &conf.MySQLConf{}
//	if err := toml.Unmarshal([]byte(mysql_master_slave_conf), &cf); err != nil {
//		t.Fatal(err)
//	}
//	if _, err := GetDB(cf, which, dbName); err != nil {
//		t.Fatal(err)
//	}
//}
