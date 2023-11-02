package advertisement

import (
	"backstage/diagnostic"
	"backstage/global/mysql"
	"fmt"
	"github.com/spf13/cast"
	"testing"
)

func TestAutoMigrate(t *testing.T) {
	diagnostic.SetupMySQL()
	err := mysql.AutoMigrate(GetWhich(), GetDbName(), GetTableName(), &Model{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestInsertModel(t *testing.T) {
	n := 20
	title := "title"
	origin := "origin"
	url := ""
	diagnostic.SetupMySQL()
	for i := 0; i < n; i++ {
		temp, err := InsertModel(&Model{
			Title:         fmt.Sprintf("%s%d", title, i+1),
			PlaceOFOrigin: fmt.Sprintf("%s%d", origin, i+1),
			Url:           fmt.Sprintf("%s%d", url, i+1),
			SellingPrice:  200,
			Visible:       1,
			Status:        0,
			Stock:         5,
			ProductId:     cast.ToInt64(i + 1),
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Log(temp)
	}

}
