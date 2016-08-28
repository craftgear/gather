package main

import (
	"os"
	"reflect"
	"testing"
)

var testDir string = "_test/"

func TestMain(m *testing.M) {
	m.Run()
}

func TestGlob(t *testing.T) {
	expected := []string{
		testDir + "B",
		testDir + "a",
		testDir + "a - 01.txt",
		testDir + "b - 02.txt",
		testDir + "c - 03.txt",
	}
	files, err := glob(testDir)

	if err != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(files, expected) == false {
		t.Fatal("files are different from expected ones.", files, expected)
	}
}

func TestExtractDirname(t *testing.T) {
	var cases = []struct {
		n        string
		expected string
	}{
		{"a", ""},
		{"a.txt", ""},
		{"a - 01.txt", "a"},
	}

	for _, v := range cases {
		if result := extractDirname(v.n, " - "); result != v.expected {
			t.Errorf("expected %+v, but got %+v", v.expected, result)
		}
	}
}

func TestMoveFail(t *testing.T) {
	fname := "./test/a - 01.txt"
	name := extractDirname(fname, " - ")
	name, err := mkDir(name, false)
	err = move(name, fname)

	if err == nil {
		t.Error("An error should be occured.")
	}
}

func TestGlobDir(t *testing.T) {
	expected := []string{testDir + "B"}
	dirs, err := globDir(testDir)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(dirs, expected) {
		t.Fatal("dirs are different from expected ones.", dirs, expected)
	}

}

func TestMkDir(t *testing.T) {
	dirName, err := mkDir(testDir+"B", false)
	if err != nil {
		t.Fatal("err should be nil.", err)
	}
	if dirName != testDir+"B" {
		t.Errorf("testDir should be %s but got %s", testDir+"B", dirName)
	}

	dirName, err = mkDir(testDir+"b", true)
	if err != nil {
		t.Fatal("err should be nil.", err)
	}
	if dirName != testDir+"B" {
		t.Errorf("testDir should be %s but got %s", testDir+"B", dirName)
	}

	dirName, err = mkDir(testDir+"a", false)
	if err != nil {
		t.Fatal(err)
	}
	if dirName != testDir+"a" {
		t.Errorf("testDir should be %s but got %s", testDir+"a", dirName)
	}

}

func TestMove(t *testing.T) {
	fname := "./test/c - 03.txt"
	dest := extractDirname(fname, " - ")
	dest, err := mkDir(dest, false)
	err = move(dest, fname)
	if err != nil {
		t.Error(err)
	}

	if _, err = os.Stat("./test/c/c - 03.txt"); err != nil {
		t.Errorf("error %v", err)
	}

	move("./test", "./test/c/c - 03.txt")
	os.Remove("./test/c")

	//TODO dryRun
}
