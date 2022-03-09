## 电商微服务项目

此为后端项目

1. 后端项目地址 https://github.com/Dlimingliang/shop_srvs
2. 后台管理前端项目地址 https://github.com/Dlimingliang/shop-admin
3. 电商前端项目地址

## 项目配置及文档地址(私有)

https://github.com/Dlimingliang/shop-document

#### 技术选型

- 使用grpc进行微服务的调用，序列化使用protobuf
- 使用gin框架，并且创建restful风格API
- API文档 使用swagger构建自动化文档
- 使用viper进行配置读取
- 使用zap进行日志输出
- 注册中心 consul
- 配置中心 nacos
- 数据库 mysql
- 缓存 redis
- 其他后面再加入

#### 模块组成

1. 用户模块
2. 商品模块
3. 库存模块
4. 订单模块