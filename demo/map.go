package main

import "fmt"

func main() {

	var kvMap map[string]string = make(map[string]string)

	kvMap["China"] = "Beijing"

	for k, v := range kvMap {
		fmt.Println(k, v)
	}

	value, ok := kvMap["中国"]
	if ok {
		fmt.Println("存在", value)
	} else {
		fmt.Println("不存在")
	}

	fmt.Println(kvMap)

}
