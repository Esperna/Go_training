#!/bin/zsh
CommandWithEcho(){
    echo $Command 
    $Command
}

Command="go run main.go -k1 Title -k2 Artist"
CommandWithEcho
Command="go run main.go -k1 Title -k2 Album"
CommandWithEcho
Command="go run main.go -k1 Title -k2 Year"
CommandWithEcho
Command="go run main.go -k1 Title -k2 Length"
CommandWithEcho
Command="go run main.go -k1 Length -k2 Year"
CommandWithEcho
Command="go run main.go -k1 Album -k2 Year"
CommandWithEcho
