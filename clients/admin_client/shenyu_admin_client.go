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

package admin_client

import (
	"encoding/json"
	"github.com/apache/shenyu-client-golang/clients/http_client"
	"github.com/apache/shenyu-client-golang/common/constants"
	"github.com/apache/shenyu-client-golang/common/shenyu_error"
	"github.com/apache/shenyu-client-golang/model"
	"github.com/pkg/errors"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"net/http"
	"sync"
)

/**
 * The ShenYuAdminClient
 **/
type ShenYuAdminClient struct {
   Acp *ShenYuAdminClientParams
   AccessTokens *sync.Map
}

type ShenYuAdminClientParams struct {
	ServerList  []string
	UserName string
	Password string
}

/**
 * NewClient
 **/
func (acc *ShenYuAdminClient) NewClient(clientParam interface{}) (client interface{}, createResult bool, err error) {
	acp, ok := clientParam.(*ShenYuAdminClientParams)
	if !ok {
		logger.Fatal("The clientParam  must not nil!")
	}
	if len(acp.ServerList) == 0{
		logger.Fatal("The clientParam ServerList must not nil!")
	}
	//token handler
	var syMap *sync.Map
	for _,address := range acp.ServerList{
		var token,err = acc.setAccessToken(acp.UserName,acp.Password,address)
		if err != nil{
			logger.Fatal("init admin client error %+v:", err)
		}
		syMap.Store(address,token)
	}
	logger.Info("Create customer admin client success!")
	return &ShenYuAdminClient{
		Acp: &ShenYuAdminClientParams{
			ServerList: acp.ServerList,
			UserName: acp.UserName,
			Password: acp.Password,
		},AccessTokens: syMap,
	},true,nil
}

/**
PersistInterface
*/
func (acc *ShenYuAdminClient) PersistInterface(metaData interface{})(registerResult bool, err error){
	var metadata,ok =  metaData.(*model.MetaDataRegister)
	if !ok {
		logger.Fatal("get consul client metaData error %+v:", err)
	}
	registerResult,err = acc.registerMetaData(metadata)
	if err != nil{
		logger.Error("admin client register fail %+v",err)
		return registerResult, err
	}
	var metadaStr,_ = json.Marshal(metadata)
	logger.Info("%s admin client register success: %s",metadata.RPCType,string(metadaStr))
	return registerResult,nil
}

/**
PersistURI
*/
func (acc *ShenYuAdminClient) PersistURI(uriRegisterData interface{})(registerResult bool, err error){
	uriRegister,ok := uriRegisterData.(*model.URIRegister)
	if !ok {
		logger.Fatal("get admin client uriregister error %+v:", err)
	}
    registerResult,err =  acc.urlRegister(uriRegister)
	if err != nil{
		logger.Error("admin client register fail %+v",err)
		return registerResult, err
	}
	logger.Info("RegisterServiceInstance,result:%+v", true)
	return registerResult, nil
}

/**
Close
*/
func (acc *ShenYuAdminClient) Close(){

}

/**
 * Register metadata to ShenYu Gateway
 **/
func (acc *ShenYuAdminClient)  registerMetaData( metaData *model.MetaDataRegister) (registerResult bool, err error) {
	for _,server := range acc.Acp.ServerList {
		var token string
		var tokenStr, ok = acc.AccessTokens.Load(server)
		if ok {
			token = tokenStr.(string)
		} else {
			return false, errors.New("token load fail")
		}
		headers := adapterHeaders(token)

		params := map[string]string{}
		if metaData.AppName == "" || metaData.Path == "" || metaData.Host == "" || metaData.Port == "" {
			return false, shenyu_error.NewShenYuError(constants.MISS_PARAM_ERROR_CODE, constants.MISS_PARAM_ERROR_MSG, err)
		}
		params["appName"] = metaData.AppName
		params["path"] = metaData.Path
		params["contextPath"] = metaData.ContextPath
		params["host"] = metaData.Host
		params["port"] = metaData.Port

		if metaData.RPCType != "" {
			params["rpcType"] = metaData.RPCType
		} else {
			params["rpcType"] = constants.RPCTYPE_HTTP
		}

		if metaData.RuleName != "" {
			params["ruleName"] = metaData.RuleName
		} else {
			params["ruleName"] = metaData.Path
		}

		tokenRequest := initShenYuCommonRequest(headers, params, constants.REGISTER_METADATA, "", server)

		registerResult, err = http_client.RegisterMetaData(tokenRequest)
		if err != nil {
			return registerResult, err
		}
	}
	return registerResult,nil
}

/**
 * Url Register to ShenYu Gateway
 **/
