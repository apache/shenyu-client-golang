# shenyu-client-golang

English | [简体中文](README_CN.md)

[![Build and Test](https://github.com/apache/shenyu-client-golang/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/apache/shenyu-client-golang/actions)
[![codecov.io](https://codecov.io/gh/apache/shenyu-client-golang/coverage.svg?branch=main)](https://app.codecov.io/gh/apache/shenyu-client-golang?branch=main)
[![GoDoc](https://godoc.org/github.com/apache/shenyu-client-golang?status.svg)](https://godoc.org/github.com/apache/shenyu-client-golang)

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
$ go get -u github.com/apache/shenyu-client-golang
```

## The Demo location

* shenyu-client-golang/example/**_client/main.go
---

## The Http type Register

**1.Fist make sure The ShenYuAdmin is Started, and ShenYuAdmin service active port is 9095.**
```go
Or you will see this error :
	
2022-05-05 15:24:28 [WARN] [github.com/apache/shenyu-client-golang/example/http_client/main.go:53] MetaDataRegister has error: The errCode is ->:503, The errMsg is  ->:Please check ShenYu admin service status

caused by:
Post "http://127.0.0.1:9095/shenyu-client/register-metadata": dial tcp 127.0.0.1:9095: connect: connection refused
2022-05-05 15:24:28 [INFO] [github.com/apache/shenyu-client-golang/example/http_client/main.go:55] finish register metadata ,the result is-> false
2022-05-05 15:24:28 [WARN] [github.com/apache/shenyu-client-golang/example/http_client/main.go:68] UrlRegister has error: The errCode is ->:503, The errMsg is  ->:Please check ShenYu admin service status

caused by:
Post "http://127.0.0.1:9095/shenyu-client/register-uri": dial tcp 127.0.0.1:9095: connect: connection refused
2022-05-05 15:24:28 [INFO] [github.com/apache/shenyu-client-golang/example/http_client/main.go:70] finish UrlRegister ,the result is-> false
	
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
2022-05-05 15:43:56 [INFO] [github.com/apache/shenyu-client-golang/clients/admin_client/shenyu_admin_client.go:51] Get ShenYu Admin response, body is -> {200 login dashboard user success {1 admin 1 true 2018-06-23 15:12:22 2022-03-09 15:08:14 eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyTmFtZSI6ImFkbWluIiwiZXhwIjoxNjUxODIzMDM2fQ.gMzPKaNlXEd1Q517qQamOpg358W9L0-0cZN3lkk06WE}}
2022-05-05 15:43:56 [INFO] [github.com/apache/shenyu-client-golang/example/http_client/main.go:40] this is ShenYu Admin client token -> eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyTmFtZSI6ImFkbWluIiwiZXhwIjoxNjUxODIzMDM2fQ.gMzPKaNlXEd1Q517qQamOpg358W9L0-0cZN3lkk06WE
2022-05-05 15:43:57 [INFO] [github.com/apache/shenyu-client-golang/example/http_client/main.go:55] finish register metadata ,the result is-> true
2022-05-05 15:43:57 [INFO] [github.com/apache/shenyu-client-golang/example/http_client/main.go:70] finish UrlRegister ,the result is-> true

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

**4. Get shenyu nacos client**
```go
    sdkClient := shenyu_sdk_client.GetFactoryClient(constants.NACOS_CLIENT)
    result, createResult, err := sdkClient.NewClient(ncp)
        if !createResult && err != nil {
        logger.Fatal("Create nacos client error : %+V", err)
    }

	nc := &nacos_client.ShenYuNacosClient{
		NacosClient: result.(*naming_client.NamingClient),
	}
```


**5.Use client to invoke RegisterServiceInstance**
```go
   instanceInfo, err := nc.GetServiceInstanceInfo(queryData)
    if instanceInfo == nil {
        logger.Fatal("Register nacos Instance error : %+V", err)
    }
        //do your logic
```

**6.Use client to invoke DeregisterServiceInstance**
```go
//DeregisterServiceInstance start
    deregisterInstanceParam := vo.DeregisterInstanceParam{
    Ip:          "10.0.0.10",
    Port:        8848,
    ServiceName: "demo.go",
    Ephemeral:   true,
    //Cluster:     "cluster-a", // default value is DEFAULT
    GroupName: "group-a", // default value is DEFAULT_GROUP
}

    serviceInstance, err := nc.DeregisterServiceInstance(deregisterInstanceParam)
        if !serviceInstance && err != nil {
        logger.Info("DeregisterServiceInstance result : %+V", serviceInstance)
}
        //do your logic
```

**7.Use client to invoke GetServiceInstanceInfo**
```go
        instanceInfo, result, err := nc.GetServiceInstanceInfo(queryData)
            if result != false && err != nil {
            logger.Fatal("Register nacos Instance error : %+V", err)
        }
        //do your logic
```

## Entire Success log
```go
2022-06-27 10:56:17 [INFO] [github.com/shenyu-client-golang/clients/nacos_client/nacos_client.go:92] RegisterServiceInstance,result:true

,param:{Ip:10.0.0.10 Port:8848 Weight:10 Enable:true Healthy:true Metadata:map[contextPath:contextPath uriMetadata:{"protocol":"testMetaDataRegister","appName":"testURLRegister","contextPath":"contextPath","rpcType":"http","host":"127.0.0.1","port":"8080"}] ClusterName: ServiceName:demo.go GroupName: Ephemeral:true}

```



---
## The Zookeeper type Register

**1.Fist make sure your Zookeeper env is correct,the set this necessary param.**
```go
    //Create ShenYuZkClient  start
    zcp := &zk_client.ZkClientParam{
    ZkServers: []string{"127.0.0.1:2181"}, //require user provide
    ZkRoot:    "/api",                     //require user provide
    }

    sdkClient := shenyu_sdk_client.GetFactoryClient(constants.ZOOKEEPER_CLIENT)
    client, createResult, err := sdkClient.NewClient(zcp)

    if !createResult && err != nil {
    logger.Fatal("Create ShenYuZkClient error : %+V", err)
    }

    zc := client.(*zk_client.ShenYuZkClient)
    defer zc.Close()
    //Create ShenYuZkClient end
```

**2. Prepare your service metaData to register**
```go
//metaData is necessary param, this will be register to shenyu gateway to use
    metaData1 := &model.MetaDataRegister{
        AppName: "testMetaDataRegister1", //require user provide
        Path:    "your/path1",            //require user provide
        Enabled: true,                    //require user provide
        Host:    "127.0.0.1",             //require user provide
        Port:    "8080",                  //require user provide
    }

    metaData2 := &model.MetaDataRegister{
        AppName: "testMetaDataRegister2", //require user provide
        Path:    "your/path2",            //require user provide
        Enabled: true,                    //require user provide
        Host:    "127.0.0.1",             //require user provide
        Port:    "8181",                  //require user provide
    }
```

**3.use client to invoke RegisterServiceInstance**
```go
   //register multiple metaData
    registerResult1, err := zc.RegisterServiceInstance(metaData1)
        if !registerResult1 && err != nil {
        logger.Fatal("Register zk Instance error : %+V", err)
    }

    registerResult2, err := zc.RegisterServiceInstance(metaData2)
        if !registerResult2 && err != nil {
        logger.Fatal("Register zk Instance error : %+V", err)
    }
    //do your logic
```

**4.use client to invoke DeregisterServiceInstance**
```go
    //your can chose to invoke,not require
    deRegisterResult1, err := zc.DeregisterServiceInstance(metaData1)
        if err != nil {
        panic(err)
        }

    deRegisterResult2, err := zc.DeregisterServiceInstance(metaData2)
        if err != nil {
        panic(err)
        }
```

**5.use client to GetServiceInstanceInfo**
```go
    //GetServiceInstanceInfo start
    instanceDetail, err := zc.GetServiceInstanceInfo(metaData1)
        nodes1, ok := instanceDetail.([]*model.MetaDataRegister)
        if !ok {
        logger.Fatal("get zk client metaData error %+v:", err)
     }
    
    //range nodes
    for index, node := range nodes1 {
        nodeJson, err := json.Marshal(node)
        if err == nil {
        logger.Info("GetNodesInfo ,success Index", index, string(nodeJson))
        }
    }
    
    instanceDetail2, err := zc.GetServiceInstanceInfo(metaData2)
        nodes2, ok := instanceDetail2.([]*model.MetaDataRegister)
        if !ok {
            logger.Fatal("get zk client metaData error %+v:", err)
    }
    //GetServiceInstanceInfo end

```

## Entire Success log
```go
2022-07-13 16:09:31 [INFO] [github.com/shenyu-client-golang/example/zk_client/main.go:105] GetNodesInfo ,success Index 0 {"appName":"testMetaDataRegister1","path":"your/path1","contextPath":"","ruleName":"","rpcType":"","enabled":true,"host":"127.0.0.1","port":"8080","pluginNames":null,"registerMetaData":false,"timeMillis":0}
2022-07-13 16:09:31 [INFO] [github.com/shenyu-client-golang/example/zk_client/main.go:119] GetNodesInfo ,success Index 0 {"appName":"testMetaDataRegister2","path":"your/path2","contextPath":"","ruleName":"","rpcType":"","enabled":true,"host":"127.0.0.1","port":"8181","pluginNames":null,"registerMetaData":false,"timeMillis":0}
2022-07-13 16:09:31 [INFO] [github.com/shenyu-client-golang/example/zk_client/main.go:132] GetNodesInfo ,success Index 0 {"appName":"testMetaDataRegister3","path":"your/path3","contextPath":"","ruleName":"","rpcType":"","enabled":true,"host":"127.0.0.1","port":"8282","pluginNames":null,"registerMetaData":false,"timeMillis":0}
2022-07-13 16:09:31 [INFO] [github.com/shenyu-client-golang/example/zk_client/main.go:139] > DeregisterServiceInstance start
2022-07-13 16:09:31 [INFO] [github.com/shenyu-client-golang/clients/zk_client/zk_client.go:213] ensureName check, path is -> /api/testMetaDataRegister1
2022-07-13 16:09:31 [INFO] [github.com/shenyu-client-golang/clients/zk_client/zk_client.go:215] ensureName check result is -> true
2022-07-13 16:09:31 [INFO] [github.com/shenyu-client-golang/clients/zk_client/zk_client.go:213] ensureName check, path is -> /api/testMetaDataRegister2
2022-07-13 16:09:31 [INFO] [github.com/shenyu-client-golang/clients/zk_client/zk_client.go:215] ensureName check result is -> true
2022-07-13 16:09:31 [INFO] [github.com/shenyu-client-golang/clients/zk_client/zk_client.go:213] ensureName check, path is -> /api/testMetaDataRegister3
2022-07-13 16:09:31 [INFO] [github.com/shenyu-client-golang/clients/zk_client/zk_client.go:215] ensureName check result is -> true
2022-07-13 16:09:31 [INFO] [github.com/shenyu-client-golang/example/zk_client/main.go:156] DeregisterServiceInstance success !
```
