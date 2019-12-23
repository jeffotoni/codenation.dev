#!/bin/bash

echo "--------------- deploy ------------------"
docker rmi $(docker images | grep "<nome>" | awk '{print $3}') --force
#GOOS=linux go build -ldflags="-s -w" -o apiserver main.go
docker build -f Dockerfile -t apiserver .

#docker push jeffotoni/apiserver 
#docker pull jeffotoni/apiserver



