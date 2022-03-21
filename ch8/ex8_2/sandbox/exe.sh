#!/bin/zsh
if [ -e "ftp.go" ]; then
    rm ftp.go
fi
ftp -n < ftp_cmd.txt
