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

package constants

/**
 * The sdk const
 **/
const (
	DEFAULT_SHENYU_ADMIN_URL        = "http://127.0.0.1:9095"
	DEFAULT_ADMIN_PASSWORD          = "123456"
	DEFAULT_ADMIN_ACCOUNT           = "admin"
	DEFAULT_REQUEST_TIME            = 1000
	DEFAULT_ADMIN_SUCCESS           = "success"
	NACOS_CLIENT                    = "Nacos"
	ZOOKEEPER_CLIENT                = "Zookeeper"
	CONSUL_CLIENT                   = "Consul"
	ETCD_CLIENT                     = "Etcd"
	DEFAULT_ZOOKEEPER_CLIENT_TIME   = 10
	DEFAULT_CONSUL_CHECK_TIMEOUT    = "1s"
	DEFAULT_CONSUL_CHECK_INTERVAL   = "3s"
	DEFAULT_CONSUL_CHECK_DEREGISTER = "15s"
	DEFAULT_ETCD_TIMEOUT            = 5

	//System default key
	ADMIN_USERNAME                  = "userName"
	ADMIN_PASSWORD                  = "password"
	DEFAULT_CONNECTION              = "Connection"
	DEFAULT_CONTENT_TYPE            = "Content-Type"
	DEFAULT_CONNECTION_VALUE        = "Keep-Alive"
	DEFAULT_CONTENT_TYPE_VALUE      = "application/json"
	DEFAULT_TOKEN_HEADER_KEY        = "X-Access-Token"
	DEFAULT_ADMIN_TOKEN_PARAM_ERROR = 500
	RPCTYPE_HTTP                    = "http"
	RPCTYPE_GRPC                    = "grpc"
	DEFAULT_BASE_PATH               = "/shenyu-client"
	REGISTER_URI                    = "/register-uri"
	REGISTER_METADATA               = "/register-metadata"
	DEFAULT_SHENYU_TOKEN            = "/platform/login"
	DEFAULT_CLIENT_ERRORCODE        = "SDK.ShenYuError"

	//System default error
	MISS_PARAM_ERROR_CODE        = "400"
	MISS_PARAM_ERROR_MSG         = "MetaDataRegister struct miss require param"
	MISS_SHENYU_ADMIN_ERROR_CODE = "503"
	MISS_SHENYU_ADMIN_ERROR_MSG  = "Please check ShenYu admin service status"
)
