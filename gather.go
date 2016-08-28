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

func globDir(path string) ([]string, error) {
	dir := []string{}
	entries, err := filepath.Glob(filepath.Join(path, "[^.]*"))
	if err != nil {
		return nil, err
	}

	for _, e := range entries {
		fi, err := os.Stat(e)
		if err != nil {
			return nil, err
		}
		if fi.IsDir() {
			dir = append(dir, e)
		}
	}

	return dir, nil
}

func mkDir(destName string, ignoreCase bool) (string, error) {
	if ignoreCase {
		path, dirName := filepath.Split(destName)

		// pathにあるディレクトリ一覧を取得
		dirs, err := globDir(path)
		if err != nil {
			return destName, err
		}
		//小文字にして比較
		for _, v := range dirs {
			// 一致したら、小文字にする前の値を返す
			if strings.ToLower(filepath.Base(v)) == strings.ToLower(dirName) {

				return v, nil
			}
		}
	}

	// 一致するものがなかったらディレクトリ作成、すでにディレクトリかファイルがある場合エラーになるので、エラーは無視する
	_ = os.Mkdir(destName, 0755)
	return destName, nil
}

func move(destName, fileName string) error {

	absDestName, err := filepath.Abs(destName)
	if err != nil {
		log.Fatal(err)
	}

	//ディレクトリにファイル移動
	newName := filepath.Join(absDestName, filepath.Base(fileName))
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
	var winRename bool

	//コマンドラインオプション解析
	// デリミタ
	flag.StringVar(&delimiter, "d", " - ", "a delimiter which separates filenames into two parts")
	flag.BoolVar(&help, "h", false, "show help")
	flag.BoolVar(&ignoreCase, "i", false, "ignore case of dir names")
	flag.BoolVar(&dryRun, "dry-run", false, "dry run")
	flag.BoolVar(&fileonly, "f", false, "move only files")
	flag.BoolVar(&winRename, "win-rename", false, "replace characters forbidden on windows platforms with 2-byte characters")
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
		destName := strings.TrimSpace(extractDirname(f, delimiter))
		if destName == "" {
			continue
		}

		//ディレクトリ作成
		destName, err = mkDir(destName, ignoreCase)
		if err != nil {
			fmt.Printf("err %v", err)
		}
		//TODO win rename 実装

		// dry run 実装
		if dryRun {
			fmt.Printf("move %s to %s\n", f, filepath.Join(destName, filepath.Base(f)))
		} else {
			if err := move(destName, f); err != nil {
				log.Fatalf("error %v", err)
			}
		}
	}
}
