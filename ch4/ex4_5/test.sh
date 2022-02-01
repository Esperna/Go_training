#!/bin/zsh

go build ./main.go
echo "./main a"
./main a
echo "./main aaa"
./main aaa
echo "./main aabbccc"
./main aabbccc 
echo "./main engineer"
./main engineer
echo "./main a b c"
./main a b c
