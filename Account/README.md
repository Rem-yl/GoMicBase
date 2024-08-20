# Account

账户服务
## 资料
- [如何在Go中使用JWT](https://juejin.cn/post/7093035836689612836)

## 目录架构

- AccountServ : 账户微服务
- AccountWeb : 账户web服务
- Conf : 项目配置
- Log : 日志配置
- Share : 全局变量，包括err等

## 启动docker
```bash
cd Account
docker-compose up -d
```

## 思路&步骤
1. 新建mysql的docker-compose
   - mysql的默认端口是3306
   - mysql的默认数据保存位置在/var/lib/mysql
   - `lsof -i :3306` : 查看3306端口是否被占用
   - `lsof -i -P -n | grep LISTEN` : 查看所有被占用的端口

2. 构建Account GRRC服务
3. 构建Account Web服务
4. JWT功能测试
   1. 运行服务
      ```bash
         # shell 1
         cd AccountServ
         go run main.go

         # shell 2
         cd AccountWeb
         go run main.go
      ```
   2. 在RapidAPI中使用POST请求 127.0.0.1:8080/account/login, 获得返回的JWT token
   3. 在RapidAPI中使用GET请求 127.0.0.1:8080/account/jwt_test, 请求头的Auth字段天填写刚刚获得的JWT token
