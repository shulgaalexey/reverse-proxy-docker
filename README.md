# reverse-proxy-docker

Reverse Proxy example on Docker containers

**Note**
Only http is supported


```

      +---------------+
      | Reverse Proxy |
      | 172.17.0.4    |
      | port 9096     |
      +---------------+
        |
        |
        | redirected from http://localhost:9096/serviceone/v1/blahblah
        |
        |           +------------+
        +---------->| Service v1 |
        |           | 172.17.0.2 |
        |           | port 9097  |
        |           +------------+
        |
        |
        |
        | redirected from http://localhost:9096/serviceone/v2/foo/bar
        |
        |           +------------+
        +---------->| Service v2 |
                    | 172.17.0.3 |
                    | port 9098  |
                    +------------+


```


## How to run

```
$ ./9097/deploy.sh
$ ./9098/deploy.sh
$ ./rproxy/deploy.sh
```

## How to use

Services resond with a simple html page, so you can use
web browser or curl.

### Access services through the reverse proxy (recommended)

```
curl http://localhost:9096/serviceone/v1
curl http://localhost:9096/serviceone/v2
```

### Access services directly (debug)

```
curl http://localhost:9097/serviceone/v1
curl http://localhost:9098/serviceone/v2
```


## Troubleshooting

Check that services are assigned with IPs:

```
"Name": "my_http_server_9097"
"IPv4Address": "172.17.0.2/16"
```

and

```
"Name": "my_http_server_9098"
"IPv4Address": "172.17.0.3/16"
```

## Cleanup

```
docker container prune
```


## Reference

Based on https://blog.charmes.net/post/reverse-proxy-go/
