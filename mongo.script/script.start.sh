#!/bin/bash
# Go Api server
# @jeffotoni

docker exec gusermeli_mongo-users_1 mongoimport --username=root --password=123456 --uri "mongodb://mongodb.local.com:27017/?authSource=admin" --type=json --mode=insert --db=meliusers --collection=users --file=/mnt/collection2.json

docker exec gusermeli_mongo-users_1 mongo --username root --password 123456 --authenticationDatabase admin --host mongodb.local.com --port 27017 --shell /mnt/index.js

echo "prontinho...."