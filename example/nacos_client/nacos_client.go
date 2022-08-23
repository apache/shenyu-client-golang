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
	"github.com/apache/shenyu-client-golang/clients/nacos_client"
	"github.com/apache/shenyu-client-golang/common/constants"
	"github.com/apache/shenyu-client-golang/common/shenyu_sdk_client"
	"github.com/apache/shenyu-client-golang/model"
	"github.com/sirupsen/logrus"
)

var (
	logger = logrus.New()
)

/**
* The nacos_client example
**/
func main() {

	//Create ShenYuNacosClient start
	//set nacos env configuration
	ncp := &nacos_client.NacosClientParam{
		ServerList:   []string{ "http://console.nacos.io:80"},   //"console.nacos.io",
		NamespaceId: "ShenyuRegisterCenter",
		UserName: "",
		Password: "",
	}

	sdkClient := shenyu_sdk_client.GetFactoryClient(constants.NACOS_CLIENT)
	client, createResult, err := sdkClient.NewClient(ncp)
	if !createResult && err != nil {
		logger.Fatalf("Create ShenYuNacosClient error : %+v", err)
	}

	nc := client.(*nacos_client.ShenYuNacosClient)
	//Create ShenYuNacosClient end

	//RegisterServiceInstance start
	//metaData is necessary param, this will be register to shenyu gateway to use
	metaData := &model.MetaDataRegister{
		AppName:     "testGoAppName2",     //require user provide
		Path:        "/your/path", //require user provide
		ContextPath: "/golang",           //require user provide
		RPCType: constants.RPCTYPE_HTTP,
		Enabled:     true,                //require user provide
		Host:        "127.0.0.1",         //require user provide
		Port:        "8080",              //require user provide
	}
	result, err := nc.PersistInterface(metaData)
	if err != nil {
		logger.Warnf("MetaDataRegister has error:%+v", err)
	}
	logger.Infof("finish register metadata ,the result is->%+v", result)

	//init urlRegister
	urlRegister := &model.URIRegister{
		Protocol:    "http://",              //require user provide
		AppName:     "testGoAppName2",        //require user provide
		ContextPath: "/golang",              //require user provide
		RPCType:     constants.RPCTYPE_HTTP, //require user provide
		Host:        "127.0.0.1",            //require user provide
		Port:        "8080",                 //require user provide
	}
	result, err = nc.PersistURI(urlRegister)
	if err != nil {
		logger.Warnf("UrlRegister has error:%+v", err)
	}
	logger.Infof("finish UrlRegister ,the result is->%+v", result)
	//do your logic
}
