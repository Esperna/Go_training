#!/bin/zsh

mkdir $1
go build ./dispCurrentFolder
var=$(pwd | xargs -I@ ./dispCurrentFolder @)
cd $1
cp ~/template/main.go ./
go mod init $var/$1
git add main.go go.mod
git commit -m"initial commit for practice"
