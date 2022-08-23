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
}

/**
BuildURINodeName
 */
func BuildURINodeName(registerDTO model.URIRegister ) string {
    var host = registerDTO.Host
    var port = registerDTO.Port
	var str = []string{host, string(port)}
	return strings.Join(str,constants.COLONS)
}
