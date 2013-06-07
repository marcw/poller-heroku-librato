#!/bin/bash
heroku create -b https://github.com/kr/heroku-buildpack-go.git
heroku addons:add librato
git push heroku master
heroku scale web=1
./add-http-check.sh google_com http://google.com 30s
./add-http-check.sh yahoo_com http://yahoo.com 30s
echo "Waiting 1 minute before opening librato dashboard"
sleep 60
heroku addons:open librato
