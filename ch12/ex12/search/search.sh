#! /bin/zsh

call() {
    echo "fetch $1"
    fetch $1
}

go build
./search &
sleep 1
call 'http://localhost:12345/search'
call 'http://localhost:12345/search?l=golang&l=programming'
call 'http://localhost:12345/search?l=golang&l=programming&max=100'
call 'http://localhost:12345/search?x=true&l=golang&l=programming'
call 'http://localhost:12345/search?q=hello&x=123'
echo "ps"
ps

