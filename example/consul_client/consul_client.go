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
	"github.com/apache/shenyu-client-golang/clients/consul_client"
	"github.com/apache/shenyu-client-golang/common/constants"
	"github.com/apache/shenyu-client-golang/common/shenyu_sdk_client"
	"github.com/apache/shenyu-client-golang/model"
	"github.com/wonderivan/logger"
)

func main() {
	var serverList = []string{
		"127.0.0.1:8500",
	}
	//Create ShenYuConsulClient  start
	ccp := &consul_client.ConsulClientParam{
		ServerList:  serverList,
		Id: "testName",
		NameSpace: "testName",
		Tags: []string{"test1"},
	}

	sdkClient := shenyu_sdk_client.GetFactoryClient(constants.CONSUL_CLIENT)
	client, createResult, err := sdkClient.NewClient(ccp)

	if !createResult && err != nil {
		logger.Fatal("Create ShenYuConsulClient error : %+V", err)
	}

	scc := client.(*consul_client.ShenYuConsulClient)

	//init MetaDataRegister
	metaData := &model.MetaDataRegister{
		AppName:     "testGoAppName2",     //require user provide
		Path:        "/your/path", //require user provide
		ContextPath: "/golang",           //require user provide
		RPCType: constants.RPCTYPE_HTTP,
		Enabled:     true,                //require user provide
		Host:        "127.0.0.1",         //require user provide
		Port:        "8080",              //require user provide
	}
	result, err := scc.PersistInterface(metaData)
	if err != nil {
		logger.Warn("MetaDataRegister has error:", err)
	}
	logger.Info("finish register metadata ,the result is->", result)

	//init urlRegister
	urlRegister := &model.URIRegister{
		Protocol:    "http://",              //require user provide
		AppName:     "testGoAppName2",        //require user provide
		ContextPath: "/golang",              //require user provide
		RPCType:     constants.RPCTYPE_HTTP, //require user provide
		Host:        "127.0.0.1",            //require user provide
		Port:        "8080",                 //require user provide
	}
	result, err = scc.PersistURI(urlRegister)
	if err != nil {
		logger.Warn("UrlRegister has error:", err)
	}
	logger.Info("finish UrlRegister ,the result is->", result)


	//do your logic

}
