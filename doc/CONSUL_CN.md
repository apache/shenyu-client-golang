
---
## 以Consul方式注册到ShenYu网关

**1.首先确保你的Consul环境是正确，然后设置这些Consul必要的参数.**
```go
    //Create ShenYuConsulClient  start
    ccp := &consul_client.ConsulClientParam{
        ServerList:  []string{"127.0.0.1:8500"},
        Id: "testName",
        Tags: []string{"test1"},
        Token:"",
    }
    
    sdkClient := shenyu_sdk_client.GetFactoryClient(constants.CONSUL_CLIENT)
    client, createResult, err := sdkClient.NewClient(ccp)
    
    if !createResult && err != nil {
    logger.Fatalf("Create ShenYuConsulClient error : %+V", err)
    }
    
    scc := client.(*consul_client.ShenYuConsulClient)
    //Create ShenYuConsulClient end
```

**2. 准备你要注册服务的元数据信息**
```go
//元数据注册
     metaData := &model.MetaDataRegister{
		AppName: "testMetaDataRegister", //需要用户提供
		Path:    "/your/path",           //需要用户提供
		Enabled: true,                   //需要用户提供
		Host:    "127.0.0.1",            //需要用户提供
		Port:    "8080",                 //需要用户提供
	}
    result, err := acc.PersistInterface(metaData)
    if err != nil {
    logger.Warnf("MetaDataRegister has error:", err)
    }
    logger.Infof("finish register metadata ,the result is->", result)


当你注册成功,你将看到这些:
finish register metadata ,the result is-> true
```

**3.以URL的方式注册到ShenYu网关. **
```go
    //URI注册
    //初始化 URI注册
    urlRegister := &model.URIRegister{
    Protocol:    "testMetaDataRegister", //需要用户提供
    AppName:     "testURLRegister",      //需要用户提供
    ContextPath: "contextPath",          //需要用户提供
    RPCType:     constants.RPCTYPE_HTTP, //需要用户提供
    Host:        "127.0.0.1",            //需要用户提供
    Port:        "8080",                 //需要用户提供
    }
    result, err = acc.PersistInterface(urlRegister)
    if err != nil {
    logger.Warnf("UrlRegister has error:", err)
    }
    logger.Infof("finish UrlRegister ,the result is->", result)
        

```

## 完整的成功日志
```go

2022-08-19 21:55:15 [INFO] [github.com/shenyu-client-golang/clients/consul_client/consul_client.go:103] http consul client register success: {"appName":"testGoAppName2","path":"/golang/your/path","pathDesc":"","contextPath":"/golang","ruleName":"","rpcType":"http","serviceName":"","methodName":"","parameterTypes":"","rpcExt":"","enabled":true,"host":"127.0.0.1","port":"8080","pluginNames":null,"registerMetaData":false,"timeMillis":0}
2022-08-19 21:55:15 [INFO] [github.com/shenyu-client-golang/example/consul_client/main.go:62] finish register metadata ,the result is-> true
2022-08-19 21:55:25 [INFO] [github.com/shenyu-client-golang/example/consul_client/main.go:78] finish UrlRegister ,the result is-> true

```

