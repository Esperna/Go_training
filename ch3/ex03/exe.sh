#!/bin/zsh

go build ./main.go
./main > out.html
open out.html
