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
	"github.com/pkg/errors"
	"github.com/wonderivan/logger"
	"net/http"
)

type HttpClient struct {
}

func (client *HttpClient) Request(r *http.Request, url string, header http.Header, timeoutMs uint64, params map[string]string) (response *http.Response, err error) {
	switch r.Method {
	case http.MethodGet:
		response, err = client.Get(url, header, timeoutMs, params)
		return
	case http.MethodPost:
		response, err = client.Post(url, header, timeoutMs, params)
		return
	case http.MethodPut:
		response, err = client.Put(url, header, timeoutMs, params)
		return
	case http.MethodDelete:
		response, err = client.Delete(url, header, timeoutMs, params)
		return
	default:
		err = errors.New("not available method")
		logger.Error("request method[%s], url[%s],header:[%s],params:[%s], not available method ", r.Method, url, ToJsonString(header), ToJsonString(params))
	}
	return
}

func (client *HttpClient) Get(url string, header http.Header, timeoutMs uint64,
	params map[string]string) (response *http.Response, err error) {
	return get(url, header, timeoutMs, params)
}
func (client *HttpClient) Post(url string, header http.Header, timeoutMs uint64,
	params map[string]string) (response *http.Response, err error) {
	return post(url, header, timeoutMs, params)
}
func (client *HttpClient) Delete(url string, header http.Header, timeoutMs uint64,
	params map[string]string) (response *http.Response, err error) {
	return delete(url, header, timeoutMs, params)
}
func (client *HttpClient) Put(url string, header http.Header, timeoutMs uint64,
	params map[string]string) (response *http.Response, err error) {
	return put(url, header, timeoutMs, params)
}
