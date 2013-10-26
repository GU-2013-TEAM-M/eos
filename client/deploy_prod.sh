#!/bin/bash 

# compile coffeescript
coffee -o lib/ -cj client.js src/service.coffee src/daemon.coffee src/controller.coffee src/communication.coffee src/graphs.coffee src/simulator.coffee

# minify js
# TODO

# moving static files
cp -r ./css ../prod/static/
cp -r ./js ../prod/static/
cp -r ./lib ../prod/static/

# moving index file to templates
cp ./index.html ../prod/static/templates/
