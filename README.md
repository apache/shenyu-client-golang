# shenyu-client-golang [中文](./README_CN.md) #

![license](https://img.shields.io/badge/license-Apache--2.0-green.svg)

---

## Shenyu-client-golang

Shenyu-client-golang for Go client allows you to access ShenYu Gateway,it supports registory go service to ShenYu
Gateway.

## Requirements

Supported Go version **over 1.12**

Supported ShenYu version **over 2.4.3**

## Installation

Use `go get` to install SDK：

```sh
$ go get -u github.com/apache/incubator-shenyu-client-golang
```

## How to use

**1.Fist make sure The ShenYuAdmin is Started, and ShenYuAdmin service active port is 9095.**

**2.Get shenyu_admin_client. (Register service need this)**

```sh
//init ShenYuAdminClient
adminClient := &model.ShenYuAdminClient{
    UserName: "admin",  //require user provide
    Password: "123456", //require user provide
}

adminTokenData, err := clients.NewShenYuAdminClient(adminClient)

The adminTokenData like this :
{
    "code":200,
    "message":"login dashboard user success",
    "data":{
        "id":"1",
        "userName":"admin",
        "role":1,
        "enabled":true,
        "dateCreated":"2018-06-23 15:12:22",
        "dateUpdated":"2022-03-09 15:08:14",
        "token":"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyTmFtZSI6ImFkbWluIiwiZXhwIjoxNjUwNjc5OTQ2fQ.K92Il2kmJ0X3FgjY4igW35-pw9nsf5VKdUyqBoyIaF4"
    }
}
```
