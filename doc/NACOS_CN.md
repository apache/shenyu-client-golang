
---
## 以Nacos方式注册到ShenYu网关

**1.首先确保你的nacos环境是正确，然后设置这些nacos必要的参数 .**
```go
//设置nacos环境配置
    ncp := &nacos_client.NacosClientParam{
        ServerList:   []string{ "http://127.0.0.1:8848"},   //"console.nacos.io",
        NamespaceId: "ShenyuRegisterCenter",
        UserName: "nacos",
        Password: "nacos",
}
    sdkClient := shenyu_sdk_client.GetFactoryClient(constants.NACOS_CLIENT)
    client, createResult, err := sdkClient.NewClient(ncp)
    if !createResult && err != nil {
    logger.Fatal("Create ShenYuNacosClient error : %+V", err)
    }
    
    nc := client.(*nacos_client.ShenYuNacosClient)
    //Create ShenYuNacosClient end
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
    result, err := nc.PersistInterface(metaData)
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
    result, err = nc.PersistInterface(urlRegister)
    if err != nil {
    logger.Warn("UrlRegister has error:", err)
    }
    logger.Info("finish UrlRegister ,the result is->", result)
```

**完整的成功日志**
```go

2022-08-19T22:06:49.169+0800	INFO	nacos_client/nacos_client.go:79	logDir:</tmp/nacos/log>   cacheDir:</tmp/nacos/cache>
2022-08-19T22:06:49.348+0800	INFO	naming_client/push_receiver.go:80	udp server start, port: 55797
2022-08-19T22:06:49.349+0800	INFO	nacos_client/nacos_client.go:79	logDir:</tmp/nacos/log>   cacheDir:</tmp/nacos/cache>
2022-08-19 22:06:49 [INFO] [github.com/shenyu-client-golang/clients/nacos_client/nacos_client.go:159] rpcType:http ->Consul client register success,meta:{"appName":"testGoAppName2","path":"/golang/your/path","pathDesc":"","contextPath":"/golang","ruleName":"","rpcType":"http","serviceName":"","methodName":"","parameterTypes":"","rpcExt":"","enabled":true,"host":"127.0.0.1","port":"8080","pluginNames":null,"registerMetaData":false,"timeMillis":0}->ruleName:
2022-08-19 22:06:49 [INFO] [github.com/shenyu-client-golang/example/nacos_client/main.go:66] finish register metadata ,the result is-> true
2022-08-19T22:06:49.523+0800	INFO	naming_client/naming_proxy.go:54	register instance namespaceId:<ShenyuRegisterCenter>,serviceName:<DEFAULT_GROUP@@shenyu.register.service.http> with instance:<{"valid":false,"marked":false,"instanceId":"","port":8080,"ip":"127.0.0.1","weight":0,"metadata":{"contextPath":"/golang/testGoAppName2","uriMetadata":"{\"protocol\":\"http://\",\"appName\":\"testGoAppName2\",\"contextPath\":\"/golang\",\"rpcType\":\"http\",\"host\":\"127.0.0.1\",\"port\":\"8080\"}"},"clusterName":"","serviceName":"","enabled":false,"healthy":false,"ephemeral":true}>
2022-08-19T22:06:49.527+0800	INFO	naming_client/beat_reactor.go:68	adding beat: <{"ip":"127.0.0.1","port":8080,"weight":0,"serviceName":"DEFAULT_GROUP@@shenyu.register.service.http","cluster":"","metadata":{"contextPath":"/golang/testGoAppName2","uriMetadata":"{\"protocol\":\"http://\",\"appName\":\"testGoAppName2\",\"contextPath\":\"/golang\",\"rpcType\":\"http\",\"host\":\"127.0.0.1\",\"port\":\"8080\"}"},"scheduled":false}> to beat map
2022-08-19 22:06:49 [INFO] [github.com/shenyu-client-golang/clients/nacos_client/nacos_client.go:190] RegisterServiceInstance,result:true

,param:{Ip:127.0.0.1 Port:8080 Weight:0 Enable:false Healthy:false Metadata:map[contextPath:/golang/testGoAppName2 uriMetadata:{"protocol":"http://","appName":"testGoAppName2","contextPath":"/golang","rpcType":"http","host":"127.0.0.1","port":"8080"}] ClusterName: ServiceName:shenyu.register.service.http GroupName:DEFAULT_GROUP Ephemeral:true}


2022-08-19 22:06:49 [INFO] [github.com/shenyu-client-golang/example/nacos_client/main.go:81] finish UrlRegister ,the result is-> true

```
