package track

import (
	"backstage/diagnostic"
	"backstage/global/config"
	"backstage/global/mgo"
	"context"
	"fmt"
	"testing"
	"time"
)

func TestInsertModel(t *testing.T) {
	diagnostic.SetupMongoDB()
	diagnostic.SetupLogger()
	id, err := mgo.InsertDoc(
		config.MongoConf(),
		context.Background(),
		GetWhich(),
		GetDBName(),
		GetTableName(),
		&Model{
			Operator:   "山水有相逢",
			Major:      "major",
			Minor:      "minor",
			Permission: "打PP",
			Request:    "request of 打PP",
			Response:   "response of 打PP",
			Timestamp:  time.Now().Unix(),
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("id: ", id)
}

func TestTimeUnixToDate(t *testing.T) {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	fmt.Println(tm.Format("2006-01-02 15:04:05"))
}

func TestBeginningOfADay(t *testing.T) {
	location, _ := time.LoadLocation("Asia/Kolkata")
	timestamp := time.Now().Unix()
	year, moon, day := time.Unix(timestamp, 0).Date()
	t.Log(time.Date(year, moon, day, 0, 0, 0, 0, location))

	timestamp = time.Now().AddDate(0, 0, -1).Unix()
	year, moon, day = time.Unix(timestamp, 0).Date()
	t.Log(time.Date(year, moon, day, 0, 0, 0, 0, location))
}
