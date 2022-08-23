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
	"github.com/apache/shenyu-client-golang/common/constants"
	"github.com/apache/shenyu-client-golang/common/utils"
	"github.com/apache/shenyu-client-golang/model"
	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"strconv"
)

var (
	logger = logrus.New()
)

/**
 * ShenYuConsulClient
 **/
type ShenYuConsulClient struct {
	Ccp          *ConsulClientParam //ConsulClientParam
	ConsulClient *api.Client        //consulClient
}

/**
 * ConsulClientParam
 **/
type ConsulClientParam struct {
	Id string // namespaceId
	//NameSpace string //namespace
	Token string
	ServerList []string //ip+port
	Tags []string
	EnableTagOverride bool
}

/**
 * NewClient
 **/
func (scc *ShenYuConsulClient) NewClient(clientParam interface{}) (client interface{}, createResult bool, err error) {
	ccp, ok := clientParam.(*ConsulClientParam)
	if !ok {
		logger.Fatalf("The clientParam  must not nil!")
	}
	if len(ccp.ServerList) == 0{
		logger.Fatalf("The clientParam ServerList must not nil!")
	}
	//use customer param to create client
	config := api.DefaultConfig()
	config.Address = ccp.ServerList[0]
	config.Token = ccp.Token
	//config.Namespace = ccp.NameSpace
	client, err = api.NewClient(config)
	if err == nil {
		logger.Infof("Create customer consul client success!")
		return &ShenYuConsulClient{
			Ccp: &ConsulClientParam{
				 Id: ccp.Id,
				// NameSpace: ccp.NameSpace,
				 ServerList: ccp.ServerList,
				 Tags: ccp.Tags,
				 EnableTagOverride: ccp.EnableTagOverride,
			},
			ConsulClient: client.(*api.Client),
		}, true, nil
	}
	logger.Errorf("init consul client error %+v:", err)
	return nil,false,err
}

/**
PersistInterface
 */
func (scc *ShenYuConsulClient) PersistInterface(metaData interface{})(registerResult bool, err error){
	var metadata,ok =  metaData.(*model.MetaDataRegister)
	if !ok {
		logger.Fatalf("get consul client metaData error %+v:", err)
	}
	utils.BuildMetadataDto(metadata)
	var contextPath = utils.BuildRealNodeRemovePrefix(metadata.ContextPath, metadata.AppName)
	var metadataNodeName = utils.BuildMetadataNodeName(*metadata)
	var metaDataPath = utils.BuildMetaDataParentPath(metadata.RPCType, contextPath)
	var realNode = utils.BuildRealNode(metaDataPath, metadataNodeName)
	realNode = utils.RemovePrefix(realNode)//remove prefix /
	var metadataStr,_ = json.Marshal(metaData)
	var putPair = &api.KVPair{
		Key: realNode,
		Value: []byte(metadataStr),
		Flags: 0,
	}
    _,err =scc.ConsulClient.KV().Put(putPair,nil)
    if err != nil{
		logger.Errorf("Consul client register failure! ,error is :%+v", err)
		return false,err
	}
	logger.Infof("%s Consul client register success: %s",metadata.RPCType,metadataStr)
	return true,nil
}

/**
PersistURI
 */
func (scc *ShenYuConsulClient) PersistURI(uriRegisterData interface{})(registerResult bool, err error){
	uriRegister,ok := uriRegisterData.(*model.URIRegister)
	if !ok {
		logger.Fatalf("get consul client uriregister error %+v:", err)
	}
	port, _ := strconv.Atoi(uriRegister.Port)
	uriRegString, _ := json.Marshal(uriRegister)

	//Integrate with MetaDataRegister
	registration := &api.AgentServiceRegistration{
		ID:        scc.Ccp.Id,
		//Namespace: scc.Ccp.NameSpace,
		Name:      uriRegister.AppName,
		Tags:      scc.Ccp.Tags,
		Port:      port,
		Address:   uriRegister.Host,
		EnableTagOverride: scc.Ccp.EnableTagOverride,
		Meta:      map[string]string{constants.UriType: string(uriRegString)},
	}

	////server checker
	//check := &api.AgentServiceCheck{
	//	Timeout:                        constants.DEFAULT_CONSUL_CHECK_TIMEOUT,
	//	Interval:                       constants.DEFAULT_CONSUL_CHECK_INTERVAL,
	//	DeregisterCriticalServiceAfter: constants.DEFAULT_CONSUL_CHECK_DEREGISTER,
	//	HTTP:                           fmt.Sprintf("%s://%s:%d/actuator/health", uriRegister.RPCType, registration.Address, registration.Port),
	//}
	//registration.Check = check

	//register
	err = scc.ConsulClient.Agent().ServiceRegister(registration)
	if err != nil {
		logger.Fatalf("RegisterServiceInstance failure! ,error is :%+v", err)
	}
	logger.Infof("RegisterServiceInstance,result:%+v", true)
	return true, nil
}

/**
Close
 */
func (scc *ShenYuConsulClient) Close(){
  var err=	scc.ConsulClient.Agent().ServiceDeregister(scc.Ccp.Id)
  if err != nil{
  	logger.Errorf("close consul fail:%=v",err)
  }
}

