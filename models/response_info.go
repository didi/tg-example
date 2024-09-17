/**
*	function:	error result
*	Author	:	dayunzhangyunfeng@didiglobal.com
*	Since	:	2018-12-03 19:45:54
 */
package models

type ResponseInfo struct {
	Errno  int32     `thrift:"errno,1,required" json:"errno"`
	Errmsg string    `thrift:"errmsg,2,required" json:"errmsg"`
	Data   *DataType `thrift:"data,3,required" json:"data"`
}

type DataType struct {
	Items          []*ItemType      `thrift:"items,1,required" json:"items"`
	Version        string           `thrift:"version,2,required" json:"version"`
}

type ItemType struct {
	Name               string  `thrift:"name,1,required" json:"name"`
	Price              float64 `thrift:"price,2,required" json:"price"`
	Score              float64   `thrift:"score,3,required" json:"score"`
	GoodsId            int64   `thrift:"goods_id,4,required" json:"goods_id"`
}

func NewResponseInfo() *ResponseInfo {
	return &ResponseInfo{}
}

func CreateErrorResponseInfo(errNo int32, errMsg string) *ResponseInfo {
	respInfo := NewResponseInfo()
	respInfo.Errno = errNo
	respInfo.Errmsg = errMsg
	//TODO ZYF add other attributes of ResponseInfo here

	return respInfo
}
