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
	"github.com/apache/incubator-shenyu-client-golang/clients/zk_client"
	"github.com/apache/incubator-shenyu-client-golang/model"
	"github.com/wonderivan/logger"
)

func main() {
	servers := []string{"127.0.0.1:2181"}                   //require user provide
	client, err := zk_client.NewClient(servers, "/api", 10) //zkRoot require user provide
	if err != nil {
		panic(err)
	}
	defer client.Close()

	//init MetaDataRegister
	metaData1 := &model.MetaDataRegister{
		AppName: "testMetaDataRegister", //require user provide
		Path:    "your/path1",           //require user provide
		Enabled: true,                   //require user provide
		Host:    "127.0.0.1",            //require user provide
		Port:    "8080",                 //require user provide
	}

	metaData2 := &model.MetaDataRegister{
		AppName: "testMetaDataRegister", //require user provide
		Path:    "your/path2",           //require user provide
		Enabled: true,                   //require user provide
		Host:    "127.0.0.1",            //require user provide
		Port:    "8181",                 //require user provide
	}

	metaData3 := &model.MetaDataRegister{
		AppName: "testMetaDataRegister", //require user provide
		Path:    "your/path3",           //require user provide
		Enabled: true,                   //require user provide
		Host:    "127.0.0.1",            //require user provide
		Port:    "8282",                 //require user provide
	}

	//register multiple metaData
	if err := client.RegisterNodeInstance(metaData1); err != nil {
		panic(err)
	}
	if err := client.RegisterNodeInstance(metaData2); err != nil {
		panic(err)
	}
	if err := client.RegisterNodeInstance(metaData3); err != nil {
		panic(err)
	}

	nodes, err := client.GetNodesInfo("testMetaDataRegister")
	if err != nil {
		panic(err)
	}

	//range nodes
	for index, node := range nodes {
		nodeJson, err := json.Marshal(node)
		if err == nil {
			logger.Info("GetNodesInfo ,success Index", index, string(nodeJson))
		}
	}

	//your can chose to invoke,not require
	err = client.DeleteNodeInstance(metaData1)
	if err != nil {
		panic(err)
	}

	err = client.DeleteNodeInstance(metaData2)
	if err != nil {
		panic(err)
	}

	err = client.DeleteNodeInstance(metaData3)
	if err != nil {
		panic(err)
	}

}
