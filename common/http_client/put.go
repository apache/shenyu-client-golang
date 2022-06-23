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
	"net/http"
	"strings"
	"time"
)

/**
 * The http_client put method Implement
 **/
func put(url string, header http.Header, timeoutMs uint64, params map[string]string) (response *http.Response, err error) {
	client := http.Client{}
	client.Timeout = time.Millisecond * time.Duration(timeoutMs)
	var body string
	for key, value := range params {
		if len(value) > 0 {
			body += key + "=" + value + "&"
		}
	}
	body = strings.TrimSuffix(body, "&")
	request, errNew := http.NewRequest(http.MethodPut, url, strings.NewReader(body))
	if errNew != nil {
		err = errNew
		return
	}
	request.Header = header
	resp, errDo := client.Do(request)
	if errDo != nil {
		err = errDo
	} else {
		response = resp
	}
	return
}
