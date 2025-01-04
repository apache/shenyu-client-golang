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
	"github.com/apache/shenyu-client-golang/clients/nacos_client"
	"github.com/apache/shenyu-client-golang/common/constants"
	"github.com/apache/shenyu-client-golang/common/shenyu_sdk_client"
	"github.com/apache/shenyu-client-golang/model"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

/**
 * TestInitNacosClient
 **/
func TestInitNacosClient(t *testing.T) {
	//set nacos env configuration
	ncp := &nacos_client.NacosClientParam{
		IpAddr:      "127.0.0.1",
		Port:        8848,
		NamespaceId: "public",
	}

	sdkClient := shenyu_sdk_client.GetFactoryClient(constants.NACOS_CLIENT)
	client, createResult, err := sdkClient.NewClient(ncp)

	assert.NotNil(t, client)
	assert.True(t, createResult)
	assert.Nil(t, err)
}

/**
 * TestInitNacosClientAndRegister
 **/
func TestInitNacosClientAndRegister(t *testing.T) {
	//set nacos env configuration
	ncp := &nacos_client.NacosClientParam{
		IpAddr:      "127.0.0.1",
		Port:        8848,
		NamespaceId: "public",
	}

	sdkClient := shenyu_sdk_client.GetFactoryClient(constants.NACOS_CLIENT)
	client, createResult, err := sdkClient.NewClient(ncp)

	assert.NotNil(t, client)
	assert.True(t, createResult)
	assert.Nil(t, err)

	nc := &nacos_client.ShenYuNacosClient{
		NacosClient: client.(*naming_client.NamingClient),
	}

	//metaData is necessary param, this will be register to shenyu gateway to use
	metaData := &model.URIRegister{
		Protocol:     "testMetaDataRegister",                 //require user provide
		AppName:      "testURLRegister",                      //require user provide
		ContextPath:  "contextPath",                          //require user provide
		RPCType:      constants.RPCTYPE_HTTP,                 //require user provide
		Host:         "127.0.0.1",                            //require user provide
		Port:         "8080",                                 //require user provide
		NamespaceIds: "649330b6-c2d7-4edc-be8e-8a54df9eb385", //require user provide
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
		ClusterName: "DEFAULT",
		GroupName:   "DEFAULT_GROUP",
	}

	instance, err := nc.RegisterServiceInstance(nacosRegisterInstance)
	assert.Nil(t, err)
	assert.NotNil(t, instance)
}

/**
 * TestRegisterAndGetInstance
 **/
func TestRegisterAndGetInstance(t *testing.T) {
	//set nacos env configuration
	ncp := &nacos_client.NacosClientParam{
		IpAddr:      "127.0.0.1",
		Port:        8848,
		NamespaceId: "public",
	}

	sdkClient := shenyu_sdk_client.GetFactoryClient(constants.NACOS_CLIENT)
	client, createResult, err := sdkClient.NewClient(ncp)

	assert.NotNil(t, client)
	assert.True(t, createResult)
	assert.Nil(t, err)

	nc := &nacos_client.ShenYuNacosClient{
		NacosClient: client.(*naming_client.NamingClient),
	}

	//metaData is necessary param, this will be register to shenyu gateway to use
	metaData := &model.URIRegister{
		Protocol:     "testMetaDataRegister",                 //require user provide
		AppName:      "testURLRegister",                      //require user provide
		ContextPath:  "contextPath",                          //require user provide
		RPCType:      constants.RPCTYPE_HTTP,                 //require user provide
		Host:         "127.0.0.1",                            //require user provide
		Port:         "8080",                                 //require user provide
		NamespaceIds: "649330b6-c2d7-4edc-be8e-8a54df9eb385", //require user provide
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
		GroupName:   "group-a",   //require user provide
		Metadata:    map[string]string{"contextPath": "contextPath", "uriMetadata": string(metaDataStringJson)},
	}

	instance, err := nc.RegisterServiceInstance(nacosRegisterInstance)
	assert.Nil(t, err)
	assert.NotNil(t, instance)

	time.Sleep(time.Second)

	queryData := vo.SelectInstancesParam{
		ServiceName: "demo.go",
		GroupName:   "group-a", //default: DEFAULT_GROUP
		//Clusters:    []string{"cluster-a"}, // default: DEFAULT
		HealthyOnly: true,
	}

	instanceInfo, err := nc.GetServiceInstanceInfo(queryData)
	assert.Nil(t, err)
	assert.NotNil(t, instanceInfo)
}

/**
 * TestRegisterAndDeregister
 **/
func TestRegisterAndDeregister(t *testing.T) {
	//set nacos env configuration
	ncp := &nacos_client.NacosClientParam{
		IpAddr:      "127.0.0.1",
		Port:        8848,
		NamespaceId: "public",
	}

	sdkClient := shenyu_sdk_client.GetFactoryClient(constants.NACOS_CLIENT)
	client, createResult, err := sdkClient.NewClient(ncp)

	assert.NotNil(t, client)
	assert.True(t, createResult)
	assert.Nil(t, err)

	nc := &nacos_client.ShenYuNacosClient{
		NacosClient: client.(*naming_client.NamingClient),
	}

	//metaData is necessary param, this will be register to shenyu gateway to use
	metaData := &model.URIRegister{
		Protocol:     "testMetaDataRegister",                 //require user provide
		AppName:      "testURLRegister",                      //require user provide
		ContextPath:  "contextPath",                          //require user provide
		RPCType:      constants.RPCTYPE_HTTP,                 //require user provide
		Host:         "127.0.0.1",                            //require user provide
		Port:         "8080",                                 //require user provide
		NamespaceIds: "649330b6-c2d7-4edc-be8e-8a54df9eb385", //require user provide
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
		GroupName:   "group-a",   //require user provide
		Metadata:    map[string]string{"contextPath": "contextPath", "uriMetadata": string(metaDataStringJson)},
	}

	instance, err := nc.RegisterServiceInstance(nacosRegisterInstance)
	assert.Nil(t, err)
	assert.NotNil(t, instance)

	time.Sleep(time.Second)
	deregisterInstanceParam := vo.DeregisterInstanceParam{
		Ip:          "10.0.0.10",
		Port:        8848,
		ServiceName: "demo.go",
		Ephemeral:   true,
		//Cluster:     "cluster-a", // default value is DEFAULT
		GroupName: "group-a", // default value is DEFAULT_GROUP
	}

	serviceInstance, err := nc.DeregisterServiceInstance(deregisterInstanceParam)
	assert.Nil(t, err)
	assert.True(t, serviceInstance)
}
