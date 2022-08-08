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
	"github.com/hashicorp/go-uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

/**
 * TestInitConsulClient
 **/
func TestInitConsulClient(t *testing.T) {
	ccp := &consul_client.ConsulClientParam{
		Host:  "127.0.0.1", //require user provide
		Port:  8500,        //require user provide
		Token: "",
	}

	sdkClient := shenyu_sdk_client.GetFactoryClient(constants.CONSUL_CLIENT)
	client, createResult, err := sdkClient.NewClient(ccp)

	assert.NotNil(t, client)
	assert.True(t, createResult)
	assert.Nil(t, err)
}

/**
 * TestRegisterServiceInstance
 **/
func TestRegisterServiceInstance(t *testing.T) {
	ccp := &consul_client.ConsulClientParam{
		Host:  "127.0.0.1", //require user provide
		Port:  8500,        //require user provide
		Token: "",
	}

	sdkClient := shenyu_sdk_client.GetFactoryClient(constants.CONSUL_CLIENT)
	client, createResult, err := sdkClient.NewClient(ccp)

	assert.Nil(t, err)
	assert.True(t, createResult)

	scc := client.(*consul_client.ShenYuConsulClient)
	//Create ShenYuConsulClient end
	uuid1, _ := uuid.GenerateUUID()
	uuid2, _ := uuid.GenerateUUID()
	uuid3, _ := uuid.GenerateUUID()

	//RegisterServiceInstance start
	//init MetaDataRegister
	metaData1 := &model.ConsulMetaDataRegister{
		ServiceId: uuid1,
		ShenYuMetaData: &model.MetaDataRegister{
			AppName: "testMetaDataRegister1", //require user provide
			Path:    "/your/path1",           //require user provide
			Enabled: true,                    //require user provide
			Host:    "127.0.0.1",             //require user provide
			Port:    "8080",                  //require user provide
			RPCType: "http",                  //require user provide
		},
	}

	metaData2 := &model.ConsulMetaDataRegister{
		ServiceId: uuid2,
		ShenYuMetaData: &model.MetaDataRegister{
			AppName: "testMetaDataRegister2", //require user provide
			Path:    "/your/path2",           //require user provide
			Enabled: true,                    //require user provide
			Host:    "127.0.0.1",             //require user provide
			Port:    "8181",                  //require user provide
			RPCType: "http",                  //require user provide
		},
	}

	metaData3 := &model.ConsulMetaDataRegister{
		ServiceId: uuid3,
		ShenYuMetaData: &model.MetaDataRegister{
			AppName: "testMetaDataRegister3", //require user provide
			Path:    "/your/path3",           //require user provide
			Enabled: true,                    //require user provide
			Host:    "127.0.0.1",             //require user provide
			Port:    "8282",                  //require user provide
			RPCType: "http",                  //require user provide
		},
	}

	//register multiple metaData
	registerResult1, err := scc.RegisterServiceInstance(metaData1)
	assert.Nil(t, err)
	assert.True(t, registerResult1)

	registerResult2, err := scc.RegisterServiceInstance(metaData2)
	assert.Nil(t, err)
	assert.True(t, registerResult2)

	registerResult3, err := scc.RegisterServiceInstance(metaData3)
	assert.Nil(t, err)
	assert.True(t, registerResult3)
}

