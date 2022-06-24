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
	"github.com/apache/incubator-shenyu-client-golang/clients/nacos_client"
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

	//init NacosRegisterInstance
	nacosRegisterInstance := model.NacosRegisterInstance{
		RegisterInstance: vo.RegisterInstanceParam{
			Ip:          "10.0.0.10",
			Port:        8848,
			ServiceName: "demo.go",
			Weight:      10,
			Enable:      true,
			Healthy:     true,
			Ephemeral:   true,
			Metadata:    map[string]string{"idc": "beijing"},
		},
	}

	client, err := nacos_client.NewNacosClient(ncp)
	if err != nil {
		logger.Fatal("create nacos client error : %+V", err)
	}

	registerResult, err := nacos_client.RegisterNacosInstance(client, nacosRegisterInstance)
	if !registerResult && err != nil {
		logger.Fatal("Register nacos Instance error : %+V", err)
	}

}
