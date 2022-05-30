#!/bin/zsh

go build ./main.go
./main < in
./main -h sha256 < in
./main -h sha384 < in
./main -h sha512 < in
./main -h sha1024 < in
