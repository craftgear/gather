[![Travis Build Status](https://travis-ci.org/craftgear/gather.svg?branch=master)](https://travis-ci.org/craftgear/gather)
[![Go Report Card](https://goreportcard.com/badge/github.com/craftgear/gather)](https://goreportcard.com/report/github.com/craftgear/gather)
[![LICENSE](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
<!--[![GoDoc](https://godoc.org/github.com/craftgear/gather?status.svg)](https://godoc.org/github.com/craftgear/gather)-->

English / [日本語](https://github.com/craftgear/gather/blob/master/README.ja.md)

# gather

*gather* is a simple tool to classify files in a directory into sub directories based on their names.

Say, you have files like named `project1 - 01.md` and `project2 - 01.md` and you execute *gather* on the directory it puts these two files into sub directories like `project1/project1 - 01.md` and `project2/project2 - 01.md`.

This time it uses a default delimiter ` - ` (space hyphen space). You can change the delimiter with `-d` option.

*gather* comes in handy when you have a slew of files like this example.

# Install

``go get github.com/craftgear/gather``

# Usage

To classify files in a current directory, run:
``gather``

To classify files in a specified directory, run:
``gather /home/username``

To change the delimiter, run with `-d` option:
``gather -d _``
In this case sub directories named by the first part of files before the first `_`.

To ignore cases on making sub directories, run with `-i` option:
``gather -i``
When `-i` option is specifed, gather puts `Project - 01.md` and `project - 01.md` under the same directory.

To display files to move without moving (so-called "dry run"), run:
``gather -dry-run``

To move only files, run with `-f` option:
``gather -f``
In default *gather* moves files and directories under sub directories. With `-f` option, it treats only files to move.

To replace characters on Windows platforms, run with `-wincase` option:
``gather -wincase``
Details on [wincase](https://github.com/craftgear/wincase).

(WIP)To rename files to names without directory names after moving, run with `-truncate` option:
``gather -truncate``
If `-truncate` option is specified, *gather* moves files and renames them at the same time. New filenames will be without direnctory names they are in. ex) `Project - 01.md` would become `Project/01.md`.

To run in verbose mode, run:
``gather -v ./``

To show help, run:
``gather -h``

### Author
craftgear (https://twitter.com/craftgear)

