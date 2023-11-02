package template

import (
	"backstage/common/conf"
	"backstage/global/mgo"
	"context"
	"github.com/BurntSushi/toml"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

var mgo_server = `
[Mongo.Backend1]
	Servers = ["mongodb://192.168.130.128:37017"]
	User = "root"
	Password = "123456"
[Mongo.Backend2]
	Servers = ["mongodb://192.168.130.128:37018"]
	User = "root"
	Password = "123456"
[Mongo.Backend3]
	Servers = ["mongodb://192.168.130.128:37019"]
	User = "root"
	Password = "123456"
`

var cf = func() *conf.MongoConf {
	cf := &conf.MongoConf{}
	if err := toml.Unmarshal([]byte(mgo_server), &cf); err != nil {
		panic(err)
	}
	return cf
}()

func TestInsert(t *testing.T) {
	for i := 1; i <= Mod; i++ {
		id, err := mgo.InsertDoc(
			cf,
			context.Background(),
			GetWhich(i),
			GetDBName(),
			GetTableName(),
			&Template{Field1: i, Field2: uuid.New().String()},
		)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("id: ", id)
	}
}

func TestInsertDocs(t *testing.T) {
	doc1 := &Template{Field2: uuid.New().String()}
	doc2 := &Template{Field2: uuid.New().String()}
	doc3 := &Template{Field2: uuid.New().String()}
	docs := []interface{}{doc1, doc2, doc3}
	for i := 1; i <= Mod; i++ {
		objIdList, err := mgo.InsertDocs(
			cf,
			context.Background(),
			GetWhich(i),
			GetDBName(),
			GetTableName(),
			docs,
		)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("objIdList: ", objIdList)
	}
}

func TestDeleteDoc(t *testing.T) {
	id := 3
	field2 := "06dde6ba-b264-4106-bfd9-5efcafbdea60"

	where := &bson.M{
		"field_2": bson.M{"$eq": field2},
	}

	affected, err := mgo.DeleteDocs(
		cf,
		context.Background(),
		GetWhich(id),
		GetDBName(),
		GetTableName(),
		where,
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("affected: ", affected)
}

func TestUpdateDoc(t *testing.T) {
	id := 3
	field2 := "50284621-ea54-4913-b5b3-8b9062b8a9fd"
	where := &bson.D{
		{"field_2",
			bson.D{
				{"$eq", field2},
			},
		},
	}

	fields := &bson.D{
		{"$set",
			bson.D{
				{"field_2", uuid.New().String()},
			},
		},
	}
	objId, err := mgo.UpdateDoc(
		cf,
		context.Background(),
		GetWhich(id),
		GetDBName(),
		GetTableName(),
		where,
		fields,
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("objId: ", objId)
}

func TestUpdateDocs(t *testing.T) {
	id := 3
	con1 := "5f5a7d0d-e19e-4390-acd3-45505a32268d"
	con2 := "a33fc1a1-34bb-40eb-8931-761ccb074864"
	con3 := "637e7da8-78ff-492f-8500-47e534570c82"
	where := &bson.D{
		{"field_2",
			bson.D{
				{"$in", bson.A{con1, con2, con3}},
			},
		},
	}

	fields := &bson.D{
		{"$set",
			bson.D{
				{"field_2", uuid.New().String()},
			},
		},
	}
	n1, n2, err := mgo.UpdateDocs(
		cf,
		context.Background(),
		GetWhich(id),
		GetDBName(),
		GetTableName(),
		where,
		fields,
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("n1: ", n1)
	t.Log("n2: ", n2)
}

func TestQuery(t *testing.T) {
	id := 3
	con1 := "b9c85f25-387c-470c-a74b-4882dd09dc8a"
	con2 := "a33fc1a1-34bb-40eb-8931-761ccb074864"
	con3 := "637e7da8-78ff-492f-8500-47e534570c82"

	where := &bson.D{
		{"field_2",
			bson.D{
				{"$in", bson.A{con1, con2, con3}},
			},
		},
	}

	cur, err := mgo.Query(
		cf,
		context.Background(),
		GetWhich(id),
		GetDBName(),
		GetTableName(),
		where,
	)
	if err != nil {
		t.Fatal(err)
	}

	for cur.Next(context.Background()) {
		_doc := &Template{}
		cur.Decode(_doc)
		t.Log("doc: ", _doc)
	}
}

func TestFind(t *testing.T) {
	id := 3
	con1 := "b9c85f25-387c-470c-a74b-4882dd09dc8a"
	con2 := "a33fc1a1-34bb-40eb-8931-761ccb074864"
	con3 := "637e7da8-78ff-492f-8500-47e534570c82"

	where := &bson.D{
		{"field_2",
			bson.D{
				{"$in", bson.A{con1, con2, con3}},
			},
		},
	}

	sort := &bson.M{"field_1": 1, "field_2": -1}

	cur, err := mgo.Find(
		cf,
		context.Background(),
		GetWhich(id),
		GetDBName(),
		GetTableName(),
		where,
		sort,
		10,
		0,
	)
	if err != nil {
		t.Fatal(err)
	}
	for cur.Next(context.Background()) {
		_doc := &Template{}
		cur.Decode(_doc)
		t.Log("doc: ", _doc)
	}
}

func TestCreateIndex(t *testing.T) {
	id := 3
	index := mongo.IndexModel{
		Keys: bson.D{
			{"field_1", 1},
			{"field_2", -1},
		},
	}
	name, err := mgo.CreateIndex(
		cf,
		context.Background(),
		GetWhich(id),
		GetDBName(),
		GetTableName(),
		index,
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("index name: ", name)
}

func TestListIndex(t *testing.T) {
	id := 3
	nameList, err := mgo.ListIndex(
		cf,
		context.Background(),
		GetWhich(id),
		GetDBName(),
		GetTableName(),
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("name list: ", nameList)
}

func TestDropIndex(t *testing.T) {
	id := 3
	indexName := "field_1_1_field_2_-1"
	err := mgo.DropIndex(
		cf,
		context.Background(),
		GetWhich(id),
		GetDBName(),
		GetTableName(),
		indexName,
	)
	if err != nil {
		t.Fatal(err)
	}
}
