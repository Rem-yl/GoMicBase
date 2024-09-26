# README

## 服务启动
```bash
cd deploy/docker_compose
docker-compose up -d

curl -I http://127.0.0.1:8848/nacos
curl -I http://127.0.0.1:8500/ui/dc1/services
```

## 问题记录
**nacos-go-sdk调用PublishConfig时报错 `NacosException: Client not connected,current status:STARTING`**
- [github-issue](https://github.com/alibaba/nacos/issues/6154)

核心原因是nacos2.x版本使用grpc调用服务，需要多开两个端口(在主端口8848上加固定偏移量1000/1001), 当9848端口无法使用时就会报这个错
```bash
lsof -i :9848
```

当有以下输出时, 说明端口正常开放
``` bash
COMMAND   PID USER   FD   TYPE             DEVICE SIZE/OFF NODE NAME
com.docke 892 yule  120u  IPv6 0xf2194499346310b9      0t0  TCP *:9848 (LISTEN)
```
