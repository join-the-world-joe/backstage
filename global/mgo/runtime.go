package mgo

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

var g_mgo_map sync.Map // it holds all connected mongo

func store(unique string, client *mongo.Client) {
	g_mgo_map.Store(unique, client)
}

func load(unique string) (*mongo.Client, error) {
	value, ok := g_mgo_map.Load(unique)
	if ok {
		return value.(*mongo.Client), nil
	}
	return nil, errors.New(fmt.Sprintf("%s doesn't exist", unique))
}

func Dump() {
	a := g_mgo_map
	fmt.Println("a = ", a)
}