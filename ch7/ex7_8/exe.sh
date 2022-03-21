#!/bin/zsh
CommandWithEcho(){
    echo $Command 
    $Command
}

Command="go run main.go -k1 Title -k2 Artist"
CommandWithEcho
