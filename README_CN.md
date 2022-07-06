# shenyu-client-golang

English | [简体中文](README_CN.md)

[![Build and Test](https://github.com/apache/incubator-shenyu-client-golang/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/apache/incubator-shenyu-client-golang/actions)
[![codecov.io](https://codecov.io/gh/apache/incubator-shenyu-client-golang/coverage.svg?branch=main)](https://app.codecov.io/gh/apache/incubator-shenyu-client-golang?branch=main)
[![GoDoc](https://godoc.org/github.com/apache/incubator-shenyu-client-golang?status.svg)](https://godoc.org/github.com/apache/incubator-shenyu-client-golang)

---

## Shenyu-client-golang
Shenyu-client-golang是提供了Go语言访问ShenYu网关的功能，并支持服务注册到ShenYu网关。

---
## 已支持注册到ShenYu网关的方式
* **以Http方式注册**
* **以Nacos方式注册**
* **以Zookeeper方式注册**

---

## 要求

要求Go语言版本 **1.12**

SDK支持ShenYu的版本 **2.4.3及以上**

## 安装方法

使用 `go get命令` 安装 SDK：

```sh
$ go get -u github.com/apache/incubator-shenyu-client-golang
```

## 代码列子路径

* incubator-shenyu-client-golang/example/**_client/main.go
---

##  以Http方式注册到ShenYu网关

**1.首先确保ShenYuAdmin是启动的，并且ShenYuAdmin服务启动的端口是9095 .**
```go
如果没启动,你将看到如下错误:
	
2022-05-05 15:24:28 [WARN] [github.com/apache/incubator-shenyu-client-golang/example/http_client/main.go:53] MetaDataRegister has error: The errCode is ->:503, The errMsg is  ->:Please check ShenYu admin service status

caused by:
Post "http://127.0.0.1:9095/shenyu-client/register-metadata": dial tcp 127.0.0.1:9095: connect: connection refused
2022-05-05 15:24:28 [INFO] [github.com/apache/incubator-shenyu-client-golang/example/http_client/main.go:55] finish register metadata ,the result is-> false
2022-05-05 15:24:28 [WARN] [github.com/apache/incubator-shenyu-client-golang/example/http_client/main.go:68] UrlRegister has error: The errCode is ->:503, The errMsg is  ->:Please check ShenYu admin service status

caused by:
Post "http://127.0.0.1:9095/shenyu-client/register-uri": dial tcp 127.0.0.1:9095: connect: connection refused
2022-05-05 15:24:28 [INFO] [github.com/apache/incubator-shenyu-client-golang/example/http_client/main.go:70] finish UrlRegister ,the result is-> false
	
```

**2.获取shenyu_admin_client. (注册服务需要这个实例)**

```go
//初始化 ShenYuAdminClient
adminClient := &model.ShenYuAdminClient{
    UserName: "admin",  //需要用户提供
    Password: "123456", //需要用户提供
}

adminToken, err := clients.NewShenYuAdminClient(adminClient)

adminToken像这样 :
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

当你成功获取到Token,你将看到这些:
this is ShenYu Admin client token -> eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyTmFtZSI6ImFkbWluIiwiZXhwIjoxNjUwNjc5OTQ2fQ.K92Il2kmJ0X3FgjY4igW35-pw9nsf5VKdUyqBoyIaF4

```


**3.注册元数据到ShenYu网关. (需要上一步的adminToken去调用)**
```go
//元数据注册(需要上一步的token: adminToken.AdminTokenData)
metaData := &model.MetaDataRegister{
		AppName: "testMetaDataRegister", //需要用户提供
		Path:    "/your/path",           //需要用户提供
		Enabled: true,                   //需要用户提供
		Host:    "127.0.0.1",            //需要用户提供
		Port:    "8080",                 //需要用户提供
	}
	result, err := clients.RegisterMetaData(adminToken.AdminTokenData, metaData)
	if err != nil {
		logger.Warn("MetaDataRegister has error:",err)
	}
	logger.Info("finish register metadata ,the result is->", result)


当你注册成功,你将看到这些:
finish register metadata ,the result is-> true
```

**4.以URL的方式注册到ShenYu网关. (需要上一步的adminToken去调用)**
```go
//URI注册(需要上一步的token: adminToken.AdminTokenData)
//初始化 URI注册
	urlRegister := &model.URIRegister{
		Protocol:    "testMetaDataRegister", //需要用户提供
		AppName:     "testURLRegister",      //需要用户提供
		ContextPath: "contextPath",          //需要用户提供
		RPCType:     constants.RPCTYPE_HTTP, //需要用户提供
		Host:        "127.0.0.1",            //需要用户提供
		Port:        "8080",                 //需要用户提供
	}
	result, err = clients.UrlRegister(adminToken.AdminTokenData, urlRegister)
	if err != nil {
		logger.Warn("UrlRegister has error:", err)
	}
	logger.Info("finish UrlRegister ,the result is->", result)
         //做你的逻辑处理
```

**5.完整的成功日志**
```go
2022-05-05 15:43:56 [INFO] [github.com/apache/incubator-shenyu-client-golang/clients/admin_client/shenyu_admin_client.go:51] Get ShenYu Admin response, body is -> {200 login dashboard user success {1 admin 1 true 2018-06-23 15:12:22 2022-03-09 15:08:14 eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyTmFtZSI6ImFkbWluIiwiZXhwIjoxNjUxODIzMDM2fQ.gMzPKaNlXEd1Q517qQamOpg358W9L0-0cZN3lkk06WE}}
2022-05-05 15:43:56 [INFO] [github.com/apache/incubator-shenyu-client-golang/example/http_client/main.go:40] this is ShenYu Admin client token -> eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyTmFtZSI6ImFkbWluIiwiZXhwIjoxNjUxODIzMDM2fQ.gMzPKaNlXEd1Q517qQamOpg358W9L0-0cZN3lkk06WE
2022-05-05 15:43:57 [INFO] [github.com/apache/incubator-shenyu-client-golang/example/http_client/main.go:55] finish register metadata ,the result is-> true
2022-05-05 15:43:57 [INFO] [github.com/apache/incubator-shenyu-client-golang/example/http_client/main.go:70] finish UrlRegister ,the result is-> true

```



---
## 以Nacos方式注册到ShenYu网关

**1.首先确保你的nacos环境是正确，然后设置这些nacos必要的参数 .**
```go
//设置nacos环境配置
    ncp := &nacos_client.NacosClientParam{
        IpAddr:      "console.nacos.io",
        Port:        80,
        NamespaceId: "e525eafa-f7d7-4029-83d9-008937f9d468",
}
```

**2. 准备你要注册服务的元数据信息**
```go
//元数据是必要的参数，这将注册到shenyu网关使用
metaData := &model.URIRegister{
        Protocol:    "testMetaDataRegister", //需要用户提供
        AppName:     "testURLRegister",      //需要用户提供
        ContextPath: "contextPath",          //需要用户提供
        RPCType:     constants.RPCTYPE_HTTP, //需要用户提供
        Host:        "127.0.0.1",            //需要用户提供
        Port:        "8080",                 //需要用户提供
}
    metaDataStringJson, _ := json.Marshal(metaData)
```

**3.准备你要注册服务的实例消息(包括元数据)**
```go
//初始化Nacos注册实例信息
    nacosRegisterInstance := vo.RegisterInstanceParam{
        Ip:          "10.0.0.10", //需要用户提供
        Port:        8848,        //需要用户提供
        ServiceName: "demo.go",   //需要用户提供
        Weight:      10,          //需要用户提供
        Enable:      true,        //需要用户提供
        Healthy:     true,        //需要用户提供
        Ephemeral:   true,        //需要用户提供
        Metadata:    map[string]string{"contextPath": "contextPath", "uriMetadata": string(metaDataStringJson)},
}
```

**4. 获取ShenYu nacos的客户端**
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


**5.使用客户端调用RegisterNacosInstance方法**
```go
    registerResult, err := nc.RegisterServiceInstance(nacosRegisterInstance)
        if !registerResult && err != nil {
    logger.Fatal("Register nacos Instance error : %+V", err)
}
        //do your logic
```

**6.使用客户端调用DeregisterServiceInstance方法**
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

**7.使用客户端调用GetServiceInstanceInfo方法**
```go
<<<<<<< main
         instanceInfo, err := nc.GetServiceInstanceInfo(queryData)
            if instanceInfo == nil {
            	logger.Fatal("Register nacos Instance error : %+V", err)
            }
=======
        instanceInfo, result, err := nc.GetServiceInstanceInfo(queryData)
            if result != false && err != nil {
            logger.Fatal("Register nacos Instance error : %+V", err)
        }
>>>>>>> main
        //do your logic
```

**完整的成功日志**
```go
2022-06-27 10:56:17 [INFO] [github.com/incubator-shenyu-client-golang/clients/nacos_client/nacos_client.go:92] RegisterServiceInstance,result:true

,param:{Ip:10.0.0.10 Port:8848 Weight:10 Enable:true Healthy:true Metadata:map[contextPath:contextPath uriMetadata:{"protocol":"testMetaDataRegister","appName":"testURLRegister","contextPath":"contextPath","rpcType":"http","host":"127.0.0.1","port":"8080"}] ClusterName: ServiceName:demo.go GroupName: Ephemeral:true}

```




---
## The Zookeeper type Register

**1.首先确保你的Zookeeper环境是正确，然后设置这些Zookeeper必要的参数 .**
```go
    servers := []string{"127.0.0.1:2181"}             //需要用户提供
        client, err := NewClient(servers, "/api", 10) //需要用户提供
        if err != nil {
            panic(err)
         }
        defer client.Close()
```

**2.  准备你要注册服务的元数据信息**
```go
//元数据是必要的参数，这将注册到shenyu网关使用
        metaData1 := &model.MetaDataRegister{
            AppName: "testMetaDataRegister", //需要用户提供
            Path:    "your/path1",           //需要用户提供
            Enabled: true,                   //需要用户提供
            Host:    "127.0.0.1",            //需要用户提供
            Port:    "8080",                 //需要用户提供
        }
```

**3.使用客户端进行节点信息注册**
```go
   //可以进行多个实例注册
    if err := client.RegisterNodeInstance(metaData1); err != nil {
        panic(err)
    }
        //做你的逻辑处理
```

**4.使用客户端进行注册节点信息删除**
```go
    //选择性调用
    err = client.DeleteNodeInstance(metaData1)
     if err != nil {
       panic(err)
}
```

**5.使用客户端获取注册节点的信息**
```go
    //遍历节点
    for index, node := range nodes {
        nodeJson, err := json.Marshal(node)
        if err == nil {
        logger.Info("GetNodesInfo ,success Index", index, string(nodeJson))
    }
}
```

## 完整的成功日志
```go
2022-06-28 15:21:57 [INFO] [github.com/incubator-shenyu-client-golang/example/zk_client/zk_client.go:80] GetNodesInfo ,success Index 0 {"appName":"testMetaDataRegister","path":"your/path1","rpcType":"","enabled":true,"host":"127.0.0.1","port":"8080","pluginNames":null,"registerMetaData":false,"timeMillis":0}
2022-06-28 15:21:57 [INFO] [github.com/incubator-shenyu-client-golang/example/zk_client/zk_client.go:80] GetNodesInfo ,success Index 1 {"appName":"testMetaDataRegister","path":"your/path3","rpcType":"","enabled":true,"host":"127.0.0.1","port":"8282","pluginNames":null,"registerMetaData":false,"timeMillis":0}
2022-06-28 15:21:57 [INFO] [github.com/incubator-shenyu-client-golang/example/zk_client/zk_client.go:80] GetNodesInfo ,success Index 2 {"appName":"testMetaDataRegister","path":"your/path2","rpcType":"","enabled":true,"host":"127.0.0.1","port":"8181","pluginNames":null,"registerMetaData":false,"timeMillis":0}

```
