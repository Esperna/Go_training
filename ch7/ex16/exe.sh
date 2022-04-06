#!/bin/zsh

CommandWithEcho(){
    echo $Command 
    $Command
}

Command="killall main"
CommandWithEcho
go build main.go
echo "./main &"
./main &
sleep 1

Command="fetch http://localhost:8000/calc"
CommandWithEcho
Command="fetch http://localhost:8000/calc?expr=1%2b2"
CommandWithEcho
Command="fetch http://localhost:8000/calc?expr=1-2"
CommandWithEcho
Command="fetch http://localhost:8000/calc?expr=x*y&x=12&y=13"
CommandWithEcho


