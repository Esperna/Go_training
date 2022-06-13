#! /bin/zsh

go build transitivePackages.go
echo "./transitivePackages $1 $2"
./transitivePackages $1 $2
