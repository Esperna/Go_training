#!/bin/zsh

go build ./main.go
echo "Input->Expected:"
echo 1              -> 1
echo 1.2            -> 1.2
echo 123            -> 123
echo 1234           -> 1,234
echo 12345          -> 12,345
echo 12345.         -> 12,345.
echo 123456         -> 123,456
echo 123456.789     -> 123,456.789
echo 1234567        -> 1,234,567
echo 1234567.8      -> 1,234,567.8
echo 1234567.89     -> 1,234,567.89
echo +1234567.89    -> +1,234,567.89
echo -123.456789    -> -123.456789

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

