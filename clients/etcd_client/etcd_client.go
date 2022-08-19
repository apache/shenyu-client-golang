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

package etcd_client

import (
	"context"
	"encoding/json"
	"github.com/apache/shenyu-client-golang/common/constants"
	"github.com/apache/shenyu-client-golang/common/utils"
	"github.com/apache/shenyu-client-golang/model"
	"github.com/wonderivan/logger"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

/**
 * ShenYuEtcdClient
 **/
type ShenYuEtcdClient struct {
	Ecp *EtcdClientParam //EtcdClientParam
    EtcdClient *clientv3.Client //EtcdClient
    GlobalLease clientv3.LeaseID //global lease
}

/**
 * EtcdClientParam
 **/
type EtcdClientParam struct {
	ServerList  []string //the customer etcd server address
	UserName string //the customer etcd server userName
	Password  string    //the customer etcd server pwd
	TTL int64 //the customer etcd key rent
}


/**
 * init NewClient
 **/
func (sec *ShenYuEtcdClient) NewClient(clientParam interface{}) (client interface{}, createResult bool, err error) {
	ecp, ok := clientParam.(*EtcdClientParam)
	if !ok {
		logger.Fatal("The clientParam  must not nil!")
	}
	if len(ecp.ServerList) > 0 {
		//use customer param to create client
		client, err := clientv3.New(clientv3.Config{
			Endpoints:   ecp.ServerList,
			DialTimeout: constants.DEFAULT_ETCD_CLIENT_TIMEOUT * time.Second,
			Username:    ecp.UserName,
			Password:    ecp.Password,
		})
		if err == nil {
			logger.Info("Create customer etcd client success!")
			//get leaseId
			var leaseId,err = genLeaseId(client,ecp.TTL)
			if err != nil{
				return nil,false,err
			}
			//rent leaseId
			go func() {
				keepAlive(client,leaseId)
			}()
			return &ShenYuEtcdClient{
				Ecp: &EtcdClientParam{
					ServerList: ecp.ServerList,
					UserName: ecp.UserName,
					Password: ecp.Password,
					TTL: ecp.TTL,
				},
				EtcdClient: client,
				GlobalLease: leaseId,
			}, true, nil
		}
		logger.Fatal("init etcd client error %+v:", err)
	}
	return
}

/**
PersistInterface
*/
func (sec *ShenYuEtcdClient) PersistInterface(metaData interface{})(registerResult bool, err error){
	var metadata,ok =  metaData.(*model.MetaDataRegister)
	if !ok {
		logger.Fatal("get etcd client metaData error %+v:", err)
	}
	var contextPath = utils.BuildRealNode(metadata.ContextPath, metadata.AppName)
	var metadataNodeName = utils.BuildMetadataNodeName(*metadata)
	var metaDataPath = utils.BuildMetaDataParentPath(metadata.RPCType, contextPath)
	var realNode = utils.BuildRealNode(metaDataPath, metadataNodeName)
	var metadataStr,_ = json.Marshal(metadata)
	err = sec.putEphemeral(realNode,metadataStr)
	if err != nil{
		return false,err
	}
	logger.Info("%s etcd client register success: %s",metadata.RPCType,metadataStr)
	return true,nil
}

/**
PersistURI
*/
func (sec *ShenYuEtcdClient) PersistURI(uriRegisterData interface{})(registerResult bool, err error){
	uriRegister,ok := uriRegisterData.(*model.URIRegister)
	if !ok {
		logger.Fatal("get etcd client uriregister error %+v:", err)
	}
	var contextPath = utils.BuildRealNode(uriRegister.ContextPath, uriRegister.AppName)
	var uriNodeName = utils.BuildURINodeName(*uriRegister)
	var uriPath = utils.BuildURIParentPath(uriRegister.RPCType, contextPath)
	var realNode = utils.BuildRealNode(uriPath, uriNodeName)
	var nodeData,_ = json.Marshal(uriRegister)
	err = sec.putEphemeral(realNode, nodeData)
    if err != nil{
    	return false, err
	}
	logger.Info("RegisterServiceInstance,result:%+v", true)
	return true, nil
}

/**
Close
*/
func (sec *ShenYuEtcdClient) Close(){
	sec.EtcdClient.Close()
}

/**
 put data with leaseID,so you should generate leaseId use GenLeaseId()
 */
func (sec *ShenYuEtcdClient) putEphemeral(key string,val []byte) error{
	ctx, cancel := context.WithTimeout(context.Background(),constants.DEFAULT_ETCD_CLIENT_TIMEOUT* time.Second)
	defer cancel()
	var _,err = sec.EtcdClient.KV.Put(ctx,key,string(val),clientv3.WithLease(sec.GlobalLease))
	return err
}

/**
* genLeaseId //get etcd  grant leaseId
**/
func  genLeaseId(client *clientv3.Client,ttl int64) (clientv3.LeaseID,error) {
	ctx, cancel := context.WithTimeout(context.Background(),constants.DEFAULT_ETCD_CLIENT_TIMEOUT* time.Second)
	defer cancel()
	//grant lease
	lease, err := client.Grant(ctx, ttl)
	if err != nil {
		logger.Error("Grant lease failed: %v\n", err)
		return 0,err
	}
	return lease.ID,nil
}


/**
* KkeepAlive
 */
func  keepAlive(client *clientv3.Client,leaseId clientv3.LeaseID)  {
	//keep alive
	kaCh, err := client.KeepAlive(context.Background(), leaseId)
	if err != nil {
		logger.Error("Keep alive with lease[%s] failed: %v\n",leaseId, err)
	}
	for {
		<-kaCh
	}
}


