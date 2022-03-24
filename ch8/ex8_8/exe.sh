#!/bin/zsh
killall reverb4
go build reverb4.go
./reverb4 &
go build ../ex8_3/netcat4.go
./netcat4

