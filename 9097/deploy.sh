#! /bin/bash

echo Building container with port 9097

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main app.go

docker build -t my_http_server_9097 .

docker run -it -p 9097:9097  --name my_http_server_9097 my_http_server_9097
