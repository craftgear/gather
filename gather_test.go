package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

var testDir = "test/"

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
	err = move(filepath.Join("./test/", name, filepath.Base(fname)), fname)

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
	originalName := "./test/c - 03.txt"
	fname := filepath.Base(originalName)
	dest := extractDirname(fname, " - ")
	if dest != "c" {
		t.Errorf("dest should be 'c'")
	}
	dest = filepath.Join("./test", dest)
	dest, err := mkDir(dest, false)
	if err != nil {
		t.Error(err)
	}
	err = move(originalName, filepath.Join(dest, fname))
	if err != nil {
		t.Error(err)
	}

	if _, err = os.Stat("./test/c/c - 03.txt"); err != nil {
		t.Errorf("error %v", err)
	}

	move("./test/c/c - 03.txt", "./test/c - 03.txt")
	os.Remove("./test/c")
}

func TestWinCase(t *testing.T) {
	var cases = []struct {
		input    string
		expected string
	}{
		{"<", "＜"},
		{">", "＞"},
		{":", "："},
		{"\"", "”"},
		{"/", "／"},
		{"\\", "＼"},
		{"|", "｜"},
		{"?", "？"},
		{"*", "＊"},
	}

	for _, v := range cases {
		if renamed := winCaseRename(v.input); renamed != v.expected {
			t.Errorf("%s should be %s", renamed, v.expected)
		}
	}
}

func TestGetFilename(t *testing.T) {
	cases := []struct {
		f       string
		wincase bool
		file    string
	}{
		{
			"./test/hoge*.txt",
			true,
			"hoge＊.txt",
		},
	}

	for _, c := range cases {
		f := getFilename(c.f, c.wincase)
		if f != c.file {
			t.Errorf("%s should be equal to %s", f, c.file)
		}
	}
}

func TestGetDestDirName(t *testing.T) {
	cases := []struct {
		filename  string
		delimiter string
		dir       string
		out       string
	}{
		{
			"hoge - fuga.txt",
			" - ",
			"./test",
			"test/hoge",
		},
	}
	for _, c := range cases {
		result := getDestDirName(c.filename, c.delimiter, c.dir)
		if result != c.out {
			t.Errorf("%s should be equal to %s", result, c.out)
		}
	}
}

//TODO truncate オプション but how?
func TestGetTruncatedFilename(t *testing.T) {
	cases := []struct {
		delimiter string
		filename  string
		out       string
	}{
		{
			delimiter: " - ",
			filename:  "hoge - fuga.txt",
			out:       "fuga.txt",
		},
		{
			delimiter: " - ",
			filename:  "hoge_fuga.txt",
			out:       "hoge_fuga.txt",
		},
		{
			delimiter: "_",
			filename:  "hoge_fuga.txt",
			out:       "fuga.txt",
		},
	}

	for _, c := range cases {
		actual := getTruncatedFilename(c.filename, c.delimiter)
		if actual != c.out {
			t.Errorf("%s should be equal %s", actual, c.out)
		}
	}
}
