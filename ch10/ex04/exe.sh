#! /bin/zsh

go build transitivePackages.go
echo "./transitivePackages $1"
./transitivePackages $1
