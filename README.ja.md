[![Travis Build Status](https://travis-ci.org/craftgear/gather.svg?branch=master)](https://travis-ci.org/craftgear/gather)
[![Go Report Card](https://goreportcard.com/badge/github.com/craftgear/gather)](https://goreportcard.com/report/github.com/craftgear/gather)
[![LICENSE](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
<!--[![GoDoc](https://godoc.org/github.com/craftgear/gather?status.svg)](https://godoc.org/github.com/craftgear/gather)-->

[English](https://github.com/craftgear/gather/blob/master/README.md) / 日本語

# gatherでできること

指定したディレクトリにあるファイルを、区切り文字で区切って、前方一致するものをディレクトリにまとめます。

`project1 - 01.md` と `project2 - 01.md` という二つのファイルがあるディレクトリで、 *gather* を実行すると、この二つのファイルはそれぞれ `project1/project1 - 01.md` と `project2/project2 - 01.md` というディレクトリ以下に移動されます。

デフォルトの区切り文字は ` - ` (半角スペース、ハイフン、半角スペース) です。デリミタは `-d` オプションで変更できます。

*gather* はこの例のようなファイルが大量にあるときに便利なツールです.

# インストール

``go get github.com/craftgear/gather``

# 使い方

カレントディレクトリのファイルを、デフォルトの区切り文字で分類するには:
``gather``

ディレクトリを指定して、ファイルの分類をするには:
``gather /home/username``

区切り文字を変更するには:
``gather -d="_"``
この場合、区切り文字は `_` アンダースコアになり、ファイルの名前をアンダースコアで区切った最初の部分がサブディレクトリの名前になります。

新しく作るディレクトリの名前のおおもじ小文字を区別しないようにするには:
``gather -i``
`-i`オプションを付けて実行すると、`Project - 01.md` と `project - 01.md` という二つのファイルは同じディレクトリ以下に置かれます。

実際にファイル移動を行わず、どのような作業が行われるかを見るには（ドライラン）:
``gather -dry-run``

ディレクトリを移動対象から外し、ファイルのみを移動するには:
``gather -f``
デフォルトではファイルとディレクトリの両方を移動対象にします。

ウィンドウズ環境でファイルに使えない文字を全角文字に変換するには:
``gather -wincase``
詳しくは [wincase](https://github.com/craftgear/wincase) をご覧ください。

(実装中) ファイル名から、サブディレクトリの部分を削除して移動するには:
``gather -truncate``

`-truncate`オプションを指定すると、ファイルを移動すると同時に、フィアル名が短くなるようにリネームを行います。`Project - 01.md` というファイルは、 `Project/01.md`　となります。

作業状況を出力しながら実行するには:
``gather -v ./``

ヘルプを表示するには:
``gather -h``

### 作者
craftgear (https://twitter.com/craftgear)


