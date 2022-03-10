#!/bin/zsh

CommandWithEcho(){
    echo $Command 
    $Command
}

Command="killall http5"
CommandWithEcho
go build http5.go
echo "./http5 &"
./http5 &

Command="fetch http://localhost:8000/list"
CommandWithEcho
Command="fetch http://localhost:8000/price"
CommandWithEcho
Command="fetch http://localhost:8000/price?item=socks"
CommandWithEcho
Command="fetch http://localhost:8000/price?item=shoes"
CommandWithEcho
Command="fetch http://localhost:8000/price?item=hat"
CommandWithEcho
Command="fetch http://localhost:8000/update?item=socks&price=6"
CommandWithEcho

