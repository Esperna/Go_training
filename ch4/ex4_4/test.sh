#!/bin/zsh

go build ./main.go
echo "./main 1"
./main 1
echo "./main 1 2 3"
./main 1 2 3
echo "./main 1 2 3 -r 1"
./main 1 2 3 -r 1
echo "./main 1 2 3 4 5 6 7 8 -r 5"
./main 1 2 3 4 5 6 7 8 -r 5
echo ./main 1 2 3 -r
./main 1 2 3 -r
echo "./main 1 2 3 4 5 6 7 8 -r 9"
./main 1 2 3 4 5 6 7 8 -r 9
echo "./main 1 2 3 4 5 6 7 8 -r 18"
./main 1 2 3 4 5 6 7 8 -r 18
