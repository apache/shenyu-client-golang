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
	"github.com/incubator-shenyu-client-golang/common/constants"
	"github.com/incubator-shenyu-client-golang/model"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"net/http"
)

/**
 * The ShenYuAdminClient
 **/
type ShenYuAdminClient struct {
	UserName string `json:"userName"` //user optional
	Password string `json:"password"` //user optional
}

func (client *ShenYuAdminClient) GetShenYuAdminUser(shenYuCommonRequest model.ShenYuCommonRequest) (adminTokenData model.AdminTokenData, err error) {
	var response *http.Response
	response, err = shenYuCommonRequest.HttpClient.Request("GET", shenYuCommonRequest.Url, shenYuCommonRequest.Header, constants.DEFAULT_REQUEST_TIME, shenYuCommonRequest.Params)
	if err != nil {
		return
	}
	var bytes []byte
	bytes, err = ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	var adminToken = model.AdminToken{}
	err = json.Unmarshal(bytes, &adminToken)

	if err != nil {
		return
	}
	logger.Info("Get body is ->", adminToken)
	if response.StatusCode == 200 {
		return model.AdminTokenData{
			ID:          adminToken.AdminTokenData.ID,
			UserName:    adminToken.AdminTokenData.UserName,
			Role:        adminToken.AdminTokenData.Role,
			Enabled:     adminToken.AdminTokenData.Enabled,
			DateCreated: adminToken.AdminTokenData.DateCreated,
			DateUpdated: adminToken.AdminTokenData.DateUpdated,
			Token:       adminToken.AdminTokenData.Token,
		}, nil
	}
	return model.AdminTokenData{}, nil
}
