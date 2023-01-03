# blog_go

A microservice module implemented by gin and gRPC framework

it can provide as a server in gRPC and be called from spring client.

## Update Log

22/12/31 Initial Commit

How to generate proto.go file?

- --go_out：表示要生成的go文件放在哪里
- -go_opt=paths=source_relative：表示路径的表达方式是相对路径
- --go-grpc_out：同go_out
- -go-grpc_opt=paths=source_relative：同上
- proto/xxx.proto ：表示生成的源proto文件

```
protoc  --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/xxx.proto
```

 
