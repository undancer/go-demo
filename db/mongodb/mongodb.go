package mongodb

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"time"
)

type DB struct {
	db *mongo.Client
}

func NewDB(config map[string]string) *DB {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var (
		client *mongo.Client
		err    error
	)

	url := config["url"]

	if client, err = mongo.Connect(ctx, url); err != nil {
		fmt.Println(err)
	}

	return &DB{client}

}

func (this *DB) fetch(ctx context.Context, database, collection string, filter interface{}, opts ...*options.FindOptions) (results []interface{}) {

	cursor, err := this.db.Database(database).Collection(collection).Find(ctx, filter, opts...)
	if err != nil {
		fmt.Println(err)
	}

	for cursor.Next(nil) {
		var result bson.M
		cursor.Decode(&result)
		results = append(results, result)
	}
	return
}

func (this *DB) FetchAll(ctx context.Context, database, collection string, filter interface{}, opts ...*options.FindOptions) (results []interface{}) {
	return this.fetch(ctx, database, collection, filter, opts...)
}
