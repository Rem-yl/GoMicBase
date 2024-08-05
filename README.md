# 说明
《从0到Go语言微服务架构师》课程学习记录

## 账户grpc服务-AccountServe
### Account 
1. 设计`model.Account`的数据结构;
2. 数据库代码写在`internal/db.go`中, 使用`gorm`连接Mysql数据库, 并初始化`account`表;
3. 使用`protobuf`定义`AccountService`grpc服务, 想好需要的服务以及互相传递的消息;
4. `biz`包写具体的`AccountService`grpc服务逻辑;
5. 写`biz`包的测试用例;

### Run Command
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

## 账户Web服务-AccountWeb
### Gin使用router+handler调grpc服务
```bash
cd AccountWeb
go run main.go
curl "http://127.0.0.1:8081/account/list"
```

### 用户登录: Gin+JWT
**JWT**
1. 配置;
2. 颁发JWT;
3. 解析JWT;
4. 刷新JWT


## 日志
### 2024-08-05
1. 账户验证登录;
2. JWT相关

### 2024-07-15
1. 学习Ch03
2. 下载RapidAPI

