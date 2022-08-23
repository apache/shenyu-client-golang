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
	"github.com/apache/shenyu-client-golang/clients/zk_client"
	"github.com/apache/shenyu-client-golang/common/constants"
	"github.com/apache/shenyu-client-golang/common/shenyu_sdk_client"
	"github.com/apache/shenyu-client-golang/model"
	"time"
)

func main() {

	//Create ShenYuZkClient  start
	zcp := &zk_client.ZkClientParam{
		ZkServers: []string{"127.0.0.1:2181"}, //require user provide
		ZkRoot:    "/api",                     //require user provide
	}

	sdkClient := shenyu_sdk_client.GetFactoryClient(constants.ZOOKEEPER_CLIENT)
	client, createResult, err := sdkClient.NewClient(zcp)

	if !createResult && err != nil {
		fmt.Printf("Create ShenYuZkClient error : %v", err)
	}

	zc := client.(*zk_client.ShenYuZkClient)
	defer zc.Close()
	//Create ShenYuZkClient end

	//RegisterServiceInstance start
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

	metaData3 := &model.MetaDataRegister{
		AppName: "testMetaDataRegister3", //require user provide
		Path:    "your/path3",            //require user provide
		Enabled: true,                    //require user provide
		Host:    "127.0.0.1",             //require user provide
		Port:    "8282",                  //require user provide
	}

	//register multiple metaData
	registerResult1, err := zc.RegisterServiceInstance(metaData1)
	if !registerResult1 && err != nil {
		fmt.Printf("Register zk Instance error : %v", err)
	}

	registerResult2, err := zc.RegisterServiceInstance(metaData2)
	if !registerResult2 && err != nil {
		fmt.Printf("Register zk Instance error : %v", err)
	}

	registerResult3, err := zc.RegisterServiceInstance(metaData3)
	if !registerResult3 && err != nil {
		fmt.Printf("Register zk Instance error : %v", err)
	}
	//RegisterServiceInstance end

	time.Sleep(time.Second)

	//GetServiceInstanceInfo start
	instanceDetail, err := zc.GetServiceInstanceInfo(metaData1)
	nodes1, ok := instanceDetail.([]*model.MetaDataRegister)
	if !ok {
		fmt.Printf("get zk client metaData error %v:", err)
	}

	//range nodes
	for index, node := range nodes1 {
		nodeJson, err := json.Marshal(node)
		if err == nil {
			fmt.Printf("GetNodesInfo ,success Index %v,%v", index, string(nodeJson))
		}
	}

	instanceDetail2, err := zc.GetServiceInstanceInfo(metaData2)
	nodes2, ok := instanceDetail2.([]*model.MetaDataRegister)
	if !ok {
		fmt.Printf("get zk client metaData error %v:", err)
	}

	//range nodes1
	for index, node := range nodes2 {
		nodeJson, err := json.Marshal(node)
		if err == nil {
			fmt.Printf("GetNodesInfo ,success Index %v,%v", index, string(nodeJson))
		}
	}

	instanceDetail3, err := zc.GetServiceInstanceInfo(metaData3)
	nodes3, ok := instanceDetail3.([]*model.MetaDataRegister)
	if !ok {
		fmt.Printf("get zk client metaData error %v:", err)
	}

	for index, node := range nodes3 {
		nodeJson, err := json.Marshal(node)
		if err == nil {
			fmt.Printf("GetNodesInfo ,success Index %v,%v", index, string(nodeJson))
		}
	}
	//GetServiceInstanceInfo end

	//DeregisterServiceInstance start
	//your can chose to invoke,not require
	fmt.Printf("> DeregisterServiceInstance start")
	deRegisterResult1, err := zc.DeregisterServiceInstance(metaData1)
	if err != nil {
		panic(err)
	}

	deRegisterResult2, err := zc.DeregisterServiceInstance(metaData2)
	if err != nil {
		panic(err)
	}

	deRegisterResult3, err := zc.DeregisterServiceInstance(metaData3)
	if err != nil {
		panic(err)
	}

	if deRegisterResult1 && deRegisterResult2 && deRegisterResult3 {
		fmt.Printf("DeregisterServiceInstance success !")
	}
	//DeregisterServiceInstance end

	//do your logic

}
