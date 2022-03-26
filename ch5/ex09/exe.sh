#!/bin/zsh

go build ./main.go
echo "./main "\$hoge\$foo""
./main "\$hoge\$foo"
echo "./main "hoge\$foo""
./main "hoge\$foo"
echo "./main "\$hello\$world""
./main "\$hello\$world"
