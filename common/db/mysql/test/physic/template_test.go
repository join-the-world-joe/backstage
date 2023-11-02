package template

import (
	"backstage/diagnostic"
	"backstage/global/mysql"
	"fmt"
	"github.com/google/uuid"
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

func TestGetTableName(t *testing.T) {
	t.Log(GetTableNameList())
}

func TestAutoMigrate(t *testing.T) {
	diagnostic.SetupMySQL()
	for i := 1; i <= Mod; i++ {
		err := mysql.AutoMigrate(GetWhich(i), GetDbName(), GetTableName(), &Template{})
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestDropTable(t *testing.T) {
	diagnostic.SetupMySQL()
	for i := 1; i <= Mod; i++ {
		err := mysql.DropTable(GetWhich(i), GetDbName(), GetTableName())
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestInsert(t *testing.T) {
	diagnostic.SetupMySQL()
	for i := 1; i <= Mod; i++ {
		n, err := mysql.Insert(GetWhich(i), GetDbName(), GetTableName(), &Template{Column1: i})
		if err != nil {
			t.Fatal(err)
		}
		t.Log("n: ", n)
	}
}

func TestUpdate(t *testing.T) {
	diagnostic.SetupMySQL()
	for i := 1; i <= Mod; i++ {
		n, err := mysql.Update(GetWhich(i), GetDbName(), GetTableName(), map[string]interface{}{"column_2": uuid.New().String()}, fmt.Sprintf("column_1=%v", i))
		if err != nil {
			t.Fatal(err)
		}
		t.Log("n: ", n)
	}
}

func TestQuery(t *testing.T) {
	sql := "select column_2 from %s where column_1 = '%v' limit 1"
	diagnostic.SetupMySQL()
	for i := 1; i <= Mod; i++ {
		rows, err := mysql.Query(
			GetWhich(i),
			GetDbName(),
			fmt.Sprintf(sql, GetTableName(), i),
		)
		if err != nil {
			t.Fatal(err)
		}
		c2 := ""
		for rows.Next() {
			rows.Scan(&c2)
		}
		rows.Close()
		t.Log("c2: ", c2)
	}
}
