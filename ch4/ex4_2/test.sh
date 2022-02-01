#!/bin/zsh

go build ./main.go
echo "./main a"
./main a
echo "./main ab"
./main ab
echo "./main abc"
./main abc 
echo "./main abc -256"
./main abc -256
echo ./main abc -384
./main abc -384
echo "./main abc -512"
./main abc -512
echo "./main abc -1024"
./main abc -1024
echo "./main abc abc abc"
./main abc abc abc

