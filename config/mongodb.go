package config

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/undancer/go-demo/db/mongodb"
	"time"
)

func init() {

	fmt.Println("初始化mongodb")
	configMongodb()
}

func configMongodb() {

	config := make(map[string]string)
	config["url"] = "mongodb://localhost:27017"

	db := mongodb.NewDB(config)

	var limit int64 = 10

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	results := db.FetchAll(ctx, "go-database", "go-collection", bson.M{}, &options.FindOptions{Limit: &limit})

	for _, result := range results {
		r := result.(bson.M)

		fmt.Println(r)
	}

}
