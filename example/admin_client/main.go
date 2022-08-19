/*
* Licensed to the Apache Software Foundation (ASF) under one or more
* contributor license agreements.  See the NOTICE file distributed with
* this work for additional information regarding copyright ownership.
* The ASF licenses this file to You under the Apache License, Version 2.0
* (the "License"); you may not use this file except in compliance with
* the License.  You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

package main

import (
	"github.com/apache/shenyu-client-golang/clients/admin_client"
	"github.com/apache/shenyu-client-golang/common/constants"
	"github.com/apache/shenyu-client-golang/common/shenyu_sdk_client"
	"github.com/apache/shenyu-client-golang/model"
	"github.com/wonderivan/logger"
)

/**
* The shenyu_admin_client example
**/
func main() {

	var serverList = []string{
		"http://127.0.0.1:9095",
	}
	//Create ShenYuAdminClient  start
	ccp := &admin_client.ShenYuAdminClientParams{
		ServerList: serverList,
		UserName:  "admin",
		Password: "123456",
	}

	sdkClient := shenyu_sdk_client.GetFactoryClient(constants.RPCTYPE_HTTP)
	client, createResult, err := sdkClient.NewClient(ccp)

	if !createResult && err != nil {
		logger.Fatal("Create ShenYuAdminClient error : %+V", err)
	}

	acc := client.(*admin_client.ShenYuAdminClient)

	//init MetaDataRegister
	metaData := &model.MetaDataRegister{
		AppName:     "testGoAppName",     //require user provide
		Path:        "/golang/your/path", //require user provide
		ContextPath: "/golang",           //require user provide
		Enabled:     true,                //require user provide
		Host:        "127.0.0.1",         //require user provide
		Port:        "8080",              //require user provide
	}
	result, err := acc.PersistInterface(metaData)
	if err != nil {
		logger.Warn("MetaDataRegister has error:", err)
	}
	logger.Info("finish register metadata ,the result is->", result)

	//init urlRegister
	urlRegister := &model.URIRegister{
		Protocol:    "http://",              //require user provide
		AppName:     "testGoAppName",        //require user provide
		ContextPath: "/golang",              //require user provide
		RPCType:     constants.RPCTYPE_HTTP, //require user provide
		Host:        "127.0.0.1",            //require user provide
		Port:        "8080",                 //require user provide
	}
	result, err = acc.PersistInterface(urlRegister)
	if err != nil {
		logger.Warn("UrlRegister has error:", err)
	}
	logger.Info("finish UrlRegister ,the result is->", result)

	//do you logic

}
