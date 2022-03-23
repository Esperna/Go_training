#!/bin/zsh
killall reverb3
go run reverb3.go &
go build ../ex8_3/netcat4.go
./netcat4
