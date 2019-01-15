package main

import (
	"fmt"
	"log"
	"net/http"
)

func handle(writer http.ResponseWriter, request *http.Request) {
	for k, v := range request.Header {
		fmt.Println("k = ", k, ", v = ", v)
	}
	fmt.Println("url = ", request.URL)
	fmt.Println("rui = ", request.RequestURI)
}

func main() {
	http.HandleFunc("/", handle)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Println(err.Error())
	}
}
