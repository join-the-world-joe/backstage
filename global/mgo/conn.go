package mgo

import (
	"backstage/common/conf"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
	"time"
)

var g_lock sync.Mutex

const (
	DefaultConnectionTimeout = 10 * time.Second
)

func _getMongoClient(cf *conf.MongoConf, which string) (*mongo.Client, error) {
	_client, err := load(which)
	if err == nil { // created
		return _client, nil
	}

	g_lock.Lock()
	defer g_lock.Unlock()

	_client, err = load(which)
	if err == nil { // created
		return _client, nil
	}

	_client, err = connectToMongo(
		cf.Mongo[which].URI,
	)
	if err != nil {
		return nil, err
	}

	store(which, _client)
	return _client, nil
}

func connectToMongo(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultConnectionTimeout)
	defer cancel()

	copts := []*options.ClientOptions{}

	copts = append(copts, options.Client().ApplyURI(uri))

	_client, err := mongo.Connect(ctx, copts...)
	if err != nil {
		return nil, err
	}

	return _client, _client.Ping(ctx, readpref.Primary())
}
