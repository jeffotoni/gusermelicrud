#!/bin/bash
# Go Api server
# @jeffotoni
echo "-------------------------------------- Clean <none> images ---------------------------------------"
docker rmi $(docker images | grep "<none>" | awk '{print $3}') --force
echo "\033[0;33m################################## build docker userwrite ##################################\033[0m"
make build
docker build --no-cache -f Dockerfile -t jeffotoni/userwrite .
docker-compose build --no-cache 
echo "\033[0;33m Pode rodar agora: $ docker-compose up -d .......\033[0m"
#docker-compose up -d