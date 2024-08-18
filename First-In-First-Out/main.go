package main

// src: https://girai.dev/blog/lru-cache-implementation-in-go/

import (
	"fmt"
)

func main() {
	fmt.Println("Implementing simple FIFO caching")
}

type FIFOCache struct {
	size  int               // capacity
	table map[string]string // hash table for list.List.Element existence check
	index []string
}

// Constructor initializes a new FIFOCache
func New(capacity int) FIFOCache {
	return FIFOCache{
		size:  capacity,
		table: make(map[string]string),
		index: make([]string, 0),
	}
}

// Get a list node from the hash map.
func (cache *FIFOCache) Get(key string) string {
	// check if list node exists
	if value, ok := cache.table[key]; ok {
		return value
	}
	return "not found"
}

// Put key and value in the LRUCache
func (cache *FIFOCache) Set(key, value string) {
	if len(cache.table) != len(cache.index) {
		fmt.Println("===== ERROR ====")
		fmt.Println("cache:", cache.table)
		fmt.Println("indexes:", cache.table)
		panic("Found size and index mismatch!")
	}

	// check if list node exists
	if ok := cache.table[key]; ok != "" {
		cache.MoveKeyToLast(key)
		cache.Update(key, value)
	} else {
		// here check what to do because key doesn't exist in cache

		// if full, delete the fist item
		if cache.IsFull() {
			// apply eviction
			cache.EvictFirst()
			cache.Add(key, value)
		} else {
			// not full
			cache.Add(key, value)
		}
	}
}

func (cache *FIFOCache) MoveKeyToLast(key string) {
	newOrdering := make([]string, 0)
	for _, value := range cache.index {
		if value != "" || value != key {
			newOrdering = append(newOrdering, value)
		}
	}
	newOrdering = append(newOrdering, key)
	cache.index = newOrdering
}

func (cache *FIFOCache) Update(key, value string) {
	cache.table[key] = value
}

func (cache *FIFOCache) EvictFirst() {
	firstItem := cache.index[0]
	rearranged := make([]string, 0)
	for i := 1; i < cache.size; i++ {
		rearranged = append(rearranged, cache.index[i])
	}
	delete(cache.table, firstItem)
	cache.index = rearranged
}

func (cache *FIFOCache) Add(key, value string) {
	cache.index = append(cache.index, key)
	cache.table[key] = value
}

func (cache *FIFOCache) IsFull() bool {
	return cache.size == len(cache.table)
}
