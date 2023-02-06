# RUN to Regenerate gRPC code

```zsh
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  ./stub_proto/service-1.proto
```


* notes
```text
you need to run the protoc-compiler yourself 
to generate the proto-file to the language do 
you need to use. Btw, the right thing to do is 
to make a different repository for the proto-file. 
Don't make the proto-file in the same place/repository
with the service repository
```