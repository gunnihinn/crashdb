#!/bin/bash

[[ -f crashdb ]] || go build -o crashdb main.go
rm -f core.*

GOTRACEBACK=crash ./crashdb &
curl -X POST --data '{"Key": "foo", "Value": 1}' localhost:8080/
curl -X POST --data '{"Key": "bar", "Value": 2}' localhost:8080/

type slapper || go get -u github.com/ikruglov/slapper
echo "Stop slapper by pressing 'q' once crashdb gets sad"
slapper -targets static/targets.txt -rate 10000

type dlv || go get -u github.com/derekparker/delve/cmd/dlv
dlv core ./crashdb core.*
