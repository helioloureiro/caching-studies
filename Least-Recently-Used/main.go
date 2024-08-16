package main

// src: https://girai.dev/blog/lru-cache-implementation-in-go/

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println("Implementing simple LFI caching")
}

type List struct {
	root Element // sentinel list element, only &root, root.prev, and root.next are used
	len  int     // current list length excluding (this) sentinel element
}

type Element struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next, prev *Element

	// The list to which this element belongs.
	list *List

	// The value stored with this element.
	Value interface{}
}

// Pair is the value of a list.List.Element.
type Pair struct {
	key   int
	value int
}

type LRUCache struct {
	cap int                   // capacity
	l   *list.List            // doubly linked list
	m   map[int]*list.Element // hash table for list.List.Element existence check
}

// Constructor initializes a new LRUCache.
func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
		l:   new(list.List),
		m:   make(map[int]*list.Element, capacity),
	}
}

// Get a list node from the hash map.
func (c *LRUCache) Get(key int) int {
	// check if list node exists
	if node, ok := c.m[key]; ok {
		val := node.Value.(*list.Element).Value.(Pair).value
		// move node to front
		c.l.MoveToFront(node)
		return val
	}
	return -1
}

// Put key and value in the LRUCache
func (c *LRUCache) Put(key int, value int) {
	// check if list node exists
	if node, ok := c.m[key]; ok {
		// move the node to front
		c.l.MoveToFront(node)
		// update the value of a list node
		node.Value.(*list.Element).Value = Pair{key: key, value: value}
	} else {
		// delete the last list node if the list is full
		if c.l.Len() == c.cap {
			// get the key that we want to delete
			idx := c.l.Back().Value.(*list.Element).Value.(Pair).key
			// delete the node pointer in the hash map by key
			delete(c.m, idx)
			// remove the last list node
			c.l.Remove(c.l.Back())
		}
		// initialize a list node
		node := &list.Element{
			Value: Pair{
				key:   key,
				value: value,
			},
		}
		// push the new list node into the list
		ptr := c.l.PushFront(node)
		// save the node pointer in the hash map
		c.m[key] = ptr
	}
}
