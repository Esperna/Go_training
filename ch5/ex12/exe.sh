#!/bin/zsh

echo "go run ./outline.go http://gopl.io > out.txt"
go run ./outline.go http://gopl.io > out.txt
echo "diff out.txt answer.txt"
diff out.txt answer.txt 
