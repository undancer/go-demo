package main

import (
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
)

type TestStruct struct {
	Name string
	ID   int32
}

func main() {
	fmt.Println("start")
	data, err := bson.Marshal(&TestStruct{Name: "Bob"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", data)

	value := TestStruct{}
	err2 := bson.Unmarshal(data, &value)
	if err2 != nil {
		panic(err)
	}
	fmt.Println("value:", value)

	mmap := bson.M{}
	err3 := bson.Unmarshal(data, &mmap)
	if err3 != nil {
		panic(err)
	}

	var m map[string]interface{} = mmap

	fmt.Println("mmap:", mmap)
	fmt.Println("m:", m)
	
}

func to(m map[string]interface{}) {
	fmt.Println(m)
}
