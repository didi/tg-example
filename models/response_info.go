/**
*	function:	error result
*	Author	:	dayunzhangyunfeng@didiglobal.com
*	Since	:	2018-12-03 19:45:54
 */
package models

import (
	"git.xiaojukeji.com/map-arch/tg-example/idl"
)

func CreateErrorResponseInfo(errNo int32, errMsg string) *idl.ResponseInfo {
	respInfo := idl.NewResponseInfo()
	respInfo.Errno = errNo
	respInfo.Errmsg = errMsg

	//TODO ZYF add other attributes of ResponseInfo here

	return respInfo
}