/**
* TestDeregisterServiceInstance
**/
func TestDeregisterServiceInstance(t *testing.T) {
	ccp := &consul_client.ConsulClientParam{
		Host:  "127.0.0.1", //require user provide
		Port:  8500,        //require user provide
		Token: "",
	}

	sdkClient := shenyu_sdk_client.GetFactoryClient(constants.CONSUL_CLIENT)
	client, createResult, err := sdkClient.NewClient(ccp)

	assert.Nil(t, err)
	assert.True(t, createResult)

	scc := client.(*consul_client.ShenYuConsulClient)
	//Create ShenYuConsulClient end
	uuid1, _ := uuid.GenerateUUID()
	uuid2, _ := uuid.GenerateUUID()
	uuid3, _ := uuid.GenerateUUID()

	//RegisterServiceInstance start
	//init MetaDataRegister
	metaData1 := &model.ConsulMetaDataRegister{
		ServiceId: uuid1,
		ShenYuMetaData: &model.MetaDataRegister{
			AppName: "testMetaDataRegister1", //require user provide
			Path:    "/your/path1",           //require user provide
			Enabled: true,                    //require user provide
			Host:    "127.0.0.1",             //require user provide
			Port:    "8080",                  //require user provide
			RPCType: "http",                  //require user provide
		},
	}

	metaData2 := &model.ConsulMetaDataRegister{
		ServiceId: uuid2,
		ShenYuMetaData: &model.MetaDataRegister{
			AppName: "testMetaDataRegister2", //require user provide
			Path:    "/your/path2",           //require user provide
			Enabled: true,                    //require user provide
			Host:    "127.0.0.1",             //require user provide
			Port:    "8181",                  //require user provide
			RPCType: "http",                  //require user provide
		},
	}

	metaData3 := &model.ConsulMetaDataRegister{
		ServiceId: uuid3,
		ShenYuMetaData: &model.MetaDataRegister{
			AppName: "testMetaDataRegister3", //require user provide
			Path:    "/your/path3",           //require user provide
			Enabled: true,                    //require user provide
			Host:    "127.0.0.1",             //require user provide
			Port:    "8282",                  //require user provide
			RPCType: "http",                  //require user provide
		},
	}

	//register multiple metaData
	registerResult1, err := scc.RegisterServiceInstance(metaData1)
	assert.Nil(t, err)
	assert.True(t, registerResult1)

	registerResult2, err := scc.RegisterServiceInstance(metaData2)
	assert.Nil(t, err)
	assert.True(t, registerResult2)

	registerResult3, err := scc.RegisterServiceInstance(metaData3)
	assert.Nil(t, err)
	assert.True(t, registerResult3)

	time.Sleep(time.Second)

	//DeregisterServiceInstance
	deRegisterResult1, err := scc.DeregisterServiceInstance(metaData1)
	assert.Nil(t, err)
	assert.True(t, deRegisterResult1)

	deRegisterResult2, err := scc.DeregisterServiceInstance(metaData2)
	assert.Nil(t, err)
	assert.True(t, deRegisterResult2)

	deRegisterResult3, err := scc.DeregisterServiceInstance(metaData3)
	assert.Nil(t, err)
	assert.True(t, deRegisterResult3)

}

/**
* TestGetServiceInstanceInfo
**/
func TestGetServiceInstanceInfo(t *testing.T) {
	ccp := &consul_client.ConsulClientParam{
		Host:  "127.0.0.1", //require user provide
		Port:  8500,        //require user provide
		Token: "",
	}

	sdkClient := shenyu_sdk_client.GetFactoryClient(constants.CONSUL_CLIENT)
	client, createResult, err := sdkClient.NewClient(ccp)

	assert.Nil(t, err)
	assert.True(t, createResult)

	scc := client.(*consul_client.ShenYuConsulClient)
	//Create ShenYuConsulClient end
	uuid1, _ := uuid.GenerateUUID()
	uuid2, _ := uuid.GenerateUUID()
	uuid3, _ := uuid.GenerateUUID()

	//RegisterServiceInstance start
	//init MetaDataRegister
	metaData1 := &model.ConsulMetaDataRegister{
		ServiceId: uuid1,
		ShenYuMetaData: &model.MetaDataRegister{
			AppName: "testMetaDataRegister1", //require user provide
			Path:    "/your/path1",           //require user provide
			Enabled: true,                    //require user provide
			Host:    "127.0.0.1",             //require user provide
			Port:    "8080",                  //require user provide
			RPCType: "http",                  //require user provide
		},
	}

	metaData2 := &model.ConsulMetaDataRegister{
		ServiceId: uuid2,
		ShenYuMetaData: &model.MetaDataRegister{
			AppName: "testMetaDataRegister2", //require user provide
			Path:    "/your/path2",           //require user provide
			Enabled: true,                    //require user provide
			Host:    "127.0.0.1",             //require user provide
			Port:    "8181",                  //require user provide
			RPCType: "http",                  //require user provide
		},
	}

	metaData3 := &model.ConsulMetaDataRegister{
		ServiceId: uuid3,
		ShenYuMetaData: &model.MetaDataRegister{
			AppName: "testMetaDataRegister3", //require user provide
			Path:    "/your/path3",           //require user provide
			Enabled: true,                    //require user provide
			Host:    "127.0.0.1",             //require user provide
			Port:    "8282",                  //require user provide
			RPCType: "http",                  //require user provide
		},
	}

	//register multiple metaData
	registerResult1, err := scc.RegisterServiceInstance(metaData1)
	assert.Nil(t, err)
	assert.True(t, registerResult1)

	registerResult2, err := scc.RegisterServiceInstance(metaData2)
	assert.Nil(t, err)
	assert.True(t, registerResult2)

	registerResult3, err := scc.RegisterServiceInstance(metaData3)
	assert.Nil(t, err)
	assert.True(t, registerResult3)

	time.Sleep(time.Second)

	//get nodes
	instanceDetail, err := scc.GetServiceInstanceInfo(metaData1)
	nodes1, ok := instanceDetail.([]*model.ConsulMetaDataRegister)
	assert.NotNil(t, nodes1)
	assert.True(t, ok)
	assert.Nil(t, err)

	instanceDetail2, err := scc.GetServiceInstanceInfo(metaData2)
	nodes2, ok := instanceDetail2.([]*model.ConsulMetaDataRegister)
	assert.NotNil(t, nodes2)
	assert.True(t, ok)
	assert.Nil(t, err)

	instanceDetail3, err := scc.GetServiceInstanceInfo(metaData3)
	nodes3, ok := instanceDetail3.([]*model.ConsulMetaDataRegister)
	assert.NotNil(t, nodes3)
	assert.True(t, ok)
	assert.Nil(t, err)

}

