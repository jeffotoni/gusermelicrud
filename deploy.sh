#!/bin/bash
# Go Api server
# @jeffotoni
echo "-------------------------------------- Clean <none> images ---------------------------------------"
docker rmi $(docker images | grep "<none>" | awk '{print $3}') --force
echo "\033[0;33m######################## carregando todos serviços ########################033[0m"
cd userwrite
echo $(pwd)
make build
sh deploy.sh

cd ../userget
make build
sh deploy.sh

cd ../

docker-compose up -d
sh mongo.script/script.start.sh
docker-compose ps


echo "\033[0;33m Prontinho .... \033[0m"