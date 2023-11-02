package template

import (
	"backstage/diagnostic"
	"backstage/global/mysql"
	"fmt"
	"github.com/google/uuid"
	"testing"
)

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

func TestAutoMigrate(t *testing.T) {
	diagnostic.SetupMySQL()
	err := mysql.AutoMigrate(GetWhich(), GetDbName(), GetTableName(), &Template{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDropTable(t *testing.T) {
	diagnostic.SetupMySQL()
	err := mysql.DropTable(GetWhich(), GetDbName(), GetTableName())
	if err != nil {
		t.Fatal(err)
	}
}

func TestInsert(t *testing.T) {
	diagnostic.SetupMySQL()
	n, err := mysql.Insert(GetWhich(), GetDbName(), GetTableName(), &Template{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("n: ", n)

}

func TestUpdate(t *testing.T) {
	column1 := 1
	where := fmt.Sprintf("column_1=%v", column1)
	diagnostic.SetupMySQL()
	n, err := mysql.Update(
		GetWhich(),
		GetDbName(),
		GetTableName(),
		map[string]interface{}{
			"column_2": uuid.New().String()},
		where,
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("n: ", n)
}

func TestQuery(t *testing.T) {
	sql := "select column_2 from %s where column_1 = '%v' limit 1"
	diagnostic.SetupMySQL()
	rows, err := mysql.Query(
		GetWhich(),
		GetDbName(),
		fmt.Sprintf(sql, GetTableName(), 1),
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
