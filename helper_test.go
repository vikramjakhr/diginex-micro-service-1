package main

import "testing"

// Reverse function test
func TestReverse(t *testing.T) {
	data := Data{Message: "abcdefg"}
	result := "gfedcba"

	if data.reverse(); data.Message != result {
		t.Errorf("Expected %s, Found %s", result, data.Message)
	}
}
