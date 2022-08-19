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

package zk_client

import (
	"encoding/json"
	"github.com/apache/shenyu-client-golang/common/constants"
	"github.com/apache/shenyu-client-golang/model"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	logger = logrus.New()
)

/**
 * ShenYuZkClient
 **/
type ShenYuZkClient struct {
	ZkClient *zk.Conn       // ZkClient
	Zcp      *ZkClientParam //client param
}

/**
 * ZkClientParam
 **/
type ZkClientParam struct {
	ZkServers []string // ZkServers ex: 127.0.0.1
	ZkRoot    string   // zkClient Root
}

/**
 * init NewClient
 **/
func (zc *ShenYuZkClient) NewClient(clientParam interface{}) (client interface{}, createResult bool, err error) {
	zcp, ok := clientParam.(*ZkClientParam)
	if !ok {
		logger.Fatalf("The clientParam  must not nil!")
	}
	//client = new(ShenYuZkClient)
	if len(zcp.ZkRoot) == 0 {
		logger.Fatalf("The param zkRoot must set a value!")
	}
	conn, _, err := zk.Connect(zcp.ZkServers, time.Duration(constants.DEFAULT_ZOOKEEPER_CLIENT_TIME)*time.Second)
	if err != nil {
		if err := zc.ensureRoot(); err != nil {
			zc.Close()
			return &ShenYuZkClient{}, false, err
		}
	}
	return &ShenYuZkClient{
		Zcp: &ZkClientParam{
			ZkRoot:    zcp.ZkRoot,
			ZkServers: zcp.ZkServers,
		},
		ZkClient: conn,
	}, true, nil
}

/**
 * DeregisterServiceInstance
 **/
func (zc *ShenYuZkClient) DeregisterServiceInstance(metaData interface{}) (deRegisterResult bool, err error) {
	mdr, ok := metaData.(*model.MetaDataRegister)
	if !ok {
		logger.Fatalf("get zk client metaData error %v:", err)
	}
	if err := zc.ensureName(mdr.AppName); err != nil {
		return false, err
	}
	path := zc.Zcp.ZkRoot + "/" + mdr.AppName
	childs, stat, err := zc.ZkClient.Children(path)
	if err != nil {
		return false, err
	}
	for _, child := range childs {
		fullPath := path + "/" + child
		err := zc.ZkClient.Delete(fullPath, stat.Version)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

/**
 * GetServiceInstanceInfo
 **/
func (zc *ShenYuZkClient) GetServiceInstanceInfo(metaData interface{}) (instances interface{}, err error) {
	mdr := zc.checkCommonParam(metaData, err)
	path := zc.Zcp.ZkRoot + "/" + mdr.AppName
	var nodes []*model.MetaDataRegister
	data, _, err := zc.ZkClient.Get(path)
	if err != nil {
		logger.Fatalf("zk Get node failure, err  %v:", err)
	}
	node := new(model.MetaDataRegister)
	err = json.Unmarshal(data, node)
	if err != nil {
		return nil, err
	}
	nodes = append(nodes, node)
	return nodes, nil
}

/**
 * GetEphemeralServiceInstanceInfo
 **/
func (zc *ShenYuZkClient) GetEphemeralServiceInstanceInfo(metaData interface{}) (instances interface{}, err error) {
	mdr := zc.checkCommonParam(metaData, err)
	path := zc.Zcp.ZkRoot + "/" + mdr.AppName
	childs, _, err := zc.ZkClient.Children(path)
	if err != nil {
		if err == zk.ErrNoNode {
			return []*model.MetaDataRegister{}, nil //default return empty MetaDataRegister
		}
		return nil, err
	}
	var nodes []*model.MetaDataRegister
	for _, child := range childs {
		fullPath := path + "/" + child
		data, _, err := zc.ZkClient.Get(fullPath)
		if err != nil {
			if err == zk.ErrNoNode {
				continue
			}
			return nil, err
		}
		node := new(model.MetaDataRegister)
		err = json.Unmarshal(data, node)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}

/**
 * RegisterNodeInstance zk node
 **/
func (zc *ShenYuZkClient) RegisterServiceInstance(metaData interface{}) (registerResult bool, err error) {
	mdr := zc.checkCommonParam(metaData, err)
	err = zc.ensureRoot()
	if err != nil {
		logger.Fatalf("ensureRoot failure, err  %v:", err)
	}
	path := zc.Zcp.ZkRoot + "/" + mdr.AppName
	data, err := json.Marshal(metaData)
	if err != nil {
		return false, err
	}
	_, err = zc.ZkClient.Create(path, data, 0, zk.WorldACL(zk.PermAll))
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
 * check common MetaDataRegister
 **/
func (zc *ShenYuZkClient) checkCommonParam(metaData interface{}, err error) *model.MetaDataRegister {
	mdr, ok := metaData.(*model.MetaDataRegister)
	if !ok {
		logger.Fatalf("get zk client metaData error %v:", err)
	}
	return mdr
}

/**
 * close zkClient
 **/
func (zc *ShenYuZkClient) Close() {
	zc.ZkClient.Close()
}

/**
 * ensure zkRoot avoid create error
 **/
func (zc *ShenYuZkClient) ensureRoot() error {
	exists, _, err := zc.ZkClient.Exists(zc.Zcp.ZkRoot)
	if err != nil {
		return err
	}
	if !exists {
		_, err := zc.ZkClient.Create(zc.Zcp.ZkRoot, []byte(""), 0, zk.WorldACL(zk.PermAll))
		if err != nil && err != zk.ErrNodeExists {
			return err
		}
	}
	return nil
}

/**
 * ensure zkRoot&nodeName
 **/
func (zc *ShenYuZkClient) ensureName(name string) error {
	path := zc.Zcp.ZkRoot + "/" + name
	logger.Infof("ensureName check, path is  %v: ->", path)
	exists, _, err := zc.ZkClient.Exists(path) //avoid create error
	logger.Infof("ensureName check result is  %v: ->", exists)
	if err != nil {
		return err
	}
	if !exists {
		_, err = zc.ZkClient.Create(path, []byte(""), 0, zk.WorldACL(zk.PermAll))
		if err != nil && err == zk.ErrNodeExists {
			logger.Infof("ensureName inner create success")
			return nil
		}
	}
	return nil
}
