
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
    var serverList = []string{
    "http://127.0.0.1:9095",
    }
    acp := &admin_client.ShenYuAdminClientParams{
    ServerList: serverList,
    UserName: "admin",  //require user provide
    Password: "123456", //require user provide
}


    sdkClient := shenyu_sdk_client.GetFactoryClient(constants.RPCTYPE_HTTP)
    client, createResult, err := sdkClient.NewClient(acp)
    acc := client.(*admin_client.ShenYuAdminClient)
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
    result, err := acc.PersistInterface(metaData)
    if err != nil {
    logger.Warn("MetaDataRegister has error:", err)
    }
    logger.Info("finish register metadata ,the result is->", result)
	
When Register success , you will see this :  
finish register metadata ,the result is-> true
```

**4.Step 3  Url  Register  to ShenYu GateWay. **
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
    result, err = acc.PersistInterface(urlRegister)
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
