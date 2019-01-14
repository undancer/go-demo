package mongodb

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"log"
	"time"
)

const (
	url        = "mongodb://localhost:27017"
	database   = "go-database"
	collection = "go-collection"
	//username   = "mongo"
	//password   = "mongo"
)

func Main() {
	var (
		ctx    context.Context
		client *mongo.Client
		err    error
	)

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

	if client, err = mongo.Connect(ctx, url); err != nil {
		log.Println(err.Error())
	}
	_database := client.Database(database)
	_collection := _database.Collection(collection)
	var limit int64 = 10
	cur, err := _collection.Find(ctx, bson.M{}, &options.FindOptions{Limit: &limit})
	if err != nil {
		log.Println(err.Error())
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result bson.M
		cur.Decode(&result)
		fmt.Println(result)
	}

}
