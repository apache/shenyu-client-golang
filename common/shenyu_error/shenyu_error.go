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

package shenyu_error

/**
 * ShenYuError
 **/
import (
	"fmt"
	"github.com/apache/shenyu-client-golang/common/constants"
)

type ShenYuError struct {
	originError error
	errorCode   string
	errMsg      string
}

func NewShenYuError(errorCode string, errMsg string, originError error) *ShenYuError {
	return &ShenYuError{
		errorCode:   errorCode,
		errMsg:      errMsg,
		originError: originError,
	}

}

func (err *ShenYuError) Error() (str string) {
	shenYuErrMsg := fmt.Sprintf("The errCode is ->:%+v, The errMsg is  ->:%+v \n\n", err.ErrorCode(), err.errMsg)
	if err.originError != nil {
		return shenYuErrMsg + "caused by:\n" + err.originError.Error()
	}
	return shenYuErrMsg
}

func (err *ShenYuError) ErrorCode() string {
	if err.errorCode == "" {
		return constants.DEFAULT_CLIENT_ERRORCODE
	} else {
		return err.errorCode
	}
}
