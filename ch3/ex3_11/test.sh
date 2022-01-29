#!/bin/zsh

go build ./main.go
echo "Expected:"
echo 1
echo 1.2
echo 123
echo 1234
echo 12345
echo 12345.
echo 123456
echo 123456.789
echo 1234567
echo 1234567.8
echo 1234567.89
echo +1234567.89
echo -123.456789

echo "Actual:"
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

