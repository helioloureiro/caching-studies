package main

import (
	"fmt"
	"testing"
)

func TestCreateCacheSize2(t *testing.T) {
	t.Log("Creating cache with size=2")
	cache := New(2) // nil
	t.Logf("cache data: %v", cache)
	t.Log("Adding (1,10) to the cache")
	cache.Set("1", "10") // nil, linked list: [1:10]
	t.Logf("cache data: %v", cache)
	t.Log("Adding (2,20) to the cache")
	cache.Set("2", "20") // nil, linked list: [2:20, 1:10]
	t.Logf("cache data: %v", cache)

	resp := cache.Get("1")
	if resp != "10" {
		t.Errorf("Cache sent: %v - expected 10", resp)
	}
	t.Logf("cache data: %v", cache)

	resp = cache.Get("2")
	if resp != "20" {
		t.Errorf("Cache sent: %v - expected 20", resp)
	}
	t.Logf("cache data: %v", cache)

	for i := 1; i < 4; i++ {
		t.Logf("Cache(%v):%v", i, cache.Get(fmt.Sprintf("%d", i)))
	}

	t.Log("Adding (3,30) to the cache  - it will replace (1, 10)")
	cache.Set("3", "30") // nil, linked list: [3:30, 1:10]
	t.Logf("cache data: %v", cache)

	for i := 1; i < 4; i++ {
		t.Logf("Cache(%v):%v", i, cache.Get(fmt.Sprintf("%d", i)))
	}
	t.Logf("cache data: %v", cache)

	resp = cache.Get("1")
	if resp != "not found" {
		t.Errorf("Cache sent: %v - expected \"not found\"", resp)
	}

	resp = cache.Get("2")
	if resp != "20" {
		t.Errorf("Cache sent: %v - expected 20", resp)
	}

	resp = cache.Get("3")
	if resp != "30" {
		t.Errorf("Cache sent: %v - expected 30", resp)
	}

	t.Logf("cache data: %v", cache)

	t.Log("Adding (4,40) to the cache  - it will replace (2, 20)")
	cache.Set("4", "40")
	t.Logf("cache data: %v", cache)

	resp = cache.Get("1")
	if resp != "not found" {
		t.Errorf("Cache sent: %v - expected \"not found\"", resp)
	}
	resp = cache.Get("2")
	if resp != "not found" {
		t.Errorf("Cache sent: %v - expected \"not found\"", resp)
	}

	resp = cache.Get("3")
	if resp != "30" {
		t.Errorf("Cache sent: %v - expected 30", resp)
	}
	resp = cache.Get("4")
	if resp != "40" {
		t.Errorf("Cache sent: %v - expected 40", resp)
	}
}

func TestCreateCacheSize10(t *testing.T) {
	// t.SkipNow()
	t.Log("Creating cache with size=10")
	cache := New(10)
	t.Log("Adding (1,10) to the cache")
	cache.Set("1", "10")
	t.Log("Adding (2,20) to the cache")
	cache.Set("2", "20")
	t.Log("Adding (3,30) to the cache")
	cache.Set("3", "30")
	t.Log("Adding (4,40) to the cache")
	cache.Set("4", "40")
	t.Log("Adding (5,50) to the cache")
	cache.Set("5", "50")
	t.Log("Adding (6,60) to the cache")
	cache.Set("6", "60")
	t.Log("Adding (7,70) to the cache")
	cache.Set("7", "70")
	t.Log("Adding (8,80) to the cache")
	cache.Set("8", "80")
	t.Log("Adding (9,90) to the cache")
	cache.Set("9", "90")
	t.Log("Adding (10,100) to the cache")
	cache.Set("10", "100")

	resp := cache.Get("1")
	if resp != "10" {
		t.Errorf("Cache sent: %v - expected 10", resp)
	}

	resp = cache.Get("5")
	if resp != "50" {
		t.Errorf("Cache sent: %v - expected 50", resp)
	}

	t.Log("Adding (11,110) to the cache - replacing (1, 10)")
	cache.Set("11", "110")
	resp = cache.Get("10")
	if resp != "100" {
		t.Errorf("Cache sent: %v - expected 100", resp)
	}
	resp = cache.Get("11")
	if resp != "110" {
		t.Errorf("Cache sent: %v - expected 110", resp)
	}
	if cache.index[0] != "2" {
		t.Errorf("Expected index \"11\" but received: %v", cache.index[0])
	}

	for i := 1; i < 12; i++ {
		t.Logf("Cache(%v):%v", i, cache.Get(fmt.Sprintf("%d", i)))
	}
	// resp = cache.Get(1)
	// if resp != -1 {
	// 	t.Errorf("Cache sent: %v - expected -1", resp)
	// }

	t.Log("Adding (12,120) to the cache - replacing (2, 20)")
	cache.Set("12", "120")
	for i := 1; i < 12; i++ {
		t.Logf("Cache(%v):%v", i, cache.Get(fmt.Sprintf("%d", i)))
	}
	if cache.index[0] != "3" {
		t.Errorf("Expected index \"11\" but received: %v", cache.index[0])
	}

}
