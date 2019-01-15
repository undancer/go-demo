package main

import (
	"fmt"
	"github.com/koding/cache"
	"time"
)

func main() {
	// create a cache with 2 second TTL
	c := cache.NewMemoryWithTTL(2 * time.Second)
	// start garbage collection for expired keys
	c.StartGC(time.Millisecond * 10)
	// set item
	err := c.Set("test_key", "test_data")
	// get item
	data, err := c.Get("test_key")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}
