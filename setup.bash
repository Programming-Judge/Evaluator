#!/bin/bash
go mod download

docker volume create judge-submissions
cd images 
cd op-bash
docker build --tag --tag op-bash .
cd ../python3
docker build --tag python3-eval .