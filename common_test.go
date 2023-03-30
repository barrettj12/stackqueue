package main

import "testing"

func checkLen(t *testing.T, collection interface{ Len() int }, expected int) {
	len := collection.Len()
	if len != expected {
		t.Errorf("collection.Len is %d, expected %d", len, expected)
	}
}
