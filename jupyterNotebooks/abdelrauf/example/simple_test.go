package main

import "testing"

func simpleTest(t *testing.T) {
	res := 2 + 2
	if res != 4 {
		t.Errorf("%d olmalidir", 4)
	}
}
