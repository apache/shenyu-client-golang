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
	"context"
	"github.com/apache/shenyu-client-golang/clients/etcd_client"
	"github.com/apache/shenyu-client-golang/common/constants"
	"github.com/apache/shenyu-client-golang/common/shenyu_sdk_client"
	"github.com/apache/shenyu-client-golang/model"
	"github.com/stretchr/testify/assert"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
	"time"
)

/**
 * TestInitEtcdClient
 **/
func TestInitEtcdClient(t *testing.T) {
	ccp := &etcd_client.EtcdClientParam{
		EtcdServers: []string{"http://127.0.0.1:2379"}, //require user provide
		TTL:    50,
		TimeOut: 100000,
	}

	sdkClient := shenyu_sdk_client.GetFactoryClient(constants.ETCD_CLIENT)
	client, createResult, err := sdkClient.NewClient(ccp)

	assert.NotNil(t, client)
	assert.True(t, createResult)
	assert.Nil(t, err)
	etcd := client.(*etcd_client.ShenYuEtcdClient)
	////leaseId
	//var leaseId = etcd.GetLeaseId()
	//etcd.GlobalLease = leaseId
	//go func() {
	//	etcd.KeepAlive()
	//}()
	defer etcd.Close()
}

/**
 * TestRegisterServiceInstanceAndGetServiceInstanceInfo
 **/
func TestRegisterServiceInstanceAndGetServiceInstanceInfo(t *testing.T) {
	ccp := &etcd_client.EtcdClientParam{
		EtcdServers: []string{"http://127.0.0.1:2379"}, //require user provide
		TTL:    50,
		TimeOut: 100000,
	}

	sdkClient := shenyu_sdk_client.GetFactoryClient(constants.ETCD_CLIENT)
	client, createResult, err := sdkClient.NewClient(ccp)

	assert.NotNil(t, client)
	assert.True(t, createResult)
	assert.Nil(t, err)

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
	assert.Nil(t, err)
	assert.True(t, registerResult1)

	registerResult2, err := etcd.RegisterServiceInstance(metaData2)
	assert.Nil(t, err)
	assert.True(t, registerResult2)


	time.Sleep(time.Second)

	instanceDetail, err := etcd.GetServiceInstanceInfo(metaData1)
	assert.NotNil(t, instanceDetail)
	assert.Nil(t, err)

	instanceDetail2, err := etcd.GetServiceInstanceInfo(metaData2)
	assert.NotNil(t, instanceDetail2)
	assert.Nil(t, err)

}

/**
 * TestDeRegisterServiceInstance
 **/
func TestDeRegisterServiceInstance(t *testing.T) {
	ccp := &etcd_client.EtcdClientParam{
		EtcdServers: []string{"http://127.0.0.1:2379"}, //require user provide
		TTL:    50,
		TimeOut: 100000,
	}

	sdkClient := shenyu_sdk_client.GetFactoryClient(constants.ETCD_CLIENT)
	client, createResult, err := sdkClient.NewClient(ccp)

	assert.NotNil(t, client)
	assert.True(t, createResult)
	assert.Nil(t, err)

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
	registerResult1, err := etcd.DeregisterServiceInstance(metaData1)
	assert.Nil(t, err)
	assert.True(t, registerResult1)

	registerResult2, err := etcd.DeregisterServiceInstance(metaData2)
	assert.Nil(t, err)
	assert.True(t, registerResult2)
}

/**
** TestGenAndGetLeaseId
 */
func TestGenAndGetLeaseId(t *testing.T){
	ccp := &etcd_client.EtcdClientParam{
		EtcdServers: []string{"http://127.0.0.1:2379"}, //require user provide
		TTL:    50,
		TimeOut: 100000,
	}

	sdkClient := shenyu_sdk_client.GetFactoryClient(constants.ETCD_CLIENT)
	client, createResult, err := sdkClient.NewClient(ccp)

	assert.NotNil(t, client)
	assert.True(t, createResult)
	assert.Nil(t, err)

	etcd := client.(*etcd_client.ShenYuEtcdClient)
	defer etcd.Close()

	leaseId := etcd.GenLeaseId()

	etcd.GlobalLease = leaseId
	etcd.EtcdClient.Put(context.TODO(),"key1","key111",clientv3.WithLease(etcd.GlobalLease))
	etcd.EtcdClient.Put(context.TODO(),"key2","key13333",clientv3.WithLease(etcd.GlobalLease))
	//rent
	go func() {
		etcd.KeepAlive()
	}()
	select {

	}
}


