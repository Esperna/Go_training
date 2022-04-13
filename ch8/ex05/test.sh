#!/bin/zsh

call () {
    echo "GOMACSPROCS=$1 go test -bench=."
    GOMACSPROCS=$1 go test -bench=.
}

call 1
call 2
