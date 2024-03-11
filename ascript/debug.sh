#!/bin/bash
# npm run asbuild:debug
npx nodemon -e ts --watch assembly --exec "npm run asbuild:debug && npm start"
