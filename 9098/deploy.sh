#! /bin/bash

echo Building container with port 9098

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main app.go

docker build -t my_http_server_9098 .

docker run -it -p 9098:9098 --name my_http_server_9098 my_http_server_9098
