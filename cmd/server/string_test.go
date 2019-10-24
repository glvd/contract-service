package main

import "testing"

// TestString ...
func TestString(t *testing.T) {
	var ss []string
	ss = []string{"a", "b", "c", "d"}
	inputString(ss...)
}

func inputString(ss ...string) {
	_ = ss
}
