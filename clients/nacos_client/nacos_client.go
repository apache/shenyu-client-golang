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
	oriNc "github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/sirupsen/logrus"
)

var (
	logger = logrus.New()
)

/**
 * ShenYuNacosClient
 **/
type ShenYuNacosClient struct {
	NacosClient *naming_client.NamingClient
}

/**
 * NacosClientParam
 **/
type NacosClientParam struct {
	IpAddr      string //the nacos server address require user provide
	Port        uint64 //the nacos server port require user provide
	NamespaceId string //the namespaceId of Nacos require user provide.
}

/**
 * create nacos client
 **/
func (nc *ShenYuNacosClient) NewClient(clientParam interface{}) (client interface{}, createResult bool, err error) {
	ncp, ok := clientParam.(*NacosClientParam)
	if !ok {
		logger.Fatalf("init nacos client error %v:", err)
	}
	checkResult := len(ncp.IpAddr) > 0 && len(ncp.NamespaceId) > 0 && ncp.Port > 0
	if checkResult {
		client, err := ncp.initNacosClient()
		if err != nil {
			logger.Fatalf("init nacos client error %v:", err)
		}
		return client, true, nil
	} else {
		logger.Fatalf("init nacos client param is missing please check")
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
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
	)

	client, err := oriNc.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err == nil {
		return client, nil
	}
	return
}

/**
 * Register Instance to Nacos
 **/
func (nc *ShenYuNacosClient) RegisterServiceInstance(metaData interface{}) (registerResult bool, err error) {
	rip, ok := metaData.(vo.RegisterInstanceParam)
	if !ok {
		logger.Fatalf("init nacos client error %v:", err)
	}
	registerResult, err = nc.NacosClient.RegisterInstance(rip)
	if err != nil {
		logger.Fatalf("RegisterServiceInstance failure! ,error is :%v", err)
	}
	logger.Infof("RegisterServiceInstance,result:%v\n\n,param:%v \n\n", registerResult, rip)
	return registerResult, nil
}

/**
 * DeregisterServiceInstance
 **/
func (nc *ShenYuNacosClient) DeregisterServiceInstance(metaData interface{}) (deRegisterResult bool, err error) {
	rip, ok := metaData.(vo.DeregisterInstanceParam)
	if !ok {
		logger.Fatalf("init nacos client error %v:", err)
	}
	deRegisterResult, err = nc.NacosClient.DeregisterInstance(rip)
	if err != nil {
		logger.Fatalf("DeregisterServiceInstance failure! ,error is :%v", err)
	}
	logger.Infof("DeregisterServiceInstance,result:%v\n\n,param:%v \n\n", deRegisterResult, rip)
	return deRegisterResult, nil
}

/**
 * GetServiceInstanceInfo
 **/
func (nc *ShenYuNacosClient) GetServiceInstanceInfo(metaData interface{}) (instances interface{}, err error) {
	rip, ok := metaData.(vo.SelectInstancesParam)
	if !ok {
		logger.Fatalf("init nacos client error %v:", err)
	}
	instances, err = nc.NacosClient.SelectInstances(rip)
	if err != nil {
		logger.Fatalf("GetServiceInstanceInfo failure! ,error is :%v", err)
	}
	logger.Infof("GetServiceInstanceInfo,result:%v\n\n,param:%v \n\n", instances, rip)
	return instances, nil
}
