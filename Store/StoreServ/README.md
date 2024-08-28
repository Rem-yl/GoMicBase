# 商场微服务
## 开发过程
1. 读取本地yaml配置
2. gorm初始化数据库
   1. gorm可以使用`foreignKey`来显式指定外键;
   2. 当没有显式指定外键时, gorm会根据字段名来推断外键, 例如`BrandID`, gorm会默认该字段指向Brand表中的ID字段