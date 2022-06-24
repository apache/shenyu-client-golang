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

package nacos_client

import (
	"github.com/apache/incubator-shenyu-client-golang/model"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/wonderivan/logger"
)

/**
 * NacosClientParam
 **/
type NacosClientParam struct {
	IpAddr      string //the nacos server address
	Port        uint64 //the nacos server port
	NamespaceId string // the namespaceId of Nacos.When namespace is public, fill in the blank string here.
}

/**
 * create nacos client
 **/
func newNacosClient(ncp *NacosClientParam) (clientProxy naming_client.INamingClient, err error) {
	checkResult := len(ncp.IpAddr) > 0 && len(ncp.NamespaceId) > 0 && ncp.Port > 0
	if checkResult {
		client, err := ncp.initNacosClient()
		if err != nil {
			logger.Fatal("init nacos client error %+v:", err)
		}
		return client, nil
	} else {
		logger.Fatal("init nacos client param is missing please check")
	}
	return
}

/**
 * use NacosClientParam to init client
 **/
func (ncp *NacosClientParam) initNacosClient() (clientProxy naming_client.INamingClient, err error) {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(ncp.IpAddr, ncp.Port),
	}

	//init ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(ncp.NamespaceId),
		constant.WithLogLevel("info"),
	)

	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		return client, nil
	}
	return
}

/**
 * register nacos instance
 **/
func registerNacosInstance(client naming_client.INamingClient, nri model.NacosRegisterInstance) (registerResult bool, err error) {
	registerResult, err = client.RegisterInstance(nri.RegisterInstance)
	if err != nil {
		logger.Fatal("RegisterServiceInstance failure! ,error is :%+v", err)
	}
	logger.Info("RegisterServiceInstance,param:%+v,result:%+v \n\n", nri.RegisterInstance, registerResult)
	return registerResult, nil
}
