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
	"fmt"
	"github.com/apache/shenyu-client-golang/common/constants"
	"github.com/apache/shenyu-client-golang/common/utils"
	"github.com/apache/shenyu-client-golang/model"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/sirupsen/logrus"
	"strings"
	"sync"
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
	NodeDataMap *sync.Map
	MasterWatch <- chan zk.Event
}

/**
 * ZkClientParam
 **/
type ZkClientParam struct {
	ServerList []string //  ex: 127.0.0.1
    Digest  string //zk user
}


/**
 * init NewClient
 **/
func (zc *ShenYuZkClient) NewClient(clientParam interface{}) (client interface{}, createResult bool, err error) {
	zcp, ok := clientParam.(*ZkClientParam)
	if !ok {
		logger.Fatalf("The clientParam  must not nil!")
	}
	//event
	eventCallbackOption := zk.WithEventCallback(callback)
	conn, watchEventChan, err := zk.Connect(zcp.ServerList, time.Duration(constants.DEFAULT_ZOOKEEPER_CLIENT_TIME)*time.Second,eventCallbackOption)
	if err != nil {
		logger.Errorf("zk connect fail %+v",err)
		return &ShenYuZkClient{}, false, err
	}
	if zcp.Digest != ""{
		err = conn.AddAuth("digest",[]byte(zcp.Digest))
		if err != nil{
			logger.Errorf("zk digest fail %+v",err)
			return &ShenYuZkClient{}, false, err
		}
	}
	return &ShenYuZkClient{
		Zcp: &ZkClientParam{
			ServerList: zcp.ServerList,
			Digest: zcp.Digest,
		},
		ZkClient: conn,
		NodeDataMap: new(sync.Map),
		MasterWatch: watchEventChan,
	}, true, nil
}

/**
PersistInterface
*/
func (zc *ShenYuZkClient) PersistInterface(metaData interface{})(registerResult bool, err error){
	var metadata,ok =  metaData.(*model.MetaDataRegister)
	if !ok {
		logger.Fatalf("get zookeeper client metaData error %+v:", err)
	}
	utils.BuildMetadataDto(metadata)
	var contextPath = utils.BuildRealNodeRemovePrefix(metadata.ContextPath, metadata.AppName)
	var metadataNodeName = utils.BuildMetadataNodeName(*metadata)
	var metaDataPath = utils.BuildMetaDataParentPath(metadata.RPCType, contextPath)
	var realNode = utils.BuildRealNode(metaDataPath, metadataNodeName)

	// create node with mode per
	err = zc.createNodeWithParent(metaDataPath, nil, zk.WorldACL(zk.PermAll), 0)
    if err != nil{
    	return false, err
	}
	var metadataStr,_ = json.Marshal(metadata)
	err = zc.createNodeOrUpdate(realNode, metadataStr,zk.WorldACL(zk.PermAll), 0)
	if err != nil{
		return false, err
	}
	logger.Infof("%s zookeeper client register success: %s",metadata.RPCType,metadataStr)
	return true,nil
}

/**
PersistURI
*/
func (zc *ShenYuZkClient) PersistURI(uriRegisterData interface{})(registerResult bool, err error){
	uriRegister,ok := uriRegisterData.(*model.URIRegister)
	if !ok {
		logger.Fatalf("get zookeeper client uriregister error %+v:", err)
	}
	var contextPath = utils.BuildRealNodeRemovePrefix(uriRegister.ContextPath,uriRegister.AppName)
	var uriNodeName = utils.BuildURINodeName(*uriRegister)
	var uriPath = utils.BuildURIParentPath(uriRegister.RPCType, contextPath)
	var realNode = utils.BuildRealNode(uriPath, uriNodeName)
	err = zc.createNodeWithParent(uriPath, nil, zk.WorldACL(zk.PermAll), 0)
    if err != nil{
    	return false, err
	}
	var nodeData,_ = json.Marshal(uriRegister)
	//set dic
	zc.NodeDataMap.Store(realNode,nodeData)
	//createMode FlagEphemeral=1 if session DisConnect will delete
	err = zc.createNodeOrUpdate(realNode,nodeData,zk.WorldACL(zk.PermAll),zk.FlagEphemeral)
	if err != nil{
		return false, err
	}
	return true, nil
}

