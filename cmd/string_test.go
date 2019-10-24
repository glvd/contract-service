package main

import "testing"

func TestString(t *testing.T) {
	var ss []string
	ss = []string{"a", "b", "c", "d"}
	inputString(ss...)
}

func inputString(ss ...string) {
	_ = ss
}
