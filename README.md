# Doc

### Generate proto
```
protoc \      
--go_out=./ \
--go_opt=paths=source_relative \
--go-grpc_out=./ --go-grpc_opt=paths=source_relative \
./proto/blogengine.proto
```

### Run tests
`go test ./...`

### Run server
`go run main.go`

### Client
Postman has a service discovery