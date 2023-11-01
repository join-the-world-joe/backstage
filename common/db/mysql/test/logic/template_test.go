package template

import (
	"fmt"
	"github.com/google/uuid"
	"go-micro-framework/diagnostic"
	"go-micro-framework/global/mysql"
	"testing"
)

func TestGetTableName(t *testing.T) {
	t.Log(GetTableNameList())
}

func TestAutoMigrate(t *testing.T) {
	diagnostic.SetupMySQL()
	for i := 1; i <= Mod; i++ {
		err := mysql.AutoMigrate(GetWhich(), GetDbName(), GetTableName(i), &Template{})
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestDropTable(t *testing.T) {
	diagnostic.SetupMySQL()
	for i := 1; i <= Mod; i++ {
		err := mysql.DropTable(GetWhich(), GetDbName(), GetTableName(i))
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestInsert(t *testing.T) {
	diagnostic.SetupMySQL()
	for i := 1; i <= Mod; i++ {
		n, err := mysql.Insert(GetWhich(), GetDbName(), GetTableName(i), &Template{Column1: i})
		if err != nil {
			t.Fatal(err)
		}
		t.Log("n: ", n)
	}
}

func TestUpdate(t *testing.T) {
	diagnostic.SetupMySQL()
	for i := 1; i <= Mod; i++ {
		n, err := mysql.Update(GetWhich(), GetDbName(), GetTableName(i), map[string]interface{}{"column_2": uuid.New().String()}, fmt.Sprintf("column_1=%v", i))
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
			GetWhich(),
			GetDbName(),
			fmt.Sprintf(sql, GetTableName(i), i),
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
