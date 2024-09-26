# README

## 服务启动
```bash
cd deploy/docker_compose
docker-compose up -d

curl -I http://127.0.0.1:8848/nacos
curl -I http://127.0.0.1:8500/ui/dc1/services
```