# gRPC-gateway example

*a simple example of client and server operation via gRPC-gateway*

## Run

- Open directory

```shell
cd grpc-gateway/
```

- clone repository:

```shell
git clone https://github.com/googleapis/googleapis.git
```

- Run server

```shell
go run server.go
```

- Run client

```shell
go run client.go
```

- Get a requests:

```shell
curl -v --http2-prior-knowledge http://localhost:8080/v1/users/322

* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
* [HTTP/2] [1] OPENED stream for http://localhost:8080/v1/users/322
* [HTTP/2] [1] [:method: GET]
* [HTTP/2] [1] [:scheme: http]
* [HTTP/2] [1] [:authority: localhost:8080]
* [HTTP/2] [1] [:path: /v1/users/322]
* [HTTP/2] [1] [user-agent: curl/8.7.1]
* [HTTP/2] [1] [accept: */*]
> GET /v1/users/322 HTTP/2
> Host: localhost:8080
> User-Agent: curl/8.7.1
> Accept: */*
> 
* Request completely sent off
< HTTP/2 200 
< content-type: application/json
< grpc-metadata-content-type: application/grpc
< content-length: 60
< date: Mon, 09 Jun 2025 13:41:24 GMT
< 
* Connection #0 to host localhost left intact
{"id":"322", "name":"John Doe", "email":"johndoe@gmail.com"}
```

```shell
curl -v --http2-prior-knowledge http://localhost:8080/v1/users
```

```text
gateway

2025/06/09 16:43:09 gRPC-Gateway listening on :8080
```

```text
server

2025/06/09 16:43:01 gRPC core listening on :50051
2025/06/09 16:43:46 Received GetUser request for ID=322
2025/06/09 16:43:46 Request 322
```

## Notes

- To change the file structure `user.proto` and logic changes, delete directory `example.com/`
- Change file `proto/user.proto`
- If protoc not installed, install:

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
```

- Add PATH

```shell
export PATH="$PATH:$(go env GOPATH)/bin"
```

- Check plugins:

```shell
which protoc-gen-go
which protoc-gen-go-grpc
```

- Generate code again:

```shell
protoc \
  -I proto \
  -I ./googleapis \
  --go_out=. \
  --go-grpc_out=. \
  --grpc-gateway_out=. \
  proto/user.proto
```
