#!/bin/zsh

go build ./issue.go
echo "Input->Expected:"
echo "./issue                   -> Invalid Number of Argument"
echo "./issue 1 2 3 4 5 6 8     -> Invalid Number of Argument"
echo "./issue -c GitHubID Token Title Body Labels Hoge -> Invalid Number of Argument"
echo "./issue -u IssueNo GitHubID Token Title Body Labels Hoge -> Invalid Number of Argument"
echo "./issue -d IssueNo GitHubID Token Title Body Labels Hoge -> Invalid Number of Argument"
echo "./issue -f IssueNo GitHubID Token Title Body Labels       -> Invalid Option"
echo "./issue -r                                                -> Successful Reading Issue info"
echo ""
echo "Actual:"
echo "./issue                   -> Invalid Number of Argument"
./issue     
echo ""
echo "./issue 1 2 3 4 5 6 8     -> Invalid Number of Argument"
./issue 1 2 3 4 5 6 8
echo ""
echo "./issue -c GitHubID Token Title Body Labels Hoge -> Invalid Number of Argument"
./issue -c GitHubID Token Title Body Labels Hoge
echo ""
echo "./issue -u IssueNo GitHubID Token Title Body Labels Hoge -> Invalid Number of Argument"
./issue -u IssueNo GitHubID Token Title Body Labels Hoge
echo ""
echo "./issue -d IssueNo GitHubID Token Title Body Labels Hoge -> Invalid Number of Argument"
./issue -d IssueNo GitHubID Token Title Body Labels Hoge
echo ""
echo "./issue -f IssueNo GitHubID Token Title Body Labels       -> Invalid Option"
./issue -f IssueNo GitHubID Token Title Body Labels
echo ""
echo "./issue -r                                                -> Successful Reading Issue info"
./issue -r 
