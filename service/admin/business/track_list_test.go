package business

import (
	"backstage/common/db/mgo/backend/track"
	"backstage/diagnostic"
	"backstage/global/config"
	"backstage/global/log"
	"backstage/global/mgo"
	"backstage/utils/timestamp"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"
)

func TestInsertTrackList(t *testing.T) {
	diagnostic.SetupMongoDB()
	diagnostic.SetupLogger()

	//location := "Asia/Shanghai"

	doc1 := &track.Model{
		Operator:   "doc1",
		Major:      "doc1",
		Minor:      "doc1",
		Permission: "doc1",
		Request:    "doc1",
		Response:   "doc1",
		Timestamp:  time.Now().Unix(),
	}
	doc2 := &track.Model{
		Operator:   "doc2",
		Major:      "doc2",
		Minor:      "doc2",
		Permission: "doc2",
		Request:    "doc2",
		Response:   "doc2",
		Timestamp:  timestamp.AddDate(time.Now().Unix(), 0, 0, -1),
	}
	doc3 := &track.Model{
		Operator:   "doc3",
		Major:      "doc3",
		Minor:      "doc3",
		Permission: "doc3",
		Request:    "doc3",
		Response:   "doc3",
		Timestamp:  timestamp.AddDate(time.Now().Unix(), 0, 0, -2),
	}

	docs := []interface{}{doc1, doc2, doc3}
	idList, err := mgo.InsertDocs(
		config.MongoConf(),
		context.Background(),
		track.GetWhich(), track.GetDBName(), track.GetTableName(),
		docs,
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ID List: ", idList)
}

func TestGetTrackByOperator(t *testing.T) {
	diagnostic.SetupMongoDB()
	diagnostic.SetupLogger()

	operator := "operator"
	value := "doc3"
	trackList := []*track.Model{}

	where := bson.M{operator: value}

	cur, err := mgo.Query(
		config.MongoConf(),
		context.Background(),
		track.GetWhich(),
		track.GetDBName(),
		track.GetTableName(),
		where,
	)
	if err != nil {
		t.Fatal(err)
	}

	for cur.Next(context.Background()) {
		doc := &track.Model{}
		err = cur.Decode(doc)
		if err != nil {
			log.Error("FetchTrackListOfCondition failure, err: ", err.Error())
			continue
		}
		trackList = append(trackList, doc)
	}

	bytes, err := json.Marshal(trackList)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("operator: ", operator)
	t.Log("track list: ", string(bytes))
}

func TestGetTrackListBetweenTimstampInterval(t *testing.T) {
	diagnostic.SetupMongoDB()
	diagnostic.SetupLogger()
	location := "Asia/Shanghai"
	year, month, day := timestamp.GetYearMonthDay()
	//begin := timestamp.GetBeginOfADay(year, month, day, location)
	begin := timestamp.GetBeginOfADay(year, month, 16, location)
	end := timestamp.GetEndOfADay(year, month, day, location)

	trackList := []*track.Model{}

	where := &bson.D{
		{"timestamp",
			bson.D{
				{"$lte", end},
				{"$gte", begin},
			},
		},
	}

	cur, err := mgo.Query(
		config.MongoConf(),
		context.Background(),
		track.GetWhich(),
		track.GetDBName(),
		track.GetTableName(),
		where,
	)
	if err != nil {
		t.Fatal(err)
	}

	for cur.Next(context.Background()) {
		doc := &track.Model{}
		err = cur.Decode(doc)
		if err != nil {
			log.Error("FetchTrackListOfCondition failure, err: ", err.Error())
			continue
		}
		trackList = append(trackList, doc)
	}

	bytes, err := json.Marshal(trackList)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("begin: ", timestamp.ToDateTimeString(begin))
	t.Log("end: ", timestamp.ToDateTimeString(end))
	t.Log("track list: ", string(bytes))
}
