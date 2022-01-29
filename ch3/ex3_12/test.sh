#!/bin/zsh

go build ./main.go
echo "Input->Expected:"
echo "a             -> Invalid Number of Argument"
echo "ab ab         -> Anagram!"
echo "ab ba         -> Anagram!"
echo "ab bac        -> Not Anagram!"
echo "abc def       -> Not Anagram!"
echo "cba bac       -> Anagram!"
echo "aabb baba     -> Anagram!"
echo "abcde adecb   -> Anagram!"
echo "bbbab aabab   -> Not Anagram!"
echo "abc def hij   -> Invalid Number of Argument"
echo ""
echo "Actual:"
./main a
./main ab ab
./main ab ba
./main ab bac
./main abc def
./main cba bac
./main aabb baba
./main abcde adecb
./main bbbab aabab
./main abc def hij

