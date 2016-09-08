#! /bin/bash
go test -v ./... \
| sed ''/PASS/s//$(printf "\033[1;32mPASS\033[0m")/'' \
| sed ''/FAIL/s//$(printf "\033[1;31mFAIL\033[0m")/'' \
| sed ''/TODO/s//$(printf "\033[1;93mTODO\033[0m")/'' \
| sed ''/ok/s//$(printf "\033[32mok\033[0m")/'' \



