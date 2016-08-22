package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
)

func glob(dir string) ([]string, error) {
	files, err := filepath.Glob("[^.]*")
	if err != nil {
		return nil, err
	}
	return files, nil
}

func main() {
	//コマンドラインオプション解析
	//1. 対象ディレク折
	dir := flag.String("dir", "./", "a directory where files are in")
	//2. デリミタ
	delimiter := flag.String("delimiter", " - ", "a delimiter which separates filenames into two parts")
	fmt.Println(*dir, *delimiter)

	//ファイルリスト一覧取得
	files, err := glob(*dir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("files = %+v\n", files)

	//デリミタでファイル名を前後に分割、デリミタが見つからなければ何もしない
	//前半分のディレクトリを探す
	//ディレクトリがなければ作成
	//ディレクトリにファイル移動

}
