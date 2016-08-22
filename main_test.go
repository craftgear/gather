package main

import (
	"os"
	"reflect"
	"testing"
)

func TestGlob(t *testing.T) {
	expectedSlice := []string{"a - 01.txt", "b - 01.txt"}
	files, err := glob("./test")
	if err != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(files, expectedSlice) {
		t.Fatal("files are different from expected ones.", files, expectedSlice)
	}
}

func TestExtractDirname(t *testing.T) {
	var cases = []struct {
		n        string
		expected string
	}{
		{"a.txt", ""},
		{"a - 01.txt", "a"},
	}

	for _, v := range cases {
		if result := extractDirname(v.n, " - "); result != v.expected {
			t.Errorf("expected %+v, but got %+v", v.expected, result)
		}
	}
}

func TestMove(t *testing.T) {
	fname := "./test/a - 01.txt"
	name := extractDirname(fname, " - ")
	move(name, fname)
	if _, err := os.Stat("./test/a/a - 01.txt"); err != nil {
		t.Errorf("error %v", err)
	}
}
