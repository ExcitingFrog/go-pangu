# Golang-pangu
[中文文档](https://github.com/ExcitingFrog/go-pangu/blob/master/READMECN.md)

Go-pangu is a Go boilerplate which follows cutting-edge solutions already adopted by the industry,  JWT(JSON Web Tokens), Postgres, Redis, Docker,  Gin, Ginkgo, Gorm. It is a solid production-ready starting point for your new backend projects.

## Features
Golang-pangu is based on following tools

|name|description|
|------|--------|
|[Go](https://github.com/golang/go)|an open source programming language that makes it easy to build simple, reliable, and efficient software.|
|[Gin](https://github.com/gin-gonic/gin)|web struct based on Go, flexible middleware，strong data binding and outstanding performance.|
|[Gorm](https://github.com/go-gorm/gorm)|The fantastic ORM library for Golang aims to be developer friendly.|
|[Ginkgo](https://github.com/onsi/ginkgo)|Ginkgo builds on Go's testing package, allowing expressive Behavior-Driven Development ("BDD") style tests.|
|[JWT](https://jwt.io/)|JSON Web Tokens. An open, industry standard RFC 7519 method for representing claims securely between two parties.|
|[Postgres](https://www.postgresql.org/)|The world's most advanced open source relational database|
|[Redis](https://redis.io/)|An open source (BSD licensed), in-memory data structure store, used as a database, cache and message broker.|
|[Docker](https://www.docker.com/)|Docker is a tool designed to make it easier to create, deploy, and run applications by using containers.|

## Struct
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

|file|function|
|------|--------|
|application.yml|config file|
|[args](https://github.com/ExcitingFrog/go-pangu/tree/master/args)|functions which can fetch params from request url|
|[conf](https://github.com/ExcitingFrog/go-pangu/tree/master/conf)|functions which can get configurations|
|[controller](https://github.com/ExcitingFrog/go-pangu/tree/master/controller)|handlers|
|[db](https://github.com/ExcitingFrog/go-pangu/tree/master/db)|database operations like migrating database|
|[jwt](https://github.com/ExcitingFrog/go-pangu/tree/master/jwt)|fuctions to create and check jwt|
|main.go|main function.Call function with "--db" parameter, "create" to create database, "migrate" to migrate tables, "dorp" to delete database|
|[middleware](https://github.com/ExcitingFrog/go-pangu/tree/master/middleware)|middleware|
|[models](https://github.com/ExcitingFrog/go-pangu/tree/master/models)|base models and basic database operations|
|[params](https://github.com/ExcitingFrog/go-pangu/tree/master/params)|struct used in data binding|
|[redis](https://github.com/ExcitingFrog/go-pangu/tree/master/redis)|redis connection and operations|
|[router](https://github.com/ExcitingFrog/go-pangu/tree/master/routers)|router|
|[test](https://github.com/ExcitingFrog/go-pangu/tree/master/test)|test|
|[i18n](https://github.com/ExcitingFrog/go-pangu/tree/master/i18n)|internationalization|
|[influx](https://github.com/ExcitingFrog/go-pangu/tree/master/influx)|influx operations include read/save point|


## Start

1. install postgres, redis
2. config application.yml
3.
```sh
`make create`(create database) or `go run main.go -db=create`
`make migrate`(migrate tables) or `go run main.go -db=migrate`
`make watch`(with hot reload) or `go run main.go`
```
4. open `http://localhost:3002/ping` in web browser, and then you will get a "pong" response

## Api examples

* ### sign_up

  Post `http://localhost:3002/users/sign_up`

  params: email, password, password_confirm

  Register user

* ### sign_in

  Post `http://localhost:3002/users/sign_in`

  params: email, password, DEVICE_TYPE, login_type

  You will get a header with authorization parameter from response after logging in successfully

* ### auth_ping

  Get `http://localhost:3002/auth_ping`

  Should add a valid user token to request this api

* ### change_password

  Post `http://localhost:3002/users/change_password`

  params: origin_password, password, password_confirm

  Modify user's password, which needs authorization

* ### cities

    Post `http://localhost:3002/cities`

    params: language

    set language to en, return cities in English.
    set language to zh, return cities in Chinese.


### sms api
(Tencent service, need to set your key in application.yml )

* ### sms

  Get `http://localhost:3002/sms`

  params: mobile

  send sms

### influx apis
(need to install influxdb and modify main.go)

* ### influx_save

  Post `http://localhost:3002/influx_save`

  params: user_name, local, version

  save struct in influxdb

* ### influx_show

  Post `http://localhost:3002/influx_save`

  get struct message in influxdb

### pay apis
(Alipay service, need to set you key in application.yml)
* ### alipay

  Post `http://localhost:3002/alipay`

  create alipay bill

* ### alipay_notify

    Post `http://localhost:3002/alipay_notify`

    receive pay details notify