/**
Close
*/
func (zc *ShenYuZkClient) Close(){
	zc.ZkClient.Close()
}

/*
 global zk event callback
 */
func callback(event zk.Event) {
	//masterWatch <- event
	//fmt.Println("###########################")
	//fmt.Println("path: ", event.Path)
	//fmt.Println("type: ", event.Type.String())
	//fmt.Println("state: ", event.State.String())
	//fmt.Println("---------------------------")
}

/**
WatchEventHandler
**/
func(zc *ShenYuZkClient) WatchEventHandler(){
	for {
		for event := range zc.MasterWatch{
			if event.State == zk.StateConnected || event.State == zk.StateConnectedReadOnly{
				if zc.NodeDataMap != nil {
					zc.NodeDataMap.Range(func(k ,v interface{}) bool{
						key, _ := k.(string)
						val, _ := v.([]byte)
						logger.Infof("watch change %s",key)
						var exists,_,_ =zc.ZkClient.Exists(key)
						if !exists {
							err := zc.createNodeOrUpdate(key, val, zk.WorldACL(zk.PermAll), zk.FlagEphemeral)
							if err != nil{
								logger.Errorf("watch eventHandler CreateNodeOrUpdate err:%+v",err)
							}
						}
						return true
					})
				}
			}
		}
	}
}

/*
 createNodeWithParent
 */
func(zc *ShenYuZkClient) createNodeWithParent(path string,data []byte, acl []zk.ACL,createMode int32) error {
	path = getZooKeeperPath(path)
	if path != constants.PathSeparator {
		path = utils.RemoveSuffix(utils.RemovePrefix(path))
	}
	tempPath := utils.RepairData(path)
	var paths = strings.Split(path,constants.PathSeparator)
	var cur = ""
	var err error
	for _,item := range paths {
	  if item == ""{
	  	continue
	  }
	  cur = fmt.Sprintf("%s%s%s",cur,constants.PathSeparator,item)
	  var exist,_,_ = zc.ZkClient.Exists(cur)
	  if exist {
			continue
		}

	  if cur == tempPath {
	    _,err =	zc.ZkClient.Create(cur,data,createMode,acl)
		} else {
		 _,err = zc.ZkClient.Create(cur,nil,createMode,acl)
		}
	}
	return err
}

/**
  create node or update nodedata
 */
func(zc *ShenYuZkClient) createNodeOrUpdate(path string,data []byte, acl []zk.ACL,createMode int32) error {
	path = getZooKeeperPath(path)
	if path != constants.PathSeparator {
		path = utils.RemoveSuffix(utils.RemovePrefix(path))
	}
	tempPath := utils.RepairData(path)
	var paths = strings.Split(path,constants.PathSeparator)
	var cur = ""
	var err error
	for _,item := range paths {
		if item == ""{
			continue
		}
		cur = fmt.Sprintf("%s%s%s",cur,constants.PathSeparator,item)
		var exist,_,_ = zc.ZkClient.Exists(cur)
		if exist {
			if cur == tempPath{
				_,err = zc.ZkClient.Set(path,data,-1)
			}
			continue
		}

		if cur == tempPath {
			_,err =	zc.ZkClient.Create(cur,data,createMode,acl)
		} else {
			_,err = zc.ZkClient.Create(cur,nil,0,acl)
		}
	}
	return err
}

/*
 check path and return correctPath
 */
func getZooKeeperPath(path string) string {
	if path == "" || path == constants.PathSeparator{
		return constants.PathSeparator
	}
	if !strings.HasPrefix(path,constants.PathSeparator){
		path = fmt.Sprintf("%s%s","/",path)
	}
    if strings.HasSuffix(path,constants.PathSeparator){
    	path = path[0:len([]rune(path)) - 1]
	}
   return  path
}


