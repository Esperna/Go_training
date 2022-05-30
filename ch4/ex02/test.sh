#!/bin/zsh

go build ./main.go
echo "./main < in"
./main < in
echo "./main -h sha256 < in"
./main -h sha256 < in
echo "./main -h sha384< in"
./main -h sha384 < in
echo "./main -h sha512 < in"
./main -h sha512 < in
echo "./main -h sha1024 < in"
./main -h sha1024 < in
