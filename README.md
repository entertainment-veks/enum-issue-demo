# How to reproduce?
0. Install Go...
1. Install dependencies
```bash
go install github.com/bufbuild/buf/cmd/buf@latest

go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```
2. Generate code
```bash
buf generate --template api/buf.gen.yaml
```
3. Run application
```bash
cd cmd 
go run .
```
4. Make a grpcurl request to get value that behavies as expected
```bash 
grpcurl -proto api/swagger.proto -plaintext localhost:50051 swagger.Swagger.GetOk
```
>{
>  "value": "OK"
>}
5. Make a grpcurl request to get responce with lost value
```bash 
grpcurl -proto api/swagger.proto -plaintext localhost:50051 swagger.Swagger.GetLost
```
>{
>  
>}