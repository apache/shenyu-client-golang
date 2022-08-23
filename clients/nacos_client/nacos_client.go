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
	"encoding/json"
	"github.com/apache/shenyu-client-golang/common/constants"
	"github.com/apache/shenyu-client-golang/common/utils"
	"github.com/apache/shenyu-client-golang/model"
	oriNc "github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/wonderivan/logger"
	"net"
	"net/url"
	"strconv"
)

/**
 * ShenYuNacosClient
 **/
type ShenYuNacosClient struct {
	NamingClient naming_client.INamingClient
	ConfigClient config_client.IConfigClient
	Ncp *NacosClientParam
	SimpleQueue *utils.SimpleQueue
}

/**
 * NacosClientParam
 **/
type NacosClientParam struct {
	ServerList  []string
	NamespaceId string // the namespaceId of Nacos.When namespace is public, fill in the blank string here  require user provide.
	UserName    string // nacos loginName
	Password    string // nacos loginPwd
	GroupName   string // options
}

/**
 * create nacos client
 **/
func (nc *ShenYuNacosClient) NewClient(clientParam interface{}) (client interface{}, createResult bool, err error) {
	ncp, ok := clientParam.(*NacosClientParam)
	if !ok {
		logger.Fatal("init nacos client error %+v:", err)
	}
	if len(ncp.ServerList) == 0{
		logger.Fatal("The clientParam ServerList must not nil!")
	}
	namingClient,configClient, err := ncp.initNacosClient()
	if err != nil {
		logger.Fatal("init nacos client error %+v:", err)
	}

	if ncp.GroupName == "" {
		ncp.GroupName = constants.DEFAULT_NACOS_GROUP_NAME
	}
	return &ShenYuNacosClient{
		Ncp: &NacosClientParam{
             ServerList: ncp.ServerList,
             UserName: ncp.UserName,
             Password: ncp.Password,
             GroupName: ncp.GroupName,
		},
		NamingClient: namingClient,
        ConfigClient: configClient,
        SimpleQueue: new(utils.SimpleQueue),
	}, true, nil
}

/**
 * use NacosClientParam to init client
 **/
func (ncp *NacosClientParam) initNacosClient() (namingClient naming_client.INamingClient,configClient config_client.IConfigClient, err error) {
	var sc []constant.ServerConfig
	for _,v := range ncp.ServerList{
		u, _ := url.Parse(v)
		host, portStr, _ := net.SplitHostPort(u.Host)
		//host = fmt.Sprintf("%s://%s",u.Scheme,host)
		var port,_ = strconv.ParseUint(portStr,10,64)
		sc = append(sc, *constant.NewServerConfig(host, port))
	}

	//init ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithUsername(ncp.UserName),
		constant.WithPassword(ncp.Password),
		constant.WithNamespaceId(ncp.NamespaceId),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
	)

	namingClient, err = oriNc.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil{
		return nil,nil,err
	}
	configClient,err = oriNc.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
		)
	if err != nil{
		return nil,nil,err
	}
	return namingClient,configClient,nil
}

/**
PersistInterface
*/
func (nc *ShenYuNacosClient) PersistInterface(metaData interface{})(registerResult bool, err error){
	var metadata,ok =  metaData.(*model.MetaDataRegister)
	if !ok {
		logger.Fatal("get nacos client metaData error %+v:", err)
	}
	utils.BuildMetadataDto(metadata)
	var contextPath = utils.BuildRealNodeRemovePrefix(metadata.ContextPath, metadata.AppName)
	var metadataStr,_ = json.Marshal(metadata)
	var configName = utils.BuildServiceConfigPath(metadata.RPCType, contextPath)

	nc.SimpleQueue.QueueAdd(string(metadataStr))

	var set,_ = json.Marshal(nc.SimpleQueue.GetAllQueueData())
	var param = vo.ConfigParam{
		DataId: configName,
		Group: nc.Ncp.GroupName,
		Content:string(set),
	}
	publishResult,err := nc.ConfigClient.PublishConfig(param)
	if !publishResult{
		logger.Error("nacos register metadata fail,please check: %+v",err)
		return publishResult,err
	}
	logger.Info("rpcType:%s ->nacos client register success,meta:%s->ruleName:%s",metadata.RPCType,metadataStr,metadata.RuleName)
	return publishResult,nil
}

/**
PersistURI
*/
func (nc *ShenYuNacosClient) PersistURI(uriRegisterData interface{})(registerResult bool, err error){
	uriRegister,ok := uriRegisterData.(*model.URIRegister)
	if !ok {
		logger.Fatal("get nacos client uriregister error %+v:", err)
	}
	//required
	var serviceName = utils.BuildServiceInstancePath(uriRegister.RPCType)
	var contextPath = utils.BuildRealNodeRemovePrefix(uriRegister.ContextPath, uriRegister.AppName)
	port, _ := strconv.ParseUint(uriRegister.Port,10,64)
	uriRegString, _ := json.Marshal(uriRegister)
    var metaData =  map[string]string{constants.CONTEXT_PATH: contextPath,constants.URI_META_DATA:string(uriRegString)}
	var param  =  vo.RegisterInstanceParam{
		ServiceName: serviceName,
		Weight: 1,
		Enable: true,
        Ephemeral: true,
        Ip: uriRegister.Host,
        Port: port,
        Metadata:metaData,
        GroupName: nc.Ncp.GroupName,
	}
	registerResult, err = nc.NamingClient.RegisterInstance(param)
	if err != nil {
		logger.Error("RegisterServiceInstance failure! ,error is :%+v", err)
		return false, err
	}
	logger.Info("RegisterServiceInstance,result:%+v\n\n,param:%+v \n\n", registerResult, param)
	return registerResult, nil
}

/**
Close
*/
func (nc *ShenYuNacosClient) Close() {

}
