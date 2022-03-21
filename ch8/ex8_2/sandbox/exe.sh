#!/bin/zsh
if [ -e "ftp.go" ]; then
    rm ftp.go
fi
if [ -e "./data/sample1.txt" ]; then
    rm ./data/sample1.txt
fi
ftp -n < ftp_cmd.txt
