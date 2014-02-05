#!/bin/bash 

# compile coffeescript
# ./src/compile

# moving static files
cp -r ./css ../beta/static/
cp -r ./js ../beta/static/
cp -r ./lib ../beta/static/

# moving index file to templates
cp ./index.html ../beta/static/templates/
