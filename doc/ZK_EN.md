
---
## The Zookeeper type Register

**1.Fist make sure your Zookeeper env is correct,the set this necessary param.**
```go
    //Create ShenYuZkClient  start
    zcp := &zk_client.ZkClientParam{
      ServerList: []string{"127.0.0.1:2181"}, //require user provide
      Digest: "",
    }

    sdkClient := shenyu_sdk_client.GetFactoryClient(constants.ZOOKEEPER_CLIENT)
    client, createResult, err := sdkClient.NewClient(zcp)

    if !createResult && err != nil {
    logger.Fatal("Create ShenYuZkClient error : %+V", err)
    }

    zc := client.(*zk_client.ShenYuZkClient)
    go func() {
    zc.WatchEventHandler()
    }()
    defer zc.Close()
    //Create ShenYuZkClient end
```


**2.Step 1 Register MetaData to ShenYu GateWay. (Need step 1 token to invoke)**
```go
//MetaDataRegister(Need Step 1 toekn adminToken.AdminTokenData)
metaData := &model.MetaDataRegister{
		AppName: "testMetaDataRegister", //require user provide
		Path:    "/your/path",           //require user provide
		Enabled: true,                   //require user provide
		Host:    "127.0.0.1",            //require user provide
		Port:    "8080",                 //require user provide
	}
    result, err := zc.PersistInterface(metaData)
    if err != nil {
    logger.Warn("MetaDataRegister has error:", err)
    }
    logger.Info("finish register metadata ,the result is->", result)
	
When Register success , you will see this :  
finish register metadata ,the result is-> true
```

**3.Step 2  Url  Register  to ShenYu GateWay.**
```go
    //URIRegister
    //init urlRegister
    urlRegister := &model.URIRegister{
    Protocol:    "testMetaDataRegister", //require user provide
    AppName:     "testURLRegister",      //require user provide
    ContextPath: "contextPath",          //require user provide
    RPCType:     constants.RPCTYPE_HTTP, //require user provide
    Host:        "127.0.0.1",            //require user provide
    Port:        "8080",                 //require user provide
    }
    result, err = zc.PersistInterface(urlRegister)
    if err != nil {
    logger.Warn("UrlRegister has error:", err)
    }
    logger.Info("finish UrlRegister ,the result is->", result)

```

## Entire Success log
```go
2022-08-19 21:55:15 [INFO] [github.com/shenyu-client-golang/clients/zk_client/zk_client.go:103] http zookeeper client register success: {"appName":"testGoAppName2","path":"/golang/your/path","pathDesc":"","contextPath":"/golang","ruleName":"","rpcType":"http","serviceName":"","methodName":"","parameterTypes":"","rpcExt":"","enabled":true,"host":"127.0.0.1","port":"8080","pluginNames":null,"registerMetaData":false,"timeMillis":0}
2022-08-19 21:55:15 [INFO] [github.com/shenyu-client-golang/example/zk_client/main.go:62] finish register metadata ,the result is-> true
2022-08-19 21:55:25 [INFO] [github.com/shenyu-client-golang/example/zk_client/main.go:78] finish UrlRegister ,the result is-> true
```
