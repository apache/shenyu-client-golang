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
	"fmt"
	"github.com/apache/shenyu-client-golang/clients"
	"github.com/apache/shenyu-client-golang/common/constants"
	"github.com/apache/shenyu-client-golang/model"
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
		fmt.Printf("this is ShenYu Admin client token %v ->", adminToken.AdminTokenData.Token)
	}

	//init MetaDataRegister
	metaData := &model.MetaDataRegister{
		AppName:     "testGoAppName",     //require user provide
		Path:        "/golang/your/path", //require user provide
		ContextPath: "/golang",           //require user provide
		Enabled:     true,                //require user provide
		Host:        "127.0.0.1",         //require user provide
		Port:        "8080",              //require user provide
	}
	result, err := clients.RegisterMetaData(adminToken.AdminTokenData, metaData)
	if err != nil {
		fmt.Printf("MetaDataRegister has error %v:", err)
	}
	fmt.Printf("finish register metadata ,the result is %v ->", result)

	//init urlRegister
	urlRegister := &model.URIRegister{
		Protocol:    "http://",              //require user provide
		AppName:     "testGoAppName",        //require user provide
		ContextPath: "/golang",              //require user provide
		RPCType:     constants.RPCTYPE_HTTP, //require user provide
		Host:        "127.0.0.1",            //require user provide
		Port:        "8080",                 //require user provide
	}
	result, err = clients.UrlRegister(adminToken.AdminTokenData, urlRegister)
	if err != nil {
		fmt.Printf("UrlRegister has error %v:", err)
	}
	fmt.Printf("finish UrlRegister ,the result is %v ->", result)

	//do you logic
}
