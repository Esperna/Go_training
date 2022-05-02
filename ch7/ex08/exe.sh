#!/bin/zsh
CommandWithEcho(){
    echo $Command 
    $Command
}

go build customSort.go trackData.go 
Command="./customSort -k1 Title -k2 Artist"
CommandWithEcho
Command="./customSort -k1 Title -k2 Album"
CommandWithEcho
Command="./customSort -k1 Title -k2 Year"
CommandWithEcho
Command="./customSort -k1 Title -k2 Length"
CommandWithEcho
Command="./customSort -k1 Length -k2 Year"
CommandWithEcho
Command="./customSort -k1 Album -k2 Year"
CommandWithEcho
