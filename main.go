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

//func mkDir(dirName string, ignoreCase bool) (string, error) {
//path, dir := filepath.Split(dname)

////pathのディレクトリ一覧
////小文字にして比較

//}

func move(dirName, fileName string) error {
	//ディレクトリ作成、すでにディレクトリかファイルがある場合エラーになるので、エラーは無視する
	// TODO ignore caseをどうやったら実現できるか？
	_ = os.Mkdir(dirName, 0755)

	newName := filepath.Join(dirName, string(filepath.Separator), filepath.Base(fileName))
	//ディレクトリにファイル移動
	if err := os.Rename(fileName, newName); err != nil {
		return err
	}

	return nil
}

func main() {
	var dir string
	var delimiter string
	var help bool
	var ignoreCase bool
	var dryRun bool
	var fileonly bool

	//コマンドラインオプション解析
	// デリミタ
	flag.StringVar(&delimiter, "d", " - ", "a delimiter which separates filenames into two parts")
	flag.BoolVar(&help, "h", false, "show help")
	flag.BoolVar(&ignoreCase, "i", false, "ignore case of dir names")
	flag.BoolVar(&dryRun, "dry-run", false, "dry run")
	flag.BoolVar(&fileonly, "f", false, "move only files")
	flag.Parse()

	if help {
		fmt.Println("\ngather - a simple utility to move files\n")
		fmt.Println("Usage: gather [options] target_dir")
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	}
	// 対象ディレクトリ
	if args := flag.Args(); len(args) > 0 {
		dir = args[0]
	} else {
		dir = "./"
	}

	if len(delimiter) == 0 {
		fmt.Println("delimiter cannot be empty.")
		os.Exit(0)
	}

	//ファイルリスト一覧取得
	files, err := glob(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		//デリミタでファイル名を前後に分割、デリミタが見つからなければ何もしない
		dirName := extractDirname(f, delimiter)
		if dirName == "" {
			continue
		}

		//TODO dirNameがスペースで終わらないようにする

		if err := move(dirName, f); err != nil {
			log.Fatalf("error %v", err)
		}
	}
}
