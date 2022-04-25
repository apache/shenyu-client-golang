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

**2.Step 1 Get shenyu_admin_client. (Register service need this)**

```sh
//init ShenYuAdminClient
adminClient := &model.ShenYuAdminClient{
    UserName: "admin",  //require user provide
    Password: "123456", //require user provide
}

adminToken, err := clients.NewShenYuAdminClient(adminClient)

The adminToken like this :
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

When you success get toekn, you will see this :
this is ShenYu Admin client token -> eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyTmFtZSI6ImFkbWluIiwiZXhwIjoxNjUwNjc5OTQ2fQ.K92Il2kmJ0X3FgjY4igW35-pw9nsf5VKdUyqBoyIaF4

```


**3.Step 2 Register MetaData to ShenYu GateWay. (Need step 1 token to invoke)**
```sh
//MetaDataRegister(Need Step 1 toekn adminToken.AdminTokenData)
metaData := &model.MetaDataRegister{
		AppName: "testMetaDataRegister", //require user provide
		Path:    "/your/path",           //require user provide
		Enabled: true,                   //require user provide
		Host:    "127.0.0.1",            //require user provide
		Port:    "8080",                 //require user provide
	}
	result, err := clients.RegisterMetaData(adminToken.AdminTokenData, metaData)
	if err != nil {
		logger.Warn("MetaDataRegister has error:",err)
	}
	logger.Info("finish register metadata ,the result is->", result)
	
	
When Register success , you will see this :  
finish register metadata ,the result is-> true
```
