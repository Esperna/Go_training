#!/bin/zsh

CommandWithEcho(){
    echo $Command 
    $Command
}

Command="killall http6"
CommandWithEcho
go build http6.go
echo "./http6 &"
./http6 &
sleep 1

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
Command="fetch http://localhost:8000/update?item=socks"
CommandWithEcho
Command="fetch http://localhost:8000/update?item=socks&price=-10"
CommandWithEcho
Command="fetch http://localhost:8000/update?item=socks&price=10yen"
CommandWithEcho
Command="fetch http://localhost:8000/delete"
CommandWithEcho
Command="fetch http://localhost:8000/delete?item=socks"
CommandWithEcho
Command="fetch http://localhost:8000/create"
CommandWithEcho
Command="fetch http://localhost:8000/create?item=socks"
CommandWithEcho
Command="fetch http://localhost:8000/create?item=shoes"
CommandWithEcho
Command="fetch http://localhost:8000/create?item=shoes&price=-10"
CommandWithEcho
Command="fetch http://localhost:8000/create?item=hat&price=10yen"
CommandWithEcho
Command="fetch http://localhost:8000/create?item=hat&price=1000"
CommandWithEcho

