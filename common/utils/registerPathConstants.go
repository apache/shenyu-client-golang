package utils

import (
	"strings"
)

const (
	 REGISTER_URI_INSTANCE_PATH  = "/shenyu/register/uri/*/*/*"

	 REGISTER_METADATA_INSTANCE_PATH = "/shenyu/register/metadata/*/*/*"

	ROOT_PATH = "/shenyu/register"

	SEPARATOR = "/"

	DOT_SEPARATOR = "."
)

/**
  build child path of "/shenyu/register/metadata/{rpcType}/".
 */
func BuildMetaDataContextPathParent(rpcType string) string {
	var str = []string{ROOT_PATH, "metadata", rpcType}
	return strings.Join(str,SEPARATOR)
}

/**
* build child path of "/shenyu/register/metadata/{rpcType}/{contextPath}/".
 */
func BuildMetaDataParentPath(rpcType string,contextPath string) string {
	contextPath = RemovePrefix(contextPath)
	var str = []string{ROOT_PATH, "metadata", rpcType, contextPath}
	return strings.Join(str,SEPARATOR)
}

/**
* Build uri path string.
* build child path of "/shenyu/register/uri/{rpcType}/".
 */
func BuildURIContextPathParent(rpcType string) string {
	var str = []string{ROOT_PATH, "uri", rpcType}
	return strings.Join(str,SEPARATOR)
}

/**
Build uri path string.
build child path of "/shenyu/register/uri/{rpcType}/{contextPath}/".
 */
func BuildURIParentPath(rpcType string,contextPath string) string {
	contextPath = RemovePrefix(contextPath)
	var str = []string{ROOT_PATH,"uri", rpcType, contextPath}
	return strings.Join(str,SEPARATOR)
}

/**
Build instance parent path string.
build child path of "/shenyu/register/instance/
 */
func BuildInstanceParentPath() string {
	var str = []string{ ROOT_PATH, "instance"}
	return strings.Join(str,SEPARATOR)
}

/**
Build real node string.
 */
func BuildRealNode(nodePath string,nodeName string) string {
	nodePath = RemoveSuffix(nodePath)
	nodeName = RemovePrefix(nodeName)
	var str = []string{nodePath, nodeName}
	return strings.Join(str,SEPARATOR)
}

/**
Build nacos instance service path string.
 build child path of "shenyu.register.service.{rpcType}".
 */
func BuildServiceInstancePath(rpcType string) string {
	var str = []string{ROOT_PATH, "service", rpcType}
	var str2 = strings.Replace(strings.Join(str,SEPARATOR),"/", DOT_SEPARATOR,-1)
	return  str2[1:]
}

/**
Build nacos config service path string.
 build child path of "shenyu.register.service.{rpcType}.{contextPath}".
 */
func BuildServiceConfigPath(rpcType string,contextPath string) string {
	contextPath = RemovePrefix(contextPath)
   var str = []string{ROOT_PATH, "service", rpcType, contextPath}
   var str2 = strings.Replace(strings.Join(str,SEPARATOR),SEPARATOR, DOT_SEPARATOR,-1)
   var serviceConfigPathOrigin = strings.Replace(str2,"*", "",-1)

   var serviceConfigPathAfterSubstring = serviceConfigPathOrigin[1:]
 if strings.HasSuffix(serviceConfigPathAfterSubstring,".") {
  return serviceConfigPathAfterSubstring[0:len(serviceConfigPathOrigin)-1]
}
return serviceConfigPathAfterSubstring
}

/**
Build node name by DOT_SEPARATOR.
 */
func BuildNodeName(serviceName string,methodName string) string {
	var str = []string{serviceName, methodName}
	return strings.Join(str,DOT_SEPARATOR)
}

