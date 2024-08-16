package main

import (
	"fmt"
	"testing"
)

func TestCreateCache(t *testing.T) {
	obj := Constructor(2)   // nil
	obj.Put(1, 10)          // nil, linked list: [1:10]
	obj.Put(2, 20)          // nil, linked list: [2:20, 1:10]
	fmt.Println(obj.Get(1)) // 10, linked list: [1:10, 2:20]
	obj.Put(3, 30)          // nil, linked list: [3:30, 1:10]
	fmt.Println(obj.Get(2)) // -1, linked list: [3:30, 1:10]
	obj.Put(4, 40)          // nil, linked list: [4:40, 3:30]
	fmt.Println(obj.Get(1)) // -1, linked list: [4:40, 3:30]
	fmt.Println(obj.Get(3)) // 30, linked list: [3:30, 4:40]
}
