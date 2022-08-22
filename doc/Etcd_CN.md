
---
## 以Etcd方式注册到ShenYu网关

**1.首先确保你的Etcd环境是正确，然后设置这些Etcd必要的参数 .**
```go
    //开始创建ShenYuEtcdClient 
    ecp := &etcd_client.EtcdClientParam{
    ServerList: []string{"http://127.0.0.1:2379"}, //需要用户提供
    UserName : "" // 可选参数etcd访问账号名
    Password : "" // 可选参数etcd访问密码
    TTL:    50, // 需要用户提供 key生存时间
    }

    sdkClient := shenyu_sdk_client.GetFactoryClient(constants.ETCD_CLIENT)
    client, createResult, err := sdkClient.NewClient(ecp)
    if !createResult && err != nil {
    logger.Fatal("Create ShenYuEtcdClient error : %+V", err)
    }
    
    etcd := client.(*etcd_client.ShenYuEtcdClient)
    defer etcd.Close()
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
    result, err := etcd.PersistInterface(metaData)
    if err != nil {
    logger.Warn("MetaDataRegister has error:", err)
    }
    logger.Info("finish register metadata ,the result is->", result)


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
    result, err = etcd.PersistInterface(urlRegister)
    if err != nil {
    logger.Warn("UrlRegister has error:", err)
    }
    logger.Info("finish UrlRegister ,the result is->", result)

```

## 完整的成功日志
```go
2022-08-19 21:55:15 [INFO] [github.com/shenyu-client-golang/clients/etcd_client/etcd_client.go:103] http etcd client register success: {"appName":"testGoAppName2","path":"/golang/your/path","pathDesc":"","contextPath":"/golang","ruleName":"","rpcType":"http","serviceName":"","methodName":"","parameterTypes":"","rpcExt":"","enabled":true,"host":"127.0.0.1","port":"8080","pluginNames":null,"registerMetaData":false,"timeMillis":0}
2022-08-19 21:55:15 [INFO] [github.com/shenyu-client-golang/example/etcd_client/main.go:62] finish register metadata ,the result is-> true
2022-08-19 21:55:25 [INFO] [github.com/shenyu-client-golang/example/etcd_client/main.go:78] finish UrlRegister ,the result is-> true
```
