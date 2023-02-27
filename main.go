package main

import (
	"fmt"

	"github.com/mars/go-cache/cache"
)

const (
	TABLE_NAME = "test_gy"
)

var cacheTable *cache.CacheTable

func Add(key string, value any) {
	cacheTable.Put(key, value)
}

func main() {
	fmt.Println("hello world")

	cache := cache.New()

	cacheTable = cache.Create(TABLE_NAME)

	Add("k1", "嘿嘿")

	res, exist := cacheTable.Get("k1")
	if !exist {
		fmt.Println("not found k1")
	} else {
		fmt.Printf("k1对应的value: %v\n", res)
	}
}
