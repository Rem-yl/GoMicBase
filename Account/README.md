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

## 思路&步骤
1. 新建mysql的docker-compose
   - mysql的默认端口是3306
   - mysql的默认数据保存位置在/var/lib/mysql
   - `lsof -i :3306` : 查看3306端口是否被占用
   - `lsof -i -P -n | grep LISTEN` : 查看所有被占用的端口

2. 构建Account GRRC服务
3. 构建Account Web服务
