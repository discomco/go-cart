#! /bin/bash

#docker-compose -f ./deploy/cmd/qr-cmd.yaml \
#               down


docker stop quadratic-cmd

docker rm quadratic-cmd

docker-compose -f ./deploy/networks.yaml \
               -f ./deploy/cmd/qr-cmd.yaml \
               up --build -d
