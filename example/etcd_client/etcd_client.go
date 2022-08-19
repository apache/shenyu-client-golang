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
	"fmt"
	"github.com/apache/shenyu-client-golang/clients/etcd_client"
	"github.com/apache/shenyu-client-golang/common/constants"
	"github.com/apache/shenyu-client-golang/common/shenyu_sdk_client"
	"github.com/apache/shenyu-client-golang/model"
	"time"
)

func main() {
	ecp := &etcd_client.EtcdClientParam{
		EtcdServers: []string{"http://127.0.0.1:2379"}, //require user provide
		TTL:         50,
	}

	sdkClient := shenyu_sdk_client.GetFactoryClient(constants.ETCD_CLIENT)
	client, createResult, err := sdkClient.NewClient(ecp)
	if !createResult && err != nil {
		fmt.Printf("Create ShenYuEtcdClient error : %v", err)
	}

	etcd := client.(*etcd_client.ShenYuEtcdClient)
	defer etcd.Close()

	//init MetaDataRegister
	metaData1 := &model.MetaDataRegister{
		AppName: "testMetaDataRegister1", //require user provide
		Path:    "your/path1",            //require user provide
		Enabled: true,                    //require user provide
		Host:    "127.0.0.1",             //require user provide
		Port:    "8080",                  //require user provide
	}

	metaData2 := &model.MetaDataRegister{
		AppName: "testMetaDataRegister2", //require user provide
		Path:    "your/path2",            //require user provide
		Enabled: true,                    //require user provide
		Host:    "127.0.0.1",             //require user provide
		Port:    "8181",                  //require user provide
	}

	//register multiple metaData
	registerResult1, err := etcd.RegisterServiceInstance(metaData1)
	if !registerResult1 && err != nil {
		fmt.Printf("Register etcd Instance error : %v", err)
	}

	registerResult2, err := etcd.RegisterServiceInstance(metaData2)
	if !registerResult2 && err != nil {
		fmt.Printf("Register etcd Instance error : %v", err)
	}

	time.Sleep(time.Second)

	instanceDetail, err := etcd.GetServiceInstanceInfo(metaData1)
	nodes1, ok := instanceDetail.([]*model.MetaDataRegister)
	if !ok {
		fmt.Printf("get etcd client metaData error %v:", err)
	}

	//range nodes
	for index, node := range nodes1 {
		nodeJson, err := json.Marshal(node)
		if err == nil {
			fmt.Printf("GetNodesInfo ,success Index %v ,%v", index, string(nodeJson))
		}
	}

	instanceDetail2, err := etcd.GetServiceInstanceInfo(metaData2)
	nodes2, ok := instanceDetail2.([]*model.MetaDataRegister)
	if !ok {
		fmt.Printf("get etcd client metaData error %v:", err)
	}

	//range nodes
	for index, node := range nodes2 {
		nodeJson, err := json.Marshal(node)
		if err == nil {
			fmt.Printf("GetNodesInfo ,success Index %v ,%v", index, string(nodeJson))
		}
	}

	fmt.Printf("> DeregisterServiceInstance start")
	deRegisterResult1, err := etcd.DeregisterServiceInstance(metaData1)
	if err != nil {
		panic(err)
	}

	deRegisterResult2, err := etcd.DeregisterServiceInstance(metaData2)
	if err != nil {
		panic(err)
	}

	if deRegisterResult1 && deRegisterResult2 {
		fmt.Printf("DeregisterServiceInstance success !")
	}
	//DeregisterServiceInstance end

	//do your logic
}
