package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func glob(dir string) ([]string, error) {
	files, err := filepath.Glob("[^.]*")
	if err != nil {
		return nil, err
	}
	return files, nil
}

func extractDirname(filename, delimiter string) string {
	a := strings.Split(filename, delimiter)
	if a[0] == filename {
		return ""
	}
	return a[0]
}

func move(fname, dname string) error {
	//ディレクトリがなければ作成
	if err := os.Mkdir(dname, os.ModeDir); err != nil {
		return err
	}

	newName := filepath.Join(dname, string(filepath.Separator), fname)
	fmt.Printf("move %v to %v\n", fname, newName)
	//ディレクトリにファイル移動
	if err := os.Rename(fname, newName); err != nil {
		return err
	}

	return nil
}

func main() {
	var dir string
	var delimiter string

	//コマンドラインオプション解析
	//1. 対象ディレクトリ
	flag.StringVar(&dir, "dir", "./", "a directory where files are in")
	//2. デリミタ
	flag.StringVar(&delimiter, "delimiter", " - ", "a delimiter which separates filenames into two parts")
	flag.Parse()

	fmt.Println(dir, delimiter)

	//ファイルリスト一覧取得
	files, err := glob(dir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("files = %+v\n", files)

	for _, v := range files {
		fmt.Printf("v = %+v\n", v)
		//デリミタでファイル名を前後に分割、デリミタが見つからなければ何もしない
		dname := extractDirname(v, delimiter)
		if dname == "" {
			continue
		}

		if err := move(dname, v); err != nil {
			log.Fatalf("error %v", err)
		}

	}

}
