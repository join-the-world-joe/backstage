package mgo

import (
	"backstage/common/conf"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient(cf *conf.MongoConf, which string) (*mongo.Client, error) {
	_, exist := cf.Mongo[which]
	if !exist {
		return nil, errors.New(fmt.Sprintf("cann't find db info of %s", which))
	}
	return _getMongoClient(cf, which)
}

func InsertDoc(cf *conf.MongoConf, ctx context.Context, which, dbName, tblName string, doc interface{}) (string, error) {
	_client, err := _getMongoClient(cf, which)
	if err != nil {
		return "", err
	}
	_ctx, err := _client.Database(dbName).Collection(tblName).InsertOne(ctx, doc)
	if err != nil {
		return "", err
	}
	id, ok := _ctx.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("ctx.InsertedID is not a ObjectId")
	}
	return id.String(), nil
}

func InsertDocs(cf *conf.MongoConf, ctx context.Context, which, dbName, tblName string, docs []interface{}) ([]string, error) {
	_client, err := _getMongoClient(cf, which)
	if err != nil {
		return nil, err
	}
	_ctx, err := _client.Database(dbName).Collection(tblName).InsertMany(ctx, docs)
	if err != nil {
		return nil, err
	}
	if len(_ctx.InsertedIDs) == 0 {
		return nil, fmt.Errorf("len(ctx.InsertedIDs) == 0")
	}
	ids := make([]string, 0)
	for k, v := range _ctx.InsertedIDs {
		if id, ok := v.(primitive.ObjectID); ok {
			ids = append(ids, id.String())
		} else {
			return nil, fmt.Errorf("%v is not a primitive.ObjectID", k)
		}
	}
	return ids, nil
}

func DeleteDoc(cf *conf.MongoConf, ctx context.Context, which, dbName, tblName string, where *bson.M) (int64, error) {
	_client, err := _getMongoClient(cf, which)
	if err != nil {
		return 0, err
	}

	ret, err := _client.Database(dbName).Collection(tblName).DeleteOne(ctx, where)
	if err != nil {
		return 0, err
	}

	return ret.DeletedCount, nil
}

func DeleteDocs(cf *conf.MongoConf, ctx context.Context, which, dbName, tblName string, where *bson.M) (int64, error) {
	_client, err := _getMongoClient(cf, which)
	if err != nil {
		return 0, err
	}

	ret, err := _client.Database(dbName).Collection(tblName).DeleteMany(ctx, where)
	if err != nil {
		return 0, err
	}

	return ret.DeletedCount, nil
}

func UpdateDoc(cf *conf.MongoConf, ctx context.Context, which, dbName, tblName string, where interface{}, fields interface{}) (string, error) {
	_client, err := _getMongoClient(cf, which)
	if err != nil {
		return "", err
	}

	if where == nil {
		fmt.Println("where == nil")
	}

	ret := _client.Database(dbName).Collection(tblName).FindOneAndUpdate(ctx, where, fields)
	if ret.Err() != nil {
		return "", ret.Err()
	}

	raw, err := ret.DecodeBytes()
	if err != nil {
		return "", err
	}

	val, err := raw.LookupErr("_id")
	if err != nil {
		return "", err
	}

	return val.ObjectID().Hex(), nil
}

func UpdateDocs(cf *conf.MongoConf, ctx context.Context, which, dbName, tblName string, where interface{}, fields interface{}) (int64, int64, error) {
	_client, err := _getMongoClient(cf, which)
	if err != nil {
		return 0, 0, err
	}

	ret, err := _client.Database(dbName).Collection(tblName).UpdateMany(ctx, where, fields)
	if err != nil {
		return 0, 0, err
	}

	return ret.MatchedCount, ret.ModifiedCount, nil
}

func Query(cf *conf.MongoConf, ctx context.Context, which, dbName, tblName string, where interface{}) (*mongo.Cursor, error) {
	_client, err := _getMongoClient(cf, which)
	if err != nil {
		return nil, err
	}

	return _client.Database(dbName).Collection(tblName).Find(ctx, where)
}

func Find(cf *conf.MongoConf, ctx context.Context, which, dbName, tblName string, where interface{}, sort interface{}, pageSize, skip int64) (*mongo.Cursor, error) {
	_client, err := _getMongoClient(cf, which)
	if err != nil {
		return nil, err
	}

	opts := []*options.FindOptions{}

	if sort != nil {
		opts = append(opts, options.Find().SetSort(sort))
	}

	if pageSize > 0 {
		opts = append(opts, options.Find().SetLimit(pageSize))
	}

	if skip > 0 {
		opts = append(opts, options.Find().SetSkip(skip))
	}

	return _client.Database(dbName).Collection(tblName).Find(ctx, where, opts...)
}

func ListIndex(cf *conf.MongoConf, ctx context.Context, which, dbName, tblName string) ([]string, error) {
	indexName := []string{}
	_client, err := _getMongoClient(cf, which)
	if err != nil {
		return nil, err
	}

	cur, err := _client.Database(dbName).Collection(tblName).Indexes().List(ctx)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		indexName = append(indexName, cur.Current.Lookup("name").StringValue())
	}

	return indexName, nil
}

func CreateIndex(cf *conf.MongoConf, ctx context.Context, which, dbName, tblName string, index mongo.IndexModel) (string, error) {
	_client, err := _getMongoClient(cf, which)
	if err != nil {
		return "", err
	}

	return _client.Database(dbName).Collection(tblName).Indexes().CreateOne(ctx, index)
}

func DropIndex(cf *conf.MongoConf, ctx context.Context, which, dbName, tblName, indexName string) error {
	_client, err := _getMongoClient(cf, which)
	if err != nil {
		return err
	}

	_, err = _client.Database(dbName).Collection(tblName).Indexes().DropOne(ctx, indexName)

	return err
}
