package main

import "testing"

func TestAppName(t *testing.T) {
	expect := "Zoo Application"
	if result := AppName(); expect != result {
		t.Errorf("\nexpect is  : %v\nbut result : %v\n", expect, result)
	}
}
