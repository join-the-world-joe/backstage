package mgo

import (
	"context"
	"go-micro-framework/common/conf"
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
		cf.Mongo[which].Servers,
		cf.Mongo[which].User,
		cf.Mongo[which].Password,
	)
	if err != nil {
		return nil, err
	}

	store(which, _client)
	return _client, nil
}

func connectToMongo(servers []string, user, password string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultConnectionTimeout)
	defer cancel()

	url := ""
	for k, v := range servers {
		if k > 0 {
			url = url + ","
		}
		url = url + v
	}

	//fmt.Println("url: ", url)

	copts := []*options.ClientOptions{}

	copts = append(copts, options.Client().ApplyURI(url))

	if user != "" && password != "" {
		copts = append(copts, options.Client().SetAuth(options.Credential{
			Username: user,
			Password: password,
		}))
	}

	_client, err := mongo.Connect(ctx, copts...)
	if err != nil {
		return nil, err
	}

	return _client, _client.Ping(ctx, readpref.Primary())
}
