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
	"github.com/wonderivan/logger"
	"io/ioutil"
	"net/http"
	"testing"
)

/**
 * The http_client test
 **/
type object struct {
	Url        string            `param:"url"`
	Header     http.Header       `param:"header"`
	Params     map[string]string `param:"params"`
	TimeoutMs  uint64            `param:"timeoutMs"`
	httpClient HttpClient        `param:"httpClient"`
}

/**
 * Test http_client get ShenYu admin token
 **/
func TestHttpClientRequest(t *testing.T) {
	headers := map[string][]string{}
	headers["Connection"] = []string{"Keep-Alive"}
	headers["Content-Type"] = []string{"application/json"}

	params := map[string]string{}
	params["userName"] = constants.DEFAULT_ADMIN_ACCOUNT
	params["password"] = constants.DEFAULT_ADMIN_PASSWORD

	obj := &object{
		Url:       "http://127.0.0.1:9095" + constants.DEFAULT_SHENYU_TOKEN,
		Header:    headers,
		Params:    params,
		TimeoutMs: 1000,
	}

	var response *http.Response
	response, err := obj.httpClient.Request("GET", obj.Url, obj.Header, 1000, obj.Params)
	if err != nil {
		return
	}
	var bytes []byte
	bytes, err = ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return
	}
	/*var adminToken = model.AdminToken{}
	err = json.Unmarshal(bytes, &adminToken)*/
	logger.Info("Get body is ->", string(bytes))
	if response.StatusCode == 200 {
		return
	} else {
		return
	}
}
