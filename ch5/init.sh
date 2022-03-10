#!/bin/zsh

if test ! -e $1 ; then
    mkdir $1
    go build ./dispCurrentFolder.go
    var=$(pwd | xargs -I@ ./dispCurrentFolder @)
    cd $1
    cp ~/template/main.go ./
    go mod init $var/$1
    git add main.go go.mod
    git commit -m"initial commit for practice"
else
    echo "$1 already exists"
fi
