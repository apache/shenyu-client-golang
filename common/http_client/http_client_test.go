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
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type object struct {
	Url       string            `param:"url"`
	Header    http.Header       `param:"header"`
	Params    map[string]string `param:"params"`
	TimeoutMs uint64            `param:"timeoutMs"`
}

func TestHttpClientRequest(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.Method + "\n"))
	})
	http.ListenAndServe("127.0.0.1:9090", nil)

	assert.Equal(t, "fe01ce2a7fbac8fafaed7c982a04e229", "md5")
}
