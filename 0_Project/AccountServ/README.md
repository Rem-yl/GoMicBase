# README

## Account 
1. 设计`model.Account`的数据结构;
2. 数据库代码写在`internal/db.go`中, 使用`gorm`连接Mysql数据库, 并初始化`account`表;
3. 使用`protobuf`定义`AccountService`grpc服务, 想好需要的服务以及互相传递的消息;
4. `biz`包写具体的`AccountService`grpc服务逻辑;
5. 写`biz`包的测试用例;

## Run Command
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