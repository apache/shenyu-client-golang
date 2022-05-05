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

package main

import (
	"github.com/incubator-shenyu-client-golang/clients"
	"github.com/incubator-shenyu-client-golang/common/constants"
	"github.com/incubator-shenyu-client-golang/model"
	"github.com/wonderivan/logger"
)

/**
 * The shenyu_http_client example
 **/
func main() {

	//init ShenYuAdminClient
	adminClient := &model.ShenYuAdminClient{
		UserName: "admin",  //user provide
		Password: "123456", //user provide
	}

	adminToken, err := clients.NewShenYuAdminClient(adminClient)
	if err == nil {
		logger.Info("this is ShenYu Admin client token ->", adminToken.AdminTokenData.Token)
	}

	//init MetaDataRegister
	metaData := &model.MetaDataRegister{
		AppName: "testMetaDataRegister", //require user provide
		Path:    "/your/path",           //require user provide
		Enabled: true,                   //require user provide
		Host:    "127.0.0.1",            //require user provide
		Port:    "8080",                 //require user provide
	}
	result, err := clients.RegisterMetaData(adminToken.AdminTokenData, metaData)
	if err != nil {
		logger.Warn("MetaDataRegister has error:", err)
	}
	logger.Info("finish register metadata ,the result is->", result)

	//init urlRegister
	urlRegister := &model.URIRegister{
		Protocol:    "testMetaDataRegister", //require user provide
		AppName:     "testURLRegister",      //require user provide
		ContextPath: "contextPath",          //require user provide
		RPCType:     constants.RPCTYPE_HTTP, //require user provide
		Host:        "127.0.0.1",            //require user provide
		Port:        "8080",                 //require user provide
	}
	result, err = clients.UrlRegister(adminToken.AdminTokenData, urlRegister)
	if err != nil {
		logger.Warn("UrlRegister has error:", err)
	}
	logger.Info("finish UrlRegister ,the result is->", result)

	//do you logic
}
