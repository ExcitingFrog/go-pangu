# Golang-pangu

[简明教程](https://github.com/ExcitingFrog/go-pangu/blob/master/document/%E7%AE%80%E6%98%8E%E6%95%99%E7%A8%8B.md)<br>
[English document](https://github.com/ExcitingFrog/go-pangu/blob/master/README.md)<br>
Golang-pangu是一个用Go语言开发的一站式服务开发的技术解决方案，它整合了 JWT(JSON Web Tokens), Postgres, Redis, Docker, Gin, Ginkgo, Gorm等多项业界尖端技术，它是后端项目开发的起点，可作为开发者强有力的生产工具。

## 产品特性
Golang-pangu基于以下的工具

|名字|描述|
|------|--------|
|[Go](https://github.com/golang/go)|最近几年最为流行的新兴语言，简单的同时拥有极高的并发性能。|
|[Gin](https://github.com/gin-gonic/gin)|基于Go语言的web框架，方便灵活的中间件，强大的数据绑定，以及极高的性能|
|[Gorm](https://github.com/go-gorm/gorm)|Go语言的全功能ORM库，用于操作数据库|
|[Ginkgo](https://github.com/onsi/ginkgo)|Ginkgo是一个BDD风格的Go测试框架，旨在帮助你有效地编写富有表现力的全方位测试。|
|[JWT](https://jwt.io/)|JSON Web Tokens，是目前最流行的跨域认证解决方案。|
|[Postgres](https://www.postgresql.org/)|高性能开源数据库，当整体负载达到很高时依旧能有强大的性能|
|[Redis](https://redis.io/)|内存数据库，拥有极高的性能|
|[Docker](https://www.docker.com/)|开发、部署、运行应用的虚拟化平台|

## 整体结构
```
.
├── application.yml  
├── args
│   ├── args.go
│   └── cmd.go
├── conf  
│   ├── conf_debug.go
│   ├── conf.go
│   └── conf_release.go
├── controller
│   ├── application.go
│   ├── auth.go
│   ├── error.go
│   └── session.go
├── db  
│   └── db.go
├── Dockerfile
├── go.mod
├── go.sum
├── jwt  
│   └── jwt.go
├── main.go
├── Makefile  
├── middleware  
│   └── middleware.go
├── models  
│   ├── base_model.go
│   └── user.go
├── params  
│   └── params.go
├── README.md
├── redis
│   └── redis.go
├── routers  
│   └── router.go
├── test
│   ├── sign_in_test.go
│   └── test_suite_test.go
└── util
    └── util.go
```

|文件|功能|
|------|--------|
|application.yml|配置文件，包含基本信息|
|[args](https://github.com/ExcitingFrog/go-pangu/tree/master/args)|包含获取url的params的函数|
|[conf](https://github.com/ExcitingFrog/go-pangu/tree/master/conf)|获取配置文件的函数|
|[controller](https://github.com/ExcitingFrog/go-pangu/tree/master/controller)|router使用的handler控件，包含各种操作具体内容|
|[db](https://github.com/ExcitingFrog/go-pangu/tree/master/db)|db操作，像是打开数据库|
|[jwt](https://github.com/ExcitingFrog/go-pangu/tree/master/jwt)|jwt相关内容 包含生成jwt与验证jwt的函数|
|main.go|程序主函数，执行时增加-db参数可选择不同的内容，create创建数据库，migrate更新表结构，drop删除数据库|
|[middleware](https://github.com/ExcitingFrog/go-pangu/tree/master/middleware)|中间件，验证token是否正确|
|[models](https://github.com/ExcitingFrog/go-pangu/tree/master/models)|基础的结构以及一些基本的数据库操作|
|[params](https://github.com/ExcitingFrog/go-pangu/tree/master/params)|数据绑定的结构|
|[redis](https://github.com/ExcitingFrog/go-pangu/tree/master/redis)|包含连接redis和redis操作函数|
|[router](https://github.com/ExcitingFrog/go-pangu/tree/master/routers)|路由|
|[test](https://github.com/ExcitingFrog/go-pangu/tree/master/test)|测试|
|[i18n](https://github.com/ExcitingFrog/go-pangu/tree/master/i18n)|国际化包|
|[influx](https://github.com/ExcitingFrog/go-pangu/tree/master/influx)|influx数据库操作|

## 开始运行
1. 安装postgres和redis数据库
2. 配置根目录下的 **application.yml** 配置文件
3.
```sh
`make create`(创建数据库) 或者 `go run main.go -db=create`
`make migrate`(初始数据库表) 或者 `go run main.go -db=migrate`
`make watch`(支持热更新) 或者 `go run main.go`
```
4. 在浏览器打开 `http://localhost:3002/ping` 会显示pong，表明服务成功部署



## Api 样例

* ### sign_up (用户注册)
  Post `http://localhost:3002/users/sign_up`

  params: email, password, password_confirm


* ### sign_in (用户登录)
    Post `http://localhost:3002/users/sign_in`

    params: email, password, login_type, DEVICE_TYPE

  成功后会在头部返回Authorization，这是后续用户接口需要的token

* ### auth_ping
    Get `http://localhost:3002/auth_ping`

  需要user的token的接口

* ### change_password
    Post `http://localhost:3002/users/change_password`

  修改用户密码，需要user的token

* ### cities

    Post `http://localhost:3002/cities`

    params: language

    设置成en时，返回英文城市列表，zh时为中文城市列表


### sms api
腾讯云服务，需要自行在application配置文件设置key

* ### sms

    Get `http://localhost:3002/sms`

    params: mobile

    发送短信接口

### influx apis
需要安装influxdb后，修改main.go然后执行初始化操作

* ### influx_save

    Post `http://localhost:3002/influx_save`

    params: user_name, local, version

    将结构化数据保存在influxdb数据库

* ### influx_show

    Post `http://localhost:3002/influx_save`

    读取influxdb数据库

### pay apis
支付宝支付接口，需要自行在application配置文件设置key
* ### alipay

    Post `http://localhost:3002/alipay`

    创建支付宝账单

* ### alipay_notify

    Post `http://localhost:3002/alipay_notify`

    接收支付信息的回调
