#!/bin/zsh

go build ./anagram.go
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
./anagram a
./anagram ab ab
./anagram ab ba
./anagram ab bac
./anagram abc def
./anagram cba bac
./anagram aabb baba
./anagram abcde adecb
./anagram bbbab aabab
./anagram abc def hij

