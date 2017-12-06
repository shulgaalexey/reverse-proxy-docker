#! /bin/bash

echo Building proxy container with port 9096

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main rproxy.go

docker build -t my_rproxy_9096 .

docker run -it -p 9096:9096  --name my_rproxy_9096 my_rproxy_9096
