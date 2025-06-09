# gRPC example

*a simple example of client and server operation via gRPC*

## Run

- Open directory

```shell
cd grpc-keyvalue/
```

- Run server

```shell
go run server.go
```

- Run client

```shell
go run client.go
```

- Input data

```text
# client

> put foo bar
2025/06/09 16:16:14 Args: [put foo baz]
OK
> get foo
2025/06/09 16:16:31 Args: [get foo]
Value: baz
```

```text
# server

2025/06/09 16:09:08 Starting core
2025/06/09 16:09:27 Received PUT key=foo value=baz
2025/06/09 16:09:32 Received GET key=foo
```

## Notes

- To change the file structure `keyvalue.proto` and logic changes, delete files: `proto/keyvalue.pb.go` and `proto/keyvalue_grpc.pb.go`
- Change file `proto/keyvalue.proto`
- If protoc not installed, install:

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
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
--proto_path=grpc/kvPR \
--go_out=proto/keyvalue.proto \
--go_opt=paths=source_relative \
--go-grpc_out=proto/keyvalue.proto \
--go-grpc_opt=paths=source_relative proto/keyvalue.proto
```