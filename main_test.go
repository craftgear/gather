package main

import (
	"reflect"
	"testing"
)

func TestGlob(t *testing.T) {
	expectedSlice := []string{"main.go", "main_test.go"}
	files, err := glob("./")
	if err != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(files, expectedSlice) {
		t.Fatal("files are different from expected ones.", files, expectedSlice)
	}
}
