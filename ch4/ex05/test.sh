#!/bin/zsh

go build ./main.go
echo "./main a"
./main a
echo "./main abc abc"
./main abc abc
echo "./main abc abc abc"
./main abc abc abc
echo "./main abc cde abc cde"
./main abc cde abc cde
echo "./main abc cde cde abc"
./main abc cde cde abc
echo "./main aa bb aa aaa aa aa aa ab ab"
./main aa bb aa aaa aa aa aa ab ab
echo "./main a aa a aaa aaa a a aa aa a aaa"
./main a aa a aaa aaa a a aa aa a aaa
