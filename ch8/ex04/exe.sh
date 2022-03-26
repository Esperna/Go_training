#!/bin/zsh
killall reverb3
#go run reverb3.go &
go build reverb3.go
./reverb3 &
go build ../ex03/netcat4.go
./netcat4
