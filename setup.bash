#!/bin/bash

go mod tidy
go mod download

docker volume create judge-submissions
cd images/python3
docker build --tag python3-eval .