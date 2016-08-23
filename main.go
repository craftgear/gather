package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// TODO エラーメッセージの多言語化

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

func move(dname, fname string) error {
	//ディレクトリ作成、すでにディレクトリかファイルがある場合エラーになるので、エラーは無視する
	// TODO ignore caseをどうやったら実現できるか？
	_ = os.Mkdir(dname, 0755)

	newName := filepath.Join(dname, string(filepath.Separator), filepath.Base(fname))
	//ディレクトリにファイル移動
	if err := os.Rename(fname, newName); err != nil {
		return err
	}

	return nil
}

func main() {
	var dir string
	var delimiter string
	//var ignoreCase bool

	//コマンドラインオプション解析
	//1. 対象ディレクトリ
	flag.StringVar(&dir, "dir", "./", "a directory where files are in")
	//2. デリミタ
	flag.StringVar(&delimiter, "delimiter", " - ", "a delimiter which separates filenames into two parts")
	//flag.StringBool(&ignoreCase, "ignore case", false, "ignore case of dir names")
	flag.Parse()

	fmt.Println(dir, delimiter)

	//ファイルリスト一覧取得
	files, err := glob(dir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("files = %+v\n", files)

	for _, f := range files {
		fmt.Printf("v = %+v\n", f)
		//デリミタでファイル名を前後に分割、デリミタが見つからなければ何もしない
		dname := extractDirname(f, delimiter)
		if dname == "" {
			continue
		}

		if err := move(dname, f); err != nil {
			log.Fatalf("error %v", err)
		}

	}

}
