package instance

import (
	"sync"
)

var instance *single
var once sync.Once

type single struct {
	Name string
}

func GetInstance() *single {
	once.Do(func() {
		instance = &single{}
	})
	return instance
}
