
---
## The Etcd type Register

**1.Fist make sure your Etcd env is correct,the set this necessary param.**
```go
    //Create ShenYuEtcdClient  start
    ecp := &etcd_client.EtcdClientParam{
    ServerList: []string{"http://127.0.0.1:2379"}, // require user provider
    UserName : "" // optional param etcd userName
    Password : "" // optional param etcd pwd
    TTL:    50, // require user provider param key live
    }
    
    sdkClient := shenyu_sdk_client.GetFactoryClient(constants.ETCD_CLIENT)
    client, createResult, err := sdkClient.NewClient(ecp)
    if !createResult && err != nil {
    logger.Fatalf("Create ShenYuEtcdClient error : %+V", err)
    }
    
    etcd := client.(*etcd_client.ShenYuEtcdClient)
    defer etcd.Close()  
    //Create ShenYuEtcdClient end
```


**2.Step 1 Register MetaData to ShenYu GateWay. **
```go
//MetaDataRegister
metaData := &model.MetaDataRegister{
		AppName: "testMetaDataRegister", //require user provide
		Path:    "/your/path",           //require user provide
		Enabled: true,                   //require user provide
		Host:    "127.0.0.1",            //require user provide
		Port:    "8080",                 //require user provide
	}
    result, err := etcd.PersistInterface(metaData)
    if err != nil {
    logger.Warnf("MetaDataRegister has error:", err)
    }
    logger.Infof("finish register metadata ,the result is->", result)
	
When Register success , you will see this :  
finish register metadata ,the result is-> true
```

**3.Step 2  Url  Register  to ShenYu GateWay. (Need step 1 token to invoke)**
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
    result, err = etcd.PersistInterface(urlRegister)
    if err != nil {
    logger.Warnf("UrlRegister has error:", err)
    }
    logger.Infof("finish UrlRegister ,the result is->", result)


```

## Entire Success log
```go
2022-08-19 21:55:15 [INFO] [github.com/shenyu-client-golang/clients/etcd_client/etcd_client.go:103] http etcd client register success: {"appName":"testGoAppName2","path":"/golang/your/path","pathDesc":"","contextPath":"/golang","ruleName":"","rpcType":"http","serviceName":"","methodName":"","parameterTypes":"","rpcExt":"","enabled":true,"host":"127.0.0.1","port":"8080","pluginNames":null,"registerMetaData":false,"timeMillis":0}
2022-08-19 21:55:15 [INFO] [github.com/shenyu-client-golang/example/etcd_client/main.go:62] finish register metadata ,the result is-> true
2022-08-19 21:55:25 [INFO] [github.com/shenyu-client-golang/example/etcd_client/main.go:78] finish UrlRegister ,the result is-> true
```
