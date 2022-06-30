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
	"github.com/apache/incubator-shenyu-client-golang/model"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/wonderivan/logger"
	"time"
)

/**
 * ServiceNode
 **/
type ServiceNode struct {
	RootName string                  `json:"rootName"` // the register zk rootName  require user provide
	MetaData *model.MetaDataRegister `json:"metaData"` // the register metaData  require user provide
}

/**
 * ZkClient
 **/
type ZkClient struct {
	zkServers []string // zkServers ex: 127.0.0.1
	zkRoot    string   // zkClient Root
	zkClient  *zk.Conn // zkClient
}

/**
 * init NewClient
 **/
func NewClient(zkServers []string, zkRoot string, timeout int) (*ZkClient, error) {
	client := new(ZkClient)
	client.zkServers = zkServers
	if len(zkRoot) == 0 {
		logger.Fatal("The param zkRoot must set a value!")
	}
	client.zkRoot = zkRoot
	conn, _, err := zk.Connect(zkServers, time.Duration(timeout)*time.Second)
	if err != nil {
		return nil, err
	}
	client.zkClient = conn
	if err := client.ensureRoot(); err != nil {
		client.Close()
		return nil, err
	}
	return client, nil
}

/**
 * close zkClient
 **/
func (s *ZkClient) Close() {
	s.zkClient.Close()
}

/**
 * ensure zkRoot avoid create error
 **/
func (s *ZkClient) ensureRoot() error {
	exists, _, err := s.zkClient.Exists(s.zkRoot)
	if err != nil {
		return err
	}
	if !exists {
		_, err := s.zkClient.Create(s.zkRoot, []byte(""), 0, zk.WorldACL(zk.PermAll))
		if err != nil && err != zk.ErrNodeExists {
			return err
		}
	}
	return nil
}

/**
 * RegisterNodeInstance zk node
 **/
func (s *ZkClient) RegisterNodeInstance(metaData *model.MetaDataRegister) error {
	if err := s.ensureName(metaData.AppName); err != nil {
		return err
	}
	path := s.zkRoot + "/" + metaData.AppName + "/n"
	data, err := json.Marshal(metaData)
	if err != nil {
		return err
	}
	_, err = s.zkClient.CreateProtectedEphemeralSequential(path, data, zk.WorldACL(zk.PermAll))
	if err != nil {
		return err
	}
	return nil
}

/**
 * DeleteNodeInstance
 **/
func (s *ZkClient) DeleteNodeInstance(metaData *model.MetaDataRegister) error {
	if err := s.ensureName(metaData.AppName); err != nil {
		return err
	}
	path := s.zkRoot + "/" + metaData.AppName
	childs, stat, err := s.zkClient.Children(path)
	if err != nil {
		return err
	}
	for _, child := range childs {
		fullPath := path + "/" + child
		err := s.zkClient.Delete(fullPath, stat.Version)
		if err != nil {
			return err
		}
	}
	return nil
}

/**
 *  * ensure zkRoot&nodeName
 **/
func (s *ZkClient) ensureName(name string) error {
	path := s.zkRoot + "/" + name
	exists, _, err := s.zkClient.Exists(path) //avoid create error
	if err != nil {
		return err
	}
	if !exists {
		_, err := s.zkClient.Create(path, []byte(""), 0, zk.WorldACL(zk.PermAll))
		if err != nil && err != zk.ErrNodeExists {
			return err
		}
	}
	return nil
}

/**
 * get zk nodes metaData
 **/
func (s *ZkClient) GetNodesInfo(name string) ([]*model.MetaDataRegister, error) {
	path := s.zkRoot + "/" + name
	childs, _, err := s.zkClient.Children(path)
	if err != nil {
		if err == zk.ErrNoNode {
			return []*model.MetaDataRegister{}, nil //default return empty MetaDataRegister
		}
		return nil, err
	}
	var nodes []*model.MetaDataRegister
	for _, child := range childs {
		fullPath := path + "/" + child
		data, _, err := s.zkClient.Get(fullPath)
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
