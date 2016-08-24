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
	files, err := filepath.Glob(dir + string(filepath.Separator) + "[^.]*")
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

//func mkDir(dname string, ignoreCase bool) error {

//}

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
	var help bool
	//var ignoreCase bool

	//コマンドラインオプション解析
	// デリミタ
	flag.StringVar(&delimiter, "d", " - ", "a delimiter which separates filenames into two parts")
	flag.BoolVar(&help, "h", false, "show help")
	//flag.StringBool(&ignoreCase, "ignore case", false, "ignore case of dir names")
	flag.Parse()

	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}
	// 対象ディレクトリ
	if args := flag.Args(); len(args) > 0 {
		dir = args[0]
	} else {
		dir = "./"
	}
	fmt.Printf("dir = %+v\n", dir)
	fmt.Printf("delimiter = %+v len='%+v'\n", delimiter, len(delimiter))

	if len(delimiter) == 0 {
		fmt.Println("delimiter cannot be empty.")
		os.Exit(0)
	}

	//ファイルリスト一覧取得
	files, err := glob(dir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("files = %+v\n", files)

	for _, f := range files {
		//fmt.Printf("v = %+v\n", f)
		//デリミタでファイル名を前後に分割、デリミタが見つからなければ何もしない
		dname := extractDirname(f, delimiter)
		//fmt.Printf("dname = %+v\n", dname)
		if dname == "" {
			continue
		}

		if err := move(dname, f); err != nil {
			log.Fatalf("error %v", err)
		}

	}

}
