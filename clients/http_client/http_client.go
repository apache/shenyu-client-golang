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

package http_client

import (
	"github.com/apache/shenyu-client-golang/common/constants"
	"github.com/apache/shenyu-client-golang/common/shenyu_error"
	"github.com/apache/shenyu-client-golang/model"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

/**
 * Register metadata to ShenYu Gateway
 **/
func RegisterMetaData(shenYuCommonRequest *model.ShenYuCommonRequest) (result bool, err error) {
	var response *http.Response
	response, err = shenYuCommonRequest.HttpClient.Request(http.MethodPost, shenYuCommonRequest.Url, shenYuCommonRequest.Header, constants.DEFAULT_REQUEST_TIME, shenYuCommonRequest.Params)
	return handleCommonResponse(response, err)

}

/**
 * Url Register to ShenYu Gateway
 **/
func DoUrlRegister(shenYuCommonRequest *model.ShenYuCommonRequest) (result bool, err error) {
	var response *http.Response
	response, err = shenYuCommonRequest.HttpClient.Request(http.MethodPost, shenYuCommonRequest.Url, shenYuCommonRequest.Header, constants.DEFAULT_REQUEST_TIME, shenYuCommonRequest.Params)
	return handleCommonResponse(response, err)
}

/**
 * handleCommonResponse
 **/
func handleCommonResponse(response *http.Response, err error) (bool, error) {
	if response == nil {
		err = shenyu_error.NewShenYuError(constants.MISS_SHENYU_ADMIN_ERROR_CODE, constants.MISS_SHENYU_ADMIN_ERROR_MSG, err)
		return false, err
	}
	var bytes []byte
	bytes, err = ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		return strings.Contains(string(bytes), constants.DEFAULT_ADMIN_SUCCESS), err
	} else {
		err = shenyu_error.NewShenYuError(strconv.Itoa(response.StatusCode), string(bytes), err)
		return false, err
	}
}
