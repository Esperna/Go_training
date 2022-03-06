#!/bin/zsh

killall clock
rm ./clock
go build ./clock.go
./clock -port 8010
./clock -port 8020
./clock -port 8030
