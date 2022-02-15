#!/bin/zsh

#1. read current issue list
#2. create new issue(title and body shall be set from vim)
#3. update issue(issueNo from shell arg and body shall be set from vim) 
#4. comment issue(issueNo(same as #3) from shell arg and body shall be set from vim)
#5. delete issue(issueNo shall be set)

go build ./issue.go
echo "./issue -r"
./issue -r
echo ""
echo "./issue -c vim"
./issue -c vim
echo ""
echo "./issue -u $1 vim"
./issue -u $1 vim 
echo ""
echo "./issue -uc $1 vim"
./issue -uc $1 vim 
echo ""
echo "./issue -d $2"
./issue -d $2
echo ""
