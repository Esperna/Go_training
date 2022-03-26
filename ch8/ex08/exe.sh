#!/bin/zsh
killall reverb4
go build reverb4.go
./reverb4 &
sleep 1
nc localhost 8000