/**
* TestEntireConsulFunction
**/
func TestEntireConsulFunction(t *testing.T) {
	ccp := &consul_client.ConsulClientParam{
		Host:  "127.0.0.1", //require user provide
		Port:  8500,        //require user provide
		Token: "",
	}

	sdkClient := shenyu_sdk_client.GetFactoryClient(constants.CONSUL_CLIENT)
	client, createResult, err := sdkClient.NewClient(ccp)

	assert.Nil(t, err)
	assert.True(t, createResult)

	scc := client.(*consul_client.ShenYuConsulClient)
	//Create ShenYuConsulClient end
	uuid1, _ := uuid.GenerateUUID()
	uuid2, _ := uuid.GenerateUUID()
	uuid3, _ := uuid.GenerateUUID()

	//RegisterServiceInstance start
	//init MetaDataRegister
	metaData1 := &model.ConsulMetaDataRegister{
		ServiceId: uuid1,
		ShenYuMetaData: &model.MetaDataRegister{
			AppName: "testEntireMetaDataRegister1", //require user provide
			Path:    "/your/entire/path1",          //require user provide
			Enabled: true,                          //require user provide
			Host:    "127.0.0.1",                   //require user provide
			Port:    "8080",                        //require user provide
			RPCType: "http",                        //require user provide
		},
	}

	metaData2 := &model.ConsulMetaDataRegister{
		ServiceId: uuid2,
		ShenYuMetaData: &model.MetaDataRegister{
			AppName: "testEntireMetaDataRegister2", //require user provide
			Path:    "/your/entire/path2",          //require user provide
			Enabled: true,                          //require user provide
			Host:    "127.0.0.1",                   //require user provide
			Port:    "8181",                        //require user provide
			RPCType: "http",                        //require user provide
		},
	}

	metaData3 := &model.ConsulMetaDataRegister{
		ServiceId: uuid3,
		ShenYuMetaData: &model.MetaDataRegister{
			AppName: "testEntireMetaDataRegister3", //require user provide
			Path:    "/your/entire/path3",          //require user provide
			Enabled: true,                          //require user provide
			Host:    "127.0.0.1",                   //require user provide
			Port:    "8282",                        //require user provide
			RPCType: "http",                        //require user provide
		},
	}

	//register multiple metaData
	registerResult1, err := scc.RegisterServiceInstance(metaData1)
	assert.Nil(t, err)
	assert.True(t, registerResult1)

	registerResult2, err := scc.RegisterServiceInstance(metaData2)
	assert.Nil(t, err)
	assert.True(t, registerResult2)

	registerResult3, err := scc.RegisterServiceInstance(metaData3)
	assert.Nil(t, err)
	assert.True(t, registerResult3)

	time.Sleep(time.Second)

	//get nodes
	instanceDetail1, err := scc.GetServiceInstanceInfo(metaData1)
	nodes1, ok := instanceDetail1.([]*model.ConsulMetaDataRegister)
	assert.NotNil(t, nodes1)
	assert.True(t, ok)
	assert.Nil(t, err)

	instanceDetail2, err := scc.GetServiceInstanceInfo(metaData2)
	nodes2, ok := instanceDetail2.([]*model.ConsulMetaDataRegister)
	assert.NotNil(t, nodes2)
	assert.True(t, ok)
	assert.Nil(t, err)

	instanceDetail3, err := scc.GetServiceInstanceInfo(metaData3)
	nodes3, ok := instanceDetail3.([]*model.ConsulMetaDataRegister)
	assert.NotNil(t, nodes3)
	assert.True(t, ok)
	assert.Nil(t, err)

	//DeregisterServiceInstance
	deRegisterResult1, err := scc.DeregisterServiceInstance(metaData1)
	assert.Nil(t, err)
	assert.True(t, deRegisterResult1)

	deRegisterResult2, err := scc.DeregisterServiceInstance(metaData2)
	assert.Nil(t, err)
	assert.True(t, deRegisterResult2)

	deRegisterResult3, err := scc.DeregisterServiceInstance(metaData3)
	assert.Nil(t, err)
	assert.True(t, deRegisterResult3)

}
