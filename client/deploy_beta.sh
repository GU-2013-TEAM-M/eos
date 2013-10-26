#!/bin/bash 

# compile coffeescript
coffee -bo lib/ -cj client.js src/service.coffee src/daemon.coffee src/controller.coffee src/communication.coffee src/graphs.coffee src/simulator.coffee

# moving static files
cp -r ./css ../beta/static/
cp -r ./js ../beta/static/
cp -r ./lib ../beta/static/

# moving index file to templates
cp ./index.html ../beta/static/templates/
