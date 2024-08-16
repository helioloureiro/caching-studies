package main

import "testing"

func TestCreateCache(t *testing.T) {
	cache := New[string, string](10, 0)
	cache.Set("hello", "world")
}
