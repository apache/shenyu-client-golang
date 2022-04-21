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
	"fmt"
	"github.com/incubator-shenyu-client-golang/common/constants"
	"github.com/incubator-shenyu-client-golang/model"
)

/**
 * Get ShenYuAdminClient
 **/
func NewShenYuAdminClient(client model.ShenYuAdminClient) (adminTokenData model.AdminTokenData, err error) {
	headers := map[string][]string{}
	headers[constants.DEFAULT_CONNECTION] = []string{constants.DEFAULT_CONNECTION_VALUE}
	headers[constants.DEFAULT_CONTENT_TYPE] = []string{constants.DEFAULT_CONTENT_TYPE_VALUE}

	params := map[string]string{}
	params[constants.ADMIN_USERNAME] = client.UserName
	params[constants.ADMIN_PASSWORD] = client.Password

	tokenRequest := &model.ShenYuCommonRequest{
		Url:       constants.DEFAULT_SHENYU_ADMIN_URL + constants.DEFAULT_SHENYU_TOKEN,
		Header:    headers,
		Params:    params,
		TimeoutMs: constants.DEFAULT_REQUEST_TIME,
	}

	//todo use  GetShenYuAdminUser
	fmt.Print(tokenRequest)

	return model.AdminTokenData{}, nil
}
