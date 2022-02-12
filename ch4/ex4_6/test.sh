#!/bin/zsh

go build ./main.go
echo "./main あ　　い　　う"
./main "あ　　い　　う"
echo "./main a   b       c"
./main "a   b       c"
echo "./main a        b 　　c"　
./main "a        b 　　c"　
