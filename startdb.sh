#!/bin/bash

docker network create mongoCluster

sleep 2

docker run -d --rm -p 27017:27017 --name mongo1 --network mongoCluster mongo:4.4 mongod --replSet myReplicaSet --bind_ip localhost,mongo1 \
  && docker run -d --rm -p 27018:27017 --name mongo2 --network mongoCluster mongo:4.4 mongod --replSet myReplicaSet --bind_ip localhost,mongo2 \
    && docker run -d --rm -p 27019:27017 --name mongo3 --network mongoCluster mongo:4.4 mongod --replSet myReplicaSet --bind_ip localhost,mongo3

sleep 5

docker exec -it mongo1 mongo --eval "rs.initiate({
 _id: \"myReplicaSet\",
 members: [
   {_id: 0, host: \"mongo1\"},
   {_id: 1, host: \"mongo2\"},
   {_id: 2, host: \"mongo3\"}
 ]
})"

# Verifique o Replica Set
# docker exec -it mongo1 mongo --eval "rs.status()"