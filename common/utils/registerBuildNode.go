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

package utils

import (
"fmt"
"github.com/apache/shenyu-client-golang/common/constants"
"github.com/apache/shenyu-client-golang/model"
"strings"
)

/*
 BuildMeataNodeName
 */
func BuildMetadataNodeName(metadata model.MetaDataRegister) string {
	var nodeName string
	var rpcType = metadata.RPCType

	if constants.RPCTYPE_HTTP == rpcType || constants.RPCTYPE_SPRING_CLOUND == rpcType {
	var ruleName = strings.Replace(metadata.RuleName,constants.PathSeparator, constants.SelectorJoinRule,-1)
	 nodeName = fmt.Sprintf("%s%s%s", metadata.ContextPath, constants.SelectorJoinRule,ruleName)
	} else {
	nodeName = BuildNodeName(metadata.ServiceName, metadata.MethodName)
	}
	 if strings.HasSuffix(nodeName,constants.PathSeparator){
		return nodeName[1:]
	 }
	 return nodeName
}

/**
 BuildMetadaDto
 */
func BuildMetadataDto(metadata *model.MetaDataRegister){
	metadata.Path = metadata.ContextPath + metadata.Path
	if metadata.RuleName == "" {
		metadata.RuleName = metadata.Path
	}
	if metadata.RPCType == ""{
		metadata.RPCType = constants.RPCTYPE_HTTP
	}
}

/**
BuildURINodeName
 */
func BuildURINodeName(registerDTO model.URIRegister ) string {
    var host = registerDTO.Host
    var port = registerDTO.Port
	var str = []string{host, port}
	return strings.Join(str,constants.COLONS)
}
