package main

import (
	"fmt"
	"testing"
)

func TestCreateCache(t *testing.T) {
	cache := New(2)           // nil
	cache.Set(1, 10)          // nil, linked list: [1:10]
	cache.Set(2, 20)          // nil, linked list: [2:20, 1:10]
	fmt.Println(cache.Get(1)) // 10, linked list: [1:10, 2:20]
	cache.Set(3, 30)          // nil, linked list: [3:30, 1:10]
	fmt.Println(cache.Get(2)) // -1, linked list: [3:30, 1:10]
	cache.Set(4, 40)          // nil, linked list: [4:40, 3:30]
	fmt.Println(cache.Get(1)) // -1, linked list: [4:40, 3:30]
	fmt.Println(cache.Get(3)) // 30, linked list: [3:30, 4:40]
}
