# RUN to Regenerate gRPC code

```zsh
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  ./stub_proto/service-1.proto
```