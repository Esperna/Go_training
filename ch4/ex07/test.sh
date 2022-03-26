#!/bin/zsh

go build ./reverseByteSlice.go
echo "./reverseByteSlice 太陽"
./reverseByteSlice 太陽
echo "./reverseByteSlice 太平洋"
./reverseByteSlice 太平洋
echo "./reverseByteSlice abcdef"
./reverseByteSlice abcdef
echo "/reverseByteSlice αβγ"
./reverseByteSlice αβγ
echo "./reverseByteSlice こんにちはHello" 
./reverseByteSlice こんにちはHello
