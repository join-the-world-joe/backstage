package mgo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"
)

func TestRawConnection1(t *testing.T) {
	URI := "mongodb://root:123456@119.23.224.221:27001/?directConnection=true"
	client, err := connectToMongo(URI)
	if err != nil {
		t.Fatal(err)
	}
	err = client.Disconnect(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}

func TestConnection2(t *testing.T) {
	dbName := "backend"
	tblName := "track"
	doc := map[string]string{"key": "value"}
	clientOpts := options.Client().ApplyURI(
		"mongodb://root:123456@119.23.224.221:27001/?directConnection=true")
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		t.Fatal(err)
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		t.Fatal(err)
	}
	_ctx, err := client.Database(dbName).Collection(tblName).InsertOne(context.Background(), &doc)
	if err != nil {
		t.Fatal(err)
	}
	id, ok := _ctx.InsertedID.(primitive.ObjectID)
	if !ok {
		t.Fatal("!ok")
	}
	t.Log("id: ", id)
}
