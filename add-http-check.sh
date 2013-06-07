#!/bin/bash

usage() {
    echo "Usage:"
    echo -e "\t$0 <key> <url> <interval>\n"
    echo "Example:"
    echo -e "\t$0 my_website http://example.org 60s\t\tWill check http://example.org every minute"
    echo -e "\t$0 other_website https://example.org 1h\tWill check https://example.org every hour"

    exit 1
}

KEY=$1
URL=$2
INTERVAL=$3

if [[ -z $1  || -z $2 || -z $3 ]];
then
    usage;
fi

HOST=`heroku apps:info | awk '/Web URL:/ {print $3}'`

JSON='{"key": "__key__", "type": "http", "interval": "__interval__", "alert": false, "notifyFix": false, "alertDelay": "0s", "config": {"url": "__url__"}}'

CONFIG=`echo $JSON | sed "s#__key__#$KEY#" | sed "s#__interval__#$INTERVAL#" | sed "s|__url__|$URL|"`
echo $CONFIG | curl -i -X POST $HOST/checks --data "@-"
