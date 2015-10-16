package main

import "testing"

func Test_check(t *testing.T) {
	players, count, apps := check(2003, 2003)

	switch {
	case len(players) != 1:
		t.Fatal("should be only one player in 2003")
	case players[0].Name() != "Wiki Gonzalez":
		t.Fatal("should be wiki!")
	case count != 1 || apps != 1:
		t.Fatal("should only be 1 player making 1 appearance")
	}
}

func Test_twice(t *testing.T) {
	Test_check(t)
}
