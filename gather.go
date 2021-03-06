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
	files, err := filepath.Glob(filepath.Join(dir, "[^.]*"))
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

func move(originalFileName, destName string) error {
	absDestName, err := filepath.Abs(destName)
	if err != nil {
		log.Fatal(err)
	}

	//ディレクトリにファイル移動
	if err := os.Rename(originalFileName, absDestName); err != nil {
		return err
	}

	return nil
}

func winCaseRename(filename string) string {
	var cases = []struct {
		input  string
		output string
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
		if strings.Index(filename, v.input) > -1 {
			filename = strings.Replace(filename, v.input, v.output, -1)
		}
	}
	return filename
}

func getFilename(f string, winCase bool) string {
	filename := filepath.Base(f)
	if winCase {
		filename = winCaseRename(filename)
	}
	return filename
}

func getDestDirName(filename, delimiter, dir string) string {
	//デリミタでファイル名を前後に分割、デリミタが見つからなければ何もしない
	newDirName := strings.TrimSpace(extractDirname(filename, delimiter))
	if newDirName == "" {
		return ""
	}

	destDirName := filepath.Join(dir, newDirName)

	return destDirName
}

func getTruncatedFilename(filename, delimiter string) string {
	temp := strings.SplitN(filename, delimiter, 2)
	if len(temp) != 2 {
		return filename
	}
	return temp[1]
}

func main() {
	var dir string
	var delimiter string
	var help bool
	var ignoreCase bool
	var dryRun bool
	var fileonly bool
	var winCase bool
	var truncate bool

	// TODO truncateオプション

	//コマンドラインオプション解析
	// デリミタ
	flag.StringVar(&delimiter, "d", " - ", "a delimiter which separates filenames into two parts")
	flag.BoolVar(&help, "h", false, "show help")
	flag.BoolVar(&ignoreCase, "i", false, "ignore case of dir names")
	flag.BoolVar(&dryRun, "dry-run", false, "dry run")
	flag.BoolVar(&fileonly, "f", false, "move files only")
	flag.BoolVar(&winCase, "wincase", false, "replace characters forbidden on windows platforms with 2-byte characters")
	flag.BoolVar(&truncate, "truncate", false, "truncate filenames to exclude the first matching part and delimiter")
	flag.Parse()

	if help {
		fmt.Print("\ngather - a simple utility to move files\n\n")
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

	for _, file := range files {
		filenameWithPath := file
		filename := getFilename(filenameWithPath, winCase)

		destDirName := getDestDirName(filename, delimiter, dir)
		if destDirName == "" {
			continue
		}

		if truncate {
			filename = getTruncatedFilename(filename, delimiter)
		}

		//ディレクトリ作成
		destDirName, err = mkDir(destDirName, ignoreCase)
		if err != nil {
			log.Fatal("mkDir failed\n", err)
		}

		if dryRun {
			fmt.Printf("move %s to %s\n", filenameWithPath, filepath.Join(destDirName, filename))
		} else {
			if err := move(filenameWithPath, filepath.Join(destDirName, filename)); err != nil {
				log.Fatal("move failed\n", err)
			}
		}
	}
}