func(acc *ShenYuAdminClient) urlRegister(urlMetaData *model.URIRegister) (registerResult bool, err error) {
	for _,server := range acc.Acp.ServerList {
		var token string
		var tokenStr, ok = acc.AccessTokens.Load(server)
		if ok {
			token = tokenStr.(string)
		} else {
			return false, errors.New("token load fail")
		}
		headers := adapterHeaders(token)
		params := map[string]string{}
		if urlMetaData.AppName == "" || urlMetaData.RPCType == "" || urlMetaData.Host == "" || urlMetaData.Port == "" {
			return false, shenyu_error.NewShenYuError(constants.MISS_PARAM_ERROR_CODE, constants.MISS_PARAM_ERROR_MSG, err)
		}
		params["protocol"] = urlMetaData.Protocol
		params["appName"] = urlMetaData.AppName
		params["contextPath"] = urlMetaData.ContextPath
		params["host"] = urlMetaData.Host
		params["port"] = urlMetaData.Port
		params["rpcType"] = urlMetaData.RPCType

		tokenRequest := initShenYuCommonRequest(headers, params, constants.REGISTER_URI, "", server)

		registerResult, err = http_client.DoUrlRegister(tokenRequest)
		if err != nil {
			return registerResult, err
		}
	}
	return true,nil
}

/*
setAccessTomen
 */
func (acc *ShenYuAdminClient) setAccessToken(userName string,password string,server string ) (string,error){
	headers := map[string][]string{}
	headers[constants.DEFAULT_CONNECTION] = []string{constants.DEFAULT_CONNECTION_VALUE}
	headers[constants.DEFAULT_CONTENT_TYPE] = []string{constants.DEFAULT_CONTENT_TYPE_VALUE}

	params := map[string]string{}
	if userName == "" || password == "" {
		params[constants.ADMIN_USERNAME] = constants.DEFAULT_ADMIN_ACCOUNT
		params[constants.ADMIN_PASSWORD] = constants.DEFAULT_ADMIN_PASSWORD
	} else {
		params[constants.ADMIN_USERNAME] = userName
		params[constants.ADMIN_PASSWORD] = password
	}

	tokenRequest := initShenYuCommonRequest(headers, params, constants.DEFAULT_SHENYU_TOKEN, "token",server)
	var token, err = getShenYUAdminToken(tokenRequest)
	if err != nil{
		logger.Error("get token fail")
		return "",err
	}
	if token == ""{
		logger.Error("get token fail,is empty %+v",err)
		return "",errors.New("get token is empty")
	}
	return token,nil
}

/**
 * initShenYuCommonRequest
 **/
func initShenYuCommonRequest(headers map[string][]string, params map[string]string, requestUrl string, busType string,serverList string) *model.ShenYuCommonRequest {
	if serverList == ""{
		serverList = constants.DEFAULT_SHENYU_ADMIN_URL
	}
	url := ""
	if len(busType) > 0 {
		url = serverList + requestUrl //get Token
	} else {
		url = serverList + constants.DEFAULT_BASE_PATH + requestUrl //register
	}
	tokenRequest := &model.ShenYuCommonRequest{
		Url:       url,
		Header:    headers,
		Params:    params,
		TimeoutMs: constants.DEFAULT_REQUEST_TIME,
	}
	return tokenRequest
}

/**
 get token
 */
func getShenYUAdminToken(shenYuCommonRequest *model.ShenYuCommonRequest) (token string, err error) {
	var response *http.Response
	response, err = shenYuCommonRequest.HttpClient.Request(http.MethodGet, shenYuCommonRequest.Url, shenYuCommonRequest.Header, constants.DEFAULT_REQUEST_TIME, shenYuCommonRequest.Params)
	if err != nil {
		return
	}
	var adminToken model.AdminToken
	var bytes []byte
	bytes, err = ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	err = json.Unmarshal(bytes, &adminToken)

	if err != nil {
		return
	}
	logger.Info("Get ShenYu Admin response, body is ->", adminToken)
	if response.StatusCode == http.StatusOK && adminToken.Code == http.StatusOK {
		return adminToken.AdminTokenData.Token, nil
	}
	if adminToken.Code == constants.DEFAULT_ADMIN_TOKEN_PARAM_ERROR {
		return "", err
	}
	return "", err
}

/**
 * adapter require Headers
 **/
func adapterHeaders(token string ) map[string][]string {
	headers := map[string][]string{}
	headers[constants.DEFAULT_CONNECTION] = []string{constants.DEFAULT_CONNECTION_VALUE}
	headers[constants.DEFAULT_CONTENT_TYPE] = []string{constants.DEFAULT_CONTENT_TYPE_VALUE}
	headers[constants.DEFAULT_TOKEN_HEADER_KEY] = []string{token}
	return headers
}
