package main

import "testing"

func TestCreateCache(t *testing.T) {
	cache := New()
	cache.Set("hello", "world")
}
