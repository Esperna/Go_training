#!/bin/zsh

go build ./comma.go
echo "Input->Expected:"
echo "1              -> 1"
echo "1.2            -> 1.2"
echo "123            -> 123"
echo "1234           -> 1,234"
echo "12345          -> 12,345"
echo "12345.         -> 12,345."
echo "123456         -> 123,456"
echo "123456.789     -> 123,456.789"
echo "1234567        -> 1,234,567"
echo "1234567.8      -> 1,234,567.8"
echo "1234567.89     -> 1,234,567.89"
echo "+1234567.89    -> +1,234,567.89"
echo "-123.456789    -> -123.456789"
echo ".12            -> .12"
echo ".123           -> .123"
echo ".1234          -> .1234"
echo ""
echo "Actual:"
./comma 1
./comma 1.2
./comma 123
./comma 1234
./comma 12345
./comma 12345.
./comma 123456
./comma 123456.789
./comma 1234567
./comma 1234567.8
./comma 1234567.89
./comma +1234567.89
./comma -123.456789
./comma .12
./comma .123
./comma .1234
