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
	"github.com/apache/shenyu-client-golang/clients/zk_client"
	"github.com/apache/shenyu-client-golang/common/constants"
	"github.com/apache/shenyu-client-golang/common/shenyu_sdk_client"
	"github.com/apache/shenyu-client-golang/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

/**
 * TestInitZkClient
 **/
func TestInitZkClient(t *testing.T) {
	zcp := &zk_client.ZkClientParam{
		ZkServers: []string{"127.0.0.1:2181"}, //require user provide
		ZkRoot:    "/api",                     //require user provide
	}

	sdkClient := shenyu_sdk_client.GetFactoryClient(constants.ZOOKEEPER_CLIENT)
	client, createResult, err := sdkClient.NewClient(zcp)

	assert.NotNil(t, client)
	assert.True(t, createResult)
	assert.Nil(t, err)
	zc := client.(*zk_client.ShenYuZkClient)
	defer zc.Close()
}

/**
 * TestRegisterServiceInstanceAndGetServiceInstanceInfo
 **/
func TestRegisterServiceInstanceAndGetServiceInstanceInfo(t *testing.T) {
	zcp := &zk_client.ZkClientParam{
		ZkServers: []string{"127.0.0.1:2181"}, //require user provide
		ZkRoot:    "/api",                     //require user provide
	}

	sdkClient := shenyu_sdk_client.GetFactoryClient(constants.ZOOKEEPER_CLIENT)
	client, createResult, err := sdkClient.NewClient(zcp)

	assert.NotNil(t, client)
	assert.True(t, createResult)
	assert.Nil(t, err)

	zc := client.(*zk_client.ShenYuZkClient)
	defer zc.Close()

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
	assert.Nil(t, err)
	assert.True(t, registerResult1)

	registerResult2, err := zc.RegisterServiceInstance(metaData2)
	assert.Nil(t, err)
	assert.True(t, registerResult2)

	registerResult3, err := zc.RegisterServiceInstance(metaData3)
	assert.Nil(t, err)
	assert.True(t, registerResult3)

	time.Sleep(time.Second)

	instanceDetail, err := zc.GetServiceInstanceInfo(metaData1)
	assert.NotNil(t, instanceDetail)
	assert.Nil(t, err)

	instanceDetail2, err := zc.GetServiceInstanceInfo(metaData2)
	assert.NotNil(t, instanceDetail2)
	assert.Nil(t, err)

	instanceDetail3, err := zc.GetServiceInstanceInfo(metaData3)
	assert.NotNil(t, instanceDetail3)
	assert.Nil(t, err)
}

/**
* TestRegisterInstanceAndDeregisterServiceInstance
**/
func TestDeregisterServiceInstance(t *testing.T) {
	zcp := &zk_client.ZkClientParam{
		ZkServers: []string{"127.0.0.1:2181"}, //require user provide
		ZkRoot:    "/api",                     //require user provide
	}

	sdkClient := shenyu_sdk_client.GetFactoryClient(constants.ZOOKEEPER_CLIENT)
	client, createResult, err := sdkClient.NewClient(zcp)

	assert.NotNil(t, client)
	assert.True(t, createResult)
	assert.Nil(t, err)

	zc := client.(*zk_client.ShenYuZkClient)
	//defer zc.Close()

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

	deRegisterResult1, err := zc.DeregisterServiceInstance(metaData1)
	assert.Nil(t, err)
	assert.True(t, deRegisterResult1)

	deRegisterResult2, err := zc.DeregisterServiceInstance(metaData2)
	assert.Nil(t, err)
	assert.True(t, deRegisterResult2)

	deRegisterResult3, err := zc.DeregisterServiceInstance(metaData3)
	assert.Nil(t, err)
	assert.True(t, deRegisterResult3)

}
