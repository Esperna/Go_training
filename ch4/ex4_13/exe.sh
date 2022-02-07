#!/bin/zsh

rm *.jpg
go build ./main.go
./main $1 $2
open *.jpg
