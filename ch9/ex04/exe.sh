#! /bin/zsh

call() {
    echo "./main -n $1"
    ./main -n $1
}

call 1
call 10
call 100
call 1000
call 10000
call 100000
call 1000000
call 10000000
call 100000000


