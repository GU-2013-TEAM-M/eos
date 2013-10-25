#!/bin/bash 

# moving static files
cp -r ./css ../prod/static/
cp -r ./js ../prod/static/
cp -r ./lib ../prod/static/

# moving index file to templates
cp ./index.html ../prod/static/templates/
