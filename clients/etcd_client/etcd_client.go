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
	EtcdServers  []string //the customer etcd server address
	UserName string //the customer etcd server userName
	Password  string    //the customer etcd server pwd
	TimeOut int64   //the customer etcd server timeout
	TTL int64 //the customer etcd key rent
}

func (sec *ShenYuEtcdClient) NewClient(clientParam interface{}) (client interface{}, createResult bool, err error) {
	ecp, ok := clientParam.(*EtcdClientParam)
	if !ok {
		logger.Fatal("init etcd client error %+v:", err)
	}
	if len(ecp.EtcdServers) > 0 {
		if ecp.TimeOut == 0 {
			ecp.TimeOut = 5
		}
		//use customer param to create client
		client, err := clientv3.New(clientv3.Config{
			Endpoints:   ecp.EtcdServers,
			DialTimeout: time.Duration(ecp.TimeOut) * time.Second,
			Username:    ecp.UserName,
			Password:    ecp.Password,
		})
		if err == nil {
			logger.Info("Create customer etcd client success!")
			////wheather gen leaseId is with by user
			//leaseId := sec.getLocalLeaseId(client,ccp)
			////keep live
			//go func() {
			//	sec.keepLocalAlive(client,leaseId)
			//}()
			return &ShenYuEtcdClient{
				Ecp: &EtcdClientParam{
					EtcdServers: ecp.EtcdServers,
					UserName: ecp.UserName,
					Password: ecp.Password,
					TimeOut: ecp.TimeOut,
					TTL: ecp.TTL,
				},
				EtcdClient: client,
				//GlobalLease: leaseId,
			}, true, nil
		}
	}
	return &ShenYuEtcdClient{}, false, err
}

/**
DeregisterServiceInstance
 */
func (sec *ShenYuEtcdClient) DeregisterServiceInstance(metaData interface{}) (deRegisterResult bool, err error) {
	mdr, ok := metaData.(*model.MetaDataRegister)
	if !ok {
		logger.Fatal("get etcd client metaData error %+v:", err)
	}
	key :=  mdr.AppName
	_,err = sec.EtcdClient.Delete(context.TODO(),key)
	// revoke by LeaseId
	//_,err = sec.EtcdClient.Revoke(context.TODO(),sec.GlobalLease)
    if err != nil{
    	return false, err
	}
	return true, nil
}

/**
* RegisterServiceInstance
 */
func (sec *ShenYuEtcdClient) GetServiceInstanceInfo(metaData interface{}) (instances interface{}, err error) {
	mdr := sec.checkCommonParam(metaData, err)
	key := mdr.AppName
	var nodes []*model.MetaDataRegister
	resp,err := sec.EtcdClient.Get(context.TODO(),key)
	if err != nil {
		logger.Error("etcd Get data failure, err:", err)
	}
	node := new(model.MetaDataRegister)
	err = json.Unmarshal(resp.Kvs[0].Value, node)
	if err != nil {
		return nil, err
	}
	nodes = append(nodes, node)
	return nodes, nil
}

/**
* RegisterServiceInstance
 **/
func (sec *ShenYuEtcdClient) RegisterServiceInstance(metaData interface{}) (registerResult bool, err error) {
	mdr := sec.checkCommonParam(metaData, err)
	data, _ := json.Marshal(metaData)
	if err != nil {
		return false, err
	}
	key := mdr.AppName
	//register  put with lease
	//_,err = sec.EtcdClient.Put(context.TODO(), key, string(data), clientv3.WithLease(sec.GlobalLease))
	_,err = sec.EtcdClient.Put(context.TODO(), key, string(data))
	if err != nil {
		logger.Fatal("RegisterServiceInstance failure! ,error is :%+v", err)
	}
	logger.Info("RegisterServiceInstance,result:%+v", true)
	return true, nil
}

/**
* GenLeaseId //get etcd  grant leaseId
**/
func (sec *ShenYuEtcdClient) GenLeaseId() clientv3.LeaseID {
	//grant lease
	lease, err := sec.EtcdClient.Grant(context.TODO(), sec.Ecp.TTL)
	if err != nil {
		logger.Error("Grant lease failed: %v\n", err)
	}
	return lease.ID
}


/**
* KeepAlive
 */
func (sec *ShenYuEtcdClient) KeepAlive()  {
	//keep alive
	kaCh, err := sec.EtcdClient.KeepAlive(context.Background(), sec.GlobalLease)
	if err != nil {
		logger.Error("Keep alive with lease[%s] failed: %v\n",sec.GlobalLease, err)
	}
	for {
		kaResp := <-kaCh
		logger.Info("ttl: ", kaResp.TTL)
	}
}

/**
 * check common MetaDataRegister
 **/
func (sec *ShenYuEtcdClient) checkCommonParam(metaData interface{}, err error) *model.MetaDataRegister {
	mdr, ok := metaData.(*model.MetaDataRegister)
	if !ok {
		logger.Fatal("get etcd client metaData error %+v:", err)
	}
	return mdr
}


/**
 * close etcdClient
 **/
func (sec *ShenYuEtcdClient) Close()  {
	sec.EtcdClient.Close()
}
