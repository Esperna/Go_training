#!/bin/zsh

go build ./tempflag.go
./tempflag
./tempflag -temp -18C
./tempflag -temp 212F
./tempflag -temp 273.15K
./tempflag -help
