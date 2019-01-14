package mongodb

import (
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func Main() {
	client, _ := mongo.NewClient("mongodb://localhost:27017")
	fmt.Println(client)
}
