package main

import (
	"flag"
	"fmt"
	"github.com/petermattis/goid"
	"strings"
	"sync"

	"github.com/chrislusf/glow/flow"
)

type SafeMap struct {
	sync.RWMutex
	Map map[int64]string
}

func newSafeMap() *SafeMap {
	m := new(SafeMap)
	m.Map = make(map[int64]string)
	return m
}

func (sm *SafeMap) read(key int64) *string {
	sm.RLock()
	value := sm.Map[key]
	sm.RUnlock()
	return &value
}

func (sm *SafeMap) write(key int64, value string) {
	sm.Lock()
	sm.Map[key] = value
	sm.Unlock()
}

func main() {
	flag.Parse()

	var m = newSafeMap()

	flow.New().TextFile(
		"/etc/passwd", 3,
	).Filter(func(line string) bool {
		println("f1", goid.Get())
		m.write(goid.Get(), "f1")
		return !strings.HasPrefix(line, "#")
	}).Map(func(line string, ch chan string) {
		println("m1", goid.Get())
		m.write(goid.Get(), "m1")
		for _, token := range strings.Split(line, ":") {
			ch <- token
		}
	}).Map(func(key string) int {
		println("m2", goid.Get())
		m.write(goid.Get(), "m2")
		return 1
	}).Reduce(func(x int, y int) int {
		println("r1", goid.Get())
		m.write(goid.Get(), "r1")
		return x + y
	}).Map(func(x int) {
		println("r2", goid.Get())
		m.write(goid.Get(), "r2")
		println("count:", x)
	}).Run()

	fmt.Println()

	mm := make(map[string]uint)

	m.RLock()
	for _, v := range m.Map {
		if _, ok := mm[v]; ok {
			mm[v]++
		} else {
			mm[v] = 1
		}
	}
	m.RUnlock()

	fmt.Println(mm)

}
