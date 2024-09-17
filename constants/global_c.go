/**
*	function:	global constants
*	Author	:	dayunzhangyunfeng@didiglobal.com
*	Since	:	2018-12-03 19:46:54
 */
package constants

const (
	CurrentAppId int64 = 200

	ErrTypeActionPanic = "action_panic"

	SceneIdTgExample = 20000

	//tag类必须首字母是空格，否则把脉无法采集到该条日志
	DLTagSystemPanic  = " system_panic"

	ErrNoOK          = 0
	ErrNoUnknown     = 1    //未知错误，错误有值，但是err不为空
	ErrNoArgsInvalid = 2    //参数解析错误
	ErrNoResultEmpty = 3    // 结果为空
	ErrNoOther      = 1000 //其它

	ContextkeyRequestInfo  	= "request_info"  //controller传来的请求信息
	ContextkeyResponseInfo 	= "response_Info" //传回给controller的返回信息
	ContextkeyEsAoiInfo		= "es_aoi_info"
	ContextkeyEsGeneralInfo	= "es_general_info"
	ContextkeyEsPreciseInfo	= "es_precise_info"
	ContextkeyHistRecallInfo	= "hist_recall_info"
	ContextkeyHotRecallInfo	= "hot_recall_info"
	ContextkeyQacRecallInfo	= "qac_recall_info"
	ContextkeyRoughRankInfo	= "rough_rank_info"
	ContextkeyPreciseRankInfo	= "precise_rank_info"
)
