#!/bin/zsh
killall reverb1
reverb1 &
go build netcat4.go
./netcat4
