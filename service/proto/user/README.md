# Generate go source
```shell
protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. user.proto
```
