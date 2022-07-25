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

package clients

import (
	"github.com/apache/shenyu-client-golang/clients/admin_client"
	"github.com/apache/shenyu-client-golang/clients/http_client"
	"github.com/apache/shenyu-client-golang/common/constants"
	"github.com/apache/shenyu-client-golang/common/shenyu_error"
	"github.com/apache/shenyu-client-golang/model"
	"reflect"
)

/**
 * Get ShenYuAdminClient
 **/
func NewShenYuAdminClient(client *model.ShenYuAdminClient) (adminToken model.AdminToken, err error) {
	headers := map[string][]string{}
	headers[constants.DEFAULT_CONNECTION] = []string{constants.DEFAULT_CONNECTION_VALUE}
	headers[constants.DEFAULT_CONTENT_TYPE] = []string{constants.DEFAULT_CONTENT_TYPE_VALUE}

	params := map[string]string{}
	if reflect.DeepEqual(client, model.ShenYuAdminClient{}) || client.UserName == "" || client.Password == "" {
		params[constants.ADMIN_USERNAME] = constants.DEFAULT_ADMIN_ACCOUNT
		params[constants.ADMIN_PASSWORD] = constants.DEFAULT_ADMIN_PASSWORD
	} else {
		params[constants.ADMIN_USERNAME] = client.UserName
		params[constants.ADMIN_PASSWORD] = client.Password
	}

	tokenRequest := initShenYuCommonRequest(headers, params, constants.DEFAULT_SHENYU_TOKEN, "token")

	adminToken, err = admin_client.GetShenYuAdminUser(tokenRequest)
	if err == nil {
		return adminToken, nil
	} else {
		return model.AdminToken{}, err
	}
}

/**
 * Register metadata to ShenYu Gateway
 **/
func RegisterMetaData(adminTokenData model.AdminTokenData, metaData *model.MetaDataRegister) (registerResult bool, err error) {
	headers := adapterHeaders(adminTokenData)

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

	tokenRequest := initShenYuCommonRequest(headers, params, constants.REGISTER_METADATA, "")

	registerResult, err = http_client.RegisterMetaData(tokenRequest)
	if err == nil {
		return registerResult, nil
	} else {
		return false, err
	}
}

/**
 * Url Register to ShenYu Gateway
 **/
func UrlRegister(adminTokenData model.AdminTokenData, urlMetaData *model.URIRegister) (registerResult bool, err error) {
	headers := adapterHeaders(adminTokenData)

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

	tokenRequest := initShenYuCommonRequest(headers, params, constants.REGISTER_URI, "")

	registerResult, err = http_client.DoUrlRegister(tokenRequest)
	if err == nil {
		return registerResult, nil
	} else {
		return false, err
	}
}

/**
 * initShenYuCommonRequest
 **/
func initShenYuCommonRequest(headers map[string][]string, params map[string]string, requestUrl string, busType string) *model.ShenYuCommonRequest {
	url := ""
	if len(busType) > 0 {
		url = constants.DEFAULT_SHENYU_ADMIN_URL + requestUrl //get Token
	} else {
		url = constants.DEFAULT_SHENYU_ADMIN_URL + constants.DEFAULT_BASE_PATH + requestUrl //register
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
 * adapter require Headers
 **/
func adapterHeaders(adminTokenData model.AdminTokenData) map[string][]string {
	headers := map[string][]string{}
	headers[constants.DEFAULT_CONNECTION] = []string{constants.DEFAULT_CONNECTION_VALUE}
	headers[constants.DEFAULT_CONTENT_TYPE] = []string{constants.DEFAULT_CONTENT_TYPE_VALUE}
	headers[constants.DEFAULT_TOKEN_HEADER_KEY] = []string{adminTokenData.Token}
	return headers
}
