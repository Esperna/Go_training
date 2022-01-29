#!/bin/zsh

go build ./main.go
./main 1
./main 1.2
./main 123
./main 1234
./main 12345
./main 12345.
./main 123456
./main 123456.789
./main 1234567
./main 1234567.8
./main 1234567.89
./main +1234567.89
./main -123.456789

