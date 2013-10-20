#!/bin/bash 

# installing python tools
sudo apt-get install build-essential python-dev

# installing django
sudo apt-get install python-django

# installing pip
sudo apt-get install python-pip

# installing pymongo
pip install pymongo

# installing the actual UI
mkdir $GOPATH/utils
cd $GOPATH/utils
git clone http://github.com/Fiedzia/Fang-of-Mongo.git

# now you need to change Fang-of-Mongo/fangofmongo/settings.py
# to have a time zone Europe/Warshaw or whatever, because the default one
# does not work (is not valid)
# then just run the server like python ./manage.py runserver
# and access it on: http://localhost:8000/fangofmongo
