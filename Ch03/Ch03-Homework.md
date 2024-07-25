# HomeWork
### 1. 介绍一下Gin
Gin 是用Go 语言编写的web 框架,它可以快速实现 API,常用于生产环境中。现在流行前后端分离,即前端对应的是网页、App,后端对应的是业务逻辑。

表现层状态转移(Representational State Transfer, REST)是由 Roy Fielding 提出的一种软件架构风格,它由一系列规范组成,满足这些规范的 API 均可称为 RESTful API
1. REST 中的实体都被抽象成资源,每一个资源都有唯一的标识——URI,所有的行为都应在资源上实现 CURD(创建、修改、查询、删除)操作。
2. 每个 RESTful API 请求都包含了所有足够完成这次操作的信息,服务端无须保持 session 信息。其好处是服务端可以方便地进行弹性扩容。

HTTP 是 RESTful API 的实现标准。HTTP 中的 GET、POST、PUT 和 DELETE 方法对应REST 资源的获取、创新、更新、删除操作。
POST : 创建一个新资源
GET : 获取一个具体的资源
PUT : 更新一个资源
DELETE : 删除一个资源 