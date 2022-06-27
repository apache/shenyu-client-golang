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
	"encoding/json"
	"github.com/apache/incubator-shenyu-client-golang/clients/nacos_client"
	"github.com/apache/incubator-shenyu-client-golang/common/constants"
	"github.com/apache/incubator-shenyu-client-golang/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/wonderivan/logger"
)

/**
 * The nacos_client example
 **/
func main() {

	//set nacos env configuration
	ncp := &nacos_client.NacosClientParam{
		IpAddr:      "console.nacos.io",
		Port:        80,
		NamespaceId: "e525eafa-f7d7-4029-83d9-008937f9d468",
	}

	//metaData is necessary param, this will be register to shenyu gateway to use
	metaData := &model.URIRegister{
		Protocol:    "testMetaDataRegister", //require user provide
		AppName:     "testURLRegister",      //require user provide
		ContextPath: "contextPath",          //require user provide
		RPCType:     constants.RPCTYPE_HTTP, //require user provide
		Host:        "127.0.0.1",            //require user provide
		Port:        "8080",                 //require user provide
	}
	metaDataStringJson, _ := json.Marshal(metaData)

	//init NacosRegisterInstance
	nacosRegisterInstance := vo.RegisterInstanceParam{
		Ip:          "10.0.0.10", //require user provide
		Port:        8848,        //require user provide
		ServiceName: "demo.go",   //require user provide
		Weight:      10,          //require user provide
		Enable:      true,        //require user provide
		Healthy:     true,        //require user provide
		Ephemeral:   true,        //require user provide
		Metadata:    map[string]string{"contextPath": "contextPath", "uriMetadata": string(metaDataStringJson)},
	}

	client, err := nacos_client.NewNacosClient(ncp)
	if err != nil {
		logger.Fatal("create nacos client error : %+V", err)
	}

	registerResult, err := nacos_client.RegisterNacosInstance(client, nacosRegisterInstance)
	if !registerResult && err != nil {
		logger.Fatal("Register nacos Instance error : %+V", err)
	}

	//do your logic
}
