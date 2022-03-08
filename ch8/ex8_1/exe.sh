#!/bin/zsh


if [ -e "clockwall" ]; then
    rm clockwall
fi
go build clockwall.go
echo "clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030"
./clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030

