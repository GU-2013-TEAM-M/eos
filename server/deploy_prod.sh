#!/bin/bash 

# exporting flags for no reason
# but generally because it doesn't seem to get them
export GOPATH=$HOME/workplace
export PATH=$PATH:$GOPATH/bin

# compiling go
go install eos/server/server
echo "server compiled"

# running tests
cd ./server
go test
echo "tests have been run"

# copying executable
cp $GOPATH/bin/server ../../prod/
echo "the application has been deployed!"
