#!/bin/bash

[[ -f crashdb ]] || go build -o crashdb main.go
rm -f core.*


if ! type slapper &> /dev/null; then
    echo -n "slapper is not installed, should we install it?"
    read -p " [y/N]" -n 1 -r
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        echo
        go get -u -v github.com/ikruglov/slapper
    else
        exit 1
    fi
fi

if ! type dlv &> /dev/null; then
    echo -n "delve is not installed, should we install it?"
    read -p " [y/N]" -n 1 -r
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        echo
        go get -u -v github.com/derekparker/delve/cmd/dlv
    else
        exit 1
    fi
fi


echo "Killing previous crashdb instances"
killall crashdb
echo "Starting crashdb..."
GOTRACEBACK=crash ./crashdb &
echo -n  "Populating keys."
curl -X POST --data '{"Key": "foo", "Value": 1}' localhost:8080/
echo -n "."
curl -X POST --data '{"Key": "bar", "Value": 2}' localhost:8080/
echo "."

echo "Stop slapper by pressing 'q' once crashdb gets sad"
echo "Press enter to continue"
read

slapper -targets static/targets.txt -rate 10000

clear

if pidof systemd &> /dev/null; then
    echo "copying coredumps /var/lib/systemd/coredump/core.crashdb*"
    cp /var/lib/systemd/coredump/core.crashdb* .
fi

dlv core ./crashdb core.*
