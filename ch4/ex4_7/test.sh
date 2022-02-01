#!/bin/zsh

go build ./main.go
echo "./main a"
./main a
echo "./main abc"
./main abc
echo "./main abcde"
./main abcde
echo "./main abc123"
./main abc123
echo ./main a b c
./main a b c
