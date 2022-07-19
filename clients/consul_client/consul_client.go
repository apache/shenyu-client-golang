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

package consul_client

import (
	"encoding/json"
	"fmt"
	"github.com/apache/incubator-shenyu-client-golang/common/constants"
	"github.com/apache/incubator-shenyu-client-golang/model"
	"github.com/hashicorp/consul/api"
	"github.com/wonderivan/logger"
	"strconv"
)

type ShenYuConsulClient struct {
	Ccp          *ConsulClientParam //ConsulClientParam
	ConsulClient *api.Client        //consulClient
}

type ConsulClientParam struct {
	host  string //the customer consul server address
	token string //the customer consul server token
	port  int    //the customer consul server port
}

/**
 * NewClient
 **/
func (scc *ShenYuConsulClient) NewClient(clientParam interface{}) (client interface{}, createResult bool, err error) {
	ccp, ok := clientParam.(*ConsulClientParam)
	if !ok {
		logger.Fatal("init consul client error %+v:", err)
	}
	if len(ccp.host) > 0 && len(ccp.token) > 0 && ccp.port > 0 {
		//use customer param to create client
		config := api.DefaultConfig()
		config.Address = ccp.host + ":" + strconv.Itoa(ccp.port)
		config.Token = ccp.token
		client, err := api.NewClient(config)
		if err == nil {
			return &ShenYuConsulClient{
				Ccp: &ConsulClientParam{
					host:  ccp.host,
					token: ccp.token,
					port:  ccp.port,
				},
				ConsulClient: client,
			}, true, nil
		}
	} else {
		//use default consul client
		config := api.DefaultConfig()
		client, err := api.NewClient(config)
		if err == nil {
			return &ShenYuConsulClient{
				Ccp: &ConsulClientParam{
					host:  ccp.host,
					token: ccp.token,
					port:  ccp.port,
				},
				ConsulClient: client,
			}, true, nil
		}
	}
	return &ShenYuConsulClient{}, false, err
}

/**
 * DeregisterServiceInstance
 **/
func (scc *ShenYuConsulClient) DeregisterServiceInstance(metaData interface{}) (deRegisterResult bool, err error) {
	mdr := scc.checkCommonParam(metaData, err)
	err = scc.ConsulClient.Agent().ServiceDeregister(mdr.AppName)
	if err != nil {
		logger.Fatal("DeregisterServiceInstance failure! ,error is :%+v", err)
	}
	logger.Info("DeregisterServiceInstance,result:%+v\n\n,param:%+v \n\n", true)
	return true, nil
}

/**
 * GetServiceInstanceInfo
 **/
func (scc *ShenYuConsulClient) GetServiceInstanceInfo(metaData interface{}) (instances interface{}, err error) {
	mdr := scc.checkCommonParam(metaData, err)
	catalogService, _, err := scc.ConsulClient.Catalog().Service(mdr.AppName, "", nil)
	if len(catalogService) > 0 && err == nil {
		return catalogService, nil
	}
	return nil, err
}

/**
 * RegisterServiceInstance
 **/
func (scc *ShenYuConsulClient) RegisterServiceInstance(metaData interface{}) (registerResult bool, err error) {
	mdr := scc.checkCommonParam(metaData, err)
	port, err := strconv.Atoi(mdr.Port)
	metaDataStringJson, _ := json.Marshal(metaData)

	//Integrate with MetaDataRegister
	registration := &api.AgentServiceRegistration{
		ID:        mdr.AppName,
		Name:      mdr.AppName,
		Port:      port,
		Address:   mdr.Host,
		Namespace: mdr.ContextPath,
		Meta:      map[string]string{"uriMetadata": string(metaDataStringJson)},
	}

	//server checker
	check := &api.AgentServiceCheck{
		Timeout:                        constants.DEFAULT_CONSUL_CHECK_TIMEOUT,
		Interval:                       constants.DEFAULT_CONSUL_CHECK_INTERVAL,
		DeregisterCriticalServiceAfter: constants.DEFAULT_CONSUL_CHECK_DEREGISTER,
		HTTP:                           fmt.Sprintf("%s://%s:%d/actuator/health", mdr.RPCType, registration.Address, registration.Port),
	}
	registration.Check = check

	//register
	err = scc.ConsulClient.Agent().ServiceRegister(registration)
	if err != nil {
		logger.Fatal("RegisterServiceInstance failure! ,error is :%+v", err)
	}
	logger.Info("RegisterServiceInstance,result:%+v\n\n,param:%+v \n\n", true)
	return true, nil
}

/**
 * check common MetaDataRegister
 **/
func (scc *ShenYuConsulClient) checkCommonParam(metaData interface{}, err error) *model.MetaDataRegister {
	mdr, ok := metaData.(*model.MetaDataRegister)
	if !ok {
		logger.Fatal("get zk client metaData error %+v:", err)
	}
	return mdr
}
