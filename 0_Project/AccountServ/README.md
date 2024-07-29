# README

- 启动docker
```bash
docker-compose up -d
```

- 生成proto文件
```bash 
protoc --go_out=. --go-grpc_out=. account.proto
```

- 运行测试文件
```bash
go test -v /Users/yule/Desktop/code/GoMicBase/0_Project/AccountServ/biz
```