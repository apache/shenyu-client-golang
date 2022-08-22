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
	"github.com/apache/shenyu-client-golang/model"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var (
	logger = logrus.New()
)

/**
 * ShenYuEtcdClient
 **/
type ShenYuEtcdClient struct {
	Ecp        *EtcdClientParam //EtcdClientParam
	EtcdClient *clientv3.Client //EtcdClient
	// GlobalLease clientv3.LeaseID //global lease
}

/**
 * EtcdClientParam
 **/
type EtcdClientParam struct {
	EtcdServers []string //the customer etcd server address
	UserName    string   //the customer etcd server userName
	Password    string   //the customer etcd server pwd
	TTL         int64    //the customer etcd key rent
}

/**
 * init NewClient
 **/
func (sec *ShenYuEtcdClient) NewClient(clientParam interface{}) (client interface{}, createResult bool, err error) {
	ecp, ok := clientParam.(*EtcdClientParam)
	if !ok {
		logger.Fatalf("The clientParam  must not nil!")
	}
	if len(ecp.EtcdServers) > 0 {
		//use customer param to create client
		client, err := clientv3.New(clientv3.Config{
			Endpoints:   ecp.EtcdServers,
			DialTimeout: constants.DEFAULT_ETCD_TIMEOUT * time.Second,
			Username:    ecp.UserName,
			Password:    ecp.Password,
		})
		if err == nil {
			logger.Infof("Create customer etcd client success!")
			return &ShenYuEtcdClient{
				Ecp: &EtcdClientParam{
					EtcdServers: ecp.EtcdServers,
					UserName:    ecp.UserName,
					Password:    ecp.Password,
					TTL:         ecp.TTL,
				},
				EtcdClient: client,
			}, true, nil
		}
		logger.Fatalf("init etcd client error %v:", err)
	}
	return
}

/**
DeregisterServiceInstance
*/
func (sec *ShenYuEtcdClient) DeregisterServiceInstance(metaData interface{}) (deRegisterResult bool, err error) {
	mdr, ok := metaData.(*model.MetaDataRegister)
	if !ok {
		logger.Fatalf("get etcd client metaData error %v:", err)
	}
	key := mdr.AppName
	ctx, cancel := context.WithTimeout(context.Background(), constants.DEFAULT_ETCD_TIMEOUT*time.Second)
	defer cancel()
	_, err = sec.EtcdClient.Delete(ctx, key)
	if err != nil {
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
	ctx, cancel := context.WithTimeout(context.Background(), constants.DEFAULT_ETCD_TIMEOUT*time.Second)
	defer cancel()
	resp, err := sec.EtcdClient.Get(ctx, key)
	if err != nil {
		logger.Error("etcd Get data failure, err:", err)
		return nil, err
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
	ctx, cancel := context.WithTimeout(context.Background(), constants.DEFAULT_ETCD_TIMEOUT*time.Second)
	defer cancel()
	_, err = sec.EtcdClient.Put(ctx, key, string(data))
	if err != nil {
		logger.Errorf("RegisterServiceInstance failure! ,error is :%v", err)
		return false, err
	}
	logger.Infof("RegisterServiceInstance,result:%v", true)
	return true, nil
}

/**
 * check common MetaDataRegister
 **/
func (sec *ShenYuEtcdClient) checkCommonParam(metaData interface{}, err error) *model.MetaDataRegister {
	mdr, ok := metaData.(*model.MetaDataRegister)
	if !ok {
		logger.Fatalf("get etcd client metaData error %v:", err)
	}
	return mdr
}

/**
 * close etcdClient
 **/
func (sec *ShenYuEtcdClient) Close() {
	sec.EtcdClient.Close()
}
