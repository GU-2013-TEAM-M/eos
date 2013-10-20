#!/bin/bash 

# installing Go language
sudo apt-get install golang

# installing MongoDB
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv 7F0CEB10
echo 'deb http://downloads-distro.mongodb.org/repo/debian-sysvinit dist 10gen' | sudo tee /etc/apt/sources.list.d/mongodb.list
sudo apt-get update
sudo apt-get install mongodb-10gen

# adding MongoDB driver for Go
sudo apt-get install bzr
go get labix.org/v2/mgo

# setting up the workplace (do it yourself)
# mkdir $HOME/go
# export GOPATH=$HOME/workplace
# export PATH=$PATH:$GOPATH/bin
