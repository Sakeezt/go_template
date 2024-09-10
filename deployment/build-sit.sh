#!bin/bash

NAME="go-template"
docker build -f ./deployment/Dockerfile-sit -t $NAME:sit .