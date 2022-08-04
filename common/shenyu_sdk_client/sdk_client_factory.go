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

package shenyu_sdk_client

import (
	"github.com/apache/shenyu-client-golang/clients/consul_client"
	"github.com/apache/shenyu-client-golang/clients/etcd_client"
	"github.com/apache/shenyu-client-golang/clients/nacos_client"
	"github.com/apache/shenyu-client-golang/clients/zk_client"
	"github.com/apache/shenyu-client-golang/common/constants"
)

/**
 * Get client by clientName
 **/
func GetFactoryClient(clientName string) SdkClient {
	switch clientName {
	case constants.NACOS_CLIENT:
		return &nacos_client.ShenYuNacosClient{}
	case constants.ZOOKEEPER_CLIENT:
		return &zk_client.ShenYuZkClient{}
	case constants.CONSUL_CLIENT:
		return &consul_client.ShenYuConsulClient{}
	case constants.ETCD_CLIENT:
		return &etcd_client.ShenYuEtcdClient{}
	default:
		return nil
	}
}
