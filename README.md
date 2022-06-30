# shenyu-client-golang

English | [简体中文](README_CN.md)

[![OSCS Status](https://www.oscs1024.com/platform/badge/apache/incubator-shenyu-client-golang.svg?size=small)](https://www.oscs1024.com/project/apache/incubator-shenyu-client-golang?ref=badge_small)
[![Build and Test](https://github.com/apache/incubator-shenyu-client-golang/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/apache/incubator-shenyu-client-golang/actions)
[![codecov.io](https://codecov.io/gh/apache/incubator-shenyu-client-golang/coverage.svg?branch=main)](https://app.codecov.io/gh/apache/incubator-shenyu-client-golang?branch=main)
[![GoDoc](https://godoc.org/github.com/apache/incubator-shenyu-client-golang?status.svg)](https://godoc.org/github.com/apache/incubator-shenyu-client-golang)

---

## Shenyu-client-golang

Shenyu-client-golang for Go client allows you to access ShenYu Gateway,it supports registory go service to ShenYu
Gateway.

---
## Supported Register Center to ShenYu Gateway

* **Http type Register**
* **Nacos type Register**
* **Zookeeper type Register**

---

## Requirements

Supported Go version **over 1.12**

SDK Supported ShenYu version **over 2.4.3**

## Installation

Use `go get` to install SDK：

```sh
$ go get -u github.com/apache/incubator-shenyu-client-golang
```

## The Demo location

* incubator-shenyu-client-golang/example/**_client/main.go
---

## The Http type Register

**1.Fist make sure The ShenYuAdmin is Started, and ShenYuAdmin service active port is 9095.**
```go
Or you will see this error :
	
2022-05-05 15:24:28 [WARN] [github.com/apache/incubator-shenyu-client-golang/example/http_client/main.go:53] MetaDataRegister has error: The errCode is ->:503, The errMsg is  ->:Please check ShenYu admin service status

caused by:
Post "http://127.0.0.1:9095/shenyu-client/register-metadata": dial tcp 127.0.0.1:9095: connect: connection refused
2022-05-05 15:24:28 [INFO] [github.com/apache/incubator-shenyu-client-golang/example/http_client/main.go:55] finish register metadata ,the result is-> false
2022-05-05 15:24:28 [WARN] [github.com/apache/incubator-shenyu-client-golang/example/http_client/main.go:68] UrlRegister has error: The errCode is ->:503, The errMsg is  ->:Please check ShenYu admin service status

caused by:
Post "http://127.0.0.1:9095/shenyu-client/register-uri": dial tcp 127.0.0.1:9095: connect: connection refused
2022-05-05 15:24:28 [INFO] [github.com/apache/incubator-shenyu-client-golang/example/http_client/main.go:70] finish UrlRegister ,the result is-> false
	
```

**2.Step 1 Get shenyu_admin_client. (Register service need this)**

```go
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
```go
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

**4.Step 3  Url  Register  to ShenYu GateWay. (Need step 1 token to invoke)**
```go
//URIRegister(Need Step 1 toekn adminToken.AdminTokenData)
//init urlRegister
	urlRegister := &model.URIRegister{
		Protocol:    "testMetaDataRegister", //require user provide
		AppName:     "testURLRegister",      //require user provide
		ContextPath: "contextPath",          //require user provide
		RPCType:     constants.RPCTYPE_HTTP, //require user provide
		Host:        "127.0.0.1",            //require user provide
		Port:        "8080",                 //require user provide
	}
	result, err = clients.UrlRegister(adminToken.AdminTokenData, urlRegister)
	if err != nil {
		logger.Warn("UrlRegister has error:", err)
	}
	logger.Info("finish UrlRegister ,the result is->", result)
```

## Entire Success log
```go
2022-05-05 15:43:56 [INFO] [github.com/apache/incubator-shenyu-client-golang/clients/admin_client/shenyu_admin_client.go:51] Get ShenYu Admin response, body is -> {200 login dashboard user success {1 admin 1 true 2018-06-23 15:12:22 2022-03-09 15:08:14 eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyTmFtZSI6ImFkbWluIiwiZXhwIjoxNjUxODIzMDM2fQ.gMzPKaNlXEd1Q517qQamOpg358W9L0-0cZN3lkk06WE}}
2022-05-05 15:43:56 [INFO] [github.com/apache/incubator-shenyu-client-golang/example/http_client/main.go:40] this is ShenYu Admin client token -> eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyTmFtZSI6ImFkbWluIiwiZXhwIjoxNjUxODIzMDM2fQ.gMzPKaNlXEd1Q517qQamOpg358W9L0-0cZN3lkk06WE
2022-05-05 15:43:57 [INFO] [github.com/apache/incubator-shenyu-client-golang/example/http_client/main.go:55] finish register metadata ,the result is-> true
2022-05-05 15:43:57 [INFO] [github.com/apache/incubator-shenyu-client-golang/example/http_client/main.go:70] finish UrlRegister ,the result is-> true

```



---
## The Nacos type Register

**1.Fist make sure your nacos env is correct,the set this necessary param.**
```go
//set nacos env configuration
    ncp := &nacos_client.NacosClientParam{
        IpAddr:      "console.nacos.io",
        Port:        80,
        NamespaceId: "e525eafa-f7d7-4029-83d9-008937f9d468",
}
```

**2. Prepare your service metaData to register**
```go
//metaData is necessary param, this will be register to shenyu gateway to use
    metaData := &model.URIRegister{
        Protocol:    "testMetaDataRegister", //require user provide
        AppName:     "testURLRegister",      //require user provide
        ContextPath: "contextPath",          //require user provide
        RPCType:     constants.RPCTYPE_HTTP, //require user provide
        Host:        "127.0.0.1",            //require user provide
        Port:        "8080",                 //require user provide
}
    metaDataStringJson, _ := json.Marshal(metaData)
```

**3. Prepare your service Instance message(include metaData)**
```go
//init NacosRegisterInstance
    nacosRegisterInstance := vo.RegisterInstanceParam{
        Ip:          "10.0.0.10", //require user provide
        Port:        8848,        //require user provide
        ServiceName: "demo.go",   //require user provide
        Weight:      10,          //require user provide
        Enable:      true,        //require user provide
        Healthy:     true,        //require user provide
        Ephemeral:   true,        //require user provide
        Metadata:    map[string]string{"contextPath": "contextPath", "uriMetadata": string(metaDataStringJson)},
}
```

**4.use client to invoke RegisterNacosInstance**
```go
    client, err := nacos_client.NewNacosClient(ncp)
        if err != nil {
        logger.Fatal("create nacos client error : %+V", err)
}

    registerResult, err := nacos_client.RegisterNacosInstance(client, nacosRegisterInstance)
        if !registerResult && err != nil {
        logger.Fatal("Register nacos Instance error : %+V", err)
}
        //do your logic
```

## Entire Success log
```go
2022-06-27 10:56:17 [INFO] [github.com/incubator-shenyu-client-golang/clients/nacos_client/nacos_client.go:92] RegisterServiceInstance,result:true

,param:{Ip:10.0.0.10 Port:8848 Weight:10 Enable:true Healthy:true Metadata:map[contextPath:contextPath uriMetadata:{"protocol":"testMetaDataRegister","appName":"testURLRegister","contextPath":"contextPath","rpcType":"http","host":"127.0.0.1","port":"8080"}] ClusterName: ServiceName:demo.go GroupName: Ephemeral:true}

```



---
## The Zookeeper type Register

**1.Fist make sure your Zookeeper env is correct,the set this necessary param.**
```go
    servers := []string{"127.0.0.1:2181"}         //require user provide
        client, err := NewClient(servers, "/api", 10) //zkRoot require user provide
        if err != nil {
            panic(err)
         }
        defer client.Close()
```

**2. Prepare your service metaData to register**
```go
//metaData is necessary param, this will be register to shenyu gateway to use
        //init MetaDataRegister
        metaData1 := &model.MetaDataRegister{
            AppName: "testMetaDataRegister", //require user provide
            Path:    "your/path1",           //require user provide
            Enabled: true,                   //require user provide
            Host:    "127.0.0.1",            //require user provide
            Port:    "8080",                 //require user provide
        }
```

**3.use client to invoke RegisterNacosInstance**
```go
   //register multiple metaData
    if err := client.RegisterNodeInstance(metaData1); err != nil {
        panic(err)
    }
        //do your logic
```

**4.use client to invoke DeleteNodeInstance**
```go
    //your can chose to invoke,not require
    err = client.DeleteNodeInstance(metaData1)
     if err != nil {
       panic(err)
}
```

**5.use client to get zk nodes**
```go
    //range nodes
    for index, node := range nodes {
        nodeJson, err := json.Marshal(node)
        if err == nil {
        logger.Info("GetNodesInfo ,success Index", index, string(nodeJson))
    }
}
```

## Entire Success log
```go
2022-06-28 15:21:57 [INFO] [github.com/incubator-shenyu-client-golang/example/zk_client/zk_client.go:80] GetNodesInfo ,success Index 0 {"appName":"testMetaDataRegister","path":"your/path1","rpcType":"","enabled":true,"host":"127.0.0.1","port":"8080","pluginNames":null,"registerMetaData":false,"timeMillis":0}
2022-06-28 15:21:57 [INFO] [github.com/incubator-shenyu-client-golang/example/zk_client/zk_client.go:80] GetNodesInfo ,success Index 1 {"appName":"testMetaDataRegister","path":"your/path3","rpcType":"","enabled":true,"host":"127.0.0.1","port":"8282","pluginNames":null,"registerMetaData":false,"timeMillis":0}
2022-06-28 15:21:57 [INFO] [github.com/incubator-shenyu-client-golang/example/zk_client/zk_client.go:80] GetNodesInfo ,success Index 2 {"appName":"testMetaDataRegister","path":"your/path2","rpcType":"","enabled":true,"host":"127.0.0.1","port":"8181","pluginNames":null,"registerMetaData":false,"timeMillis":0}

```
