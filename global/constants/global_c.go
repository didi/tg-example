/**
*	function:	global constants
*	Author	:	dayunzhangyunfeng@didiglobal.com
*	Since	:	2018-12-03 19:46:54
 */
package constants

const (
	//系统id
	CurrentAppId int64 = 200

	CurrentSystemName = "tg-example"

	//上报至metrics的统计曲线名称
	METRICS_NAME = "rpc_tg_example_rt"

	ErrTypeActionPanic = "action_panic"

	SceneIdTgExample = 20000

	ExperimentTgExample = "tg-example-exp"

	//tag类必须首字母是空格，否则把脉无法采集到该条日志
	DLTagSystemPanic  = " system_panic"

	ErrNoOK          = 0
	ErrNoUnknown     = 1    //未知错误，错误有值，但是err不为空
	ErrNoArgsInvalid = 2    //参数解析错误
	ErrNoResultEmpty = 3    // 结果为空
	ErrrNoOther      = 1000 //其它

	CONTEXTKEY_REQUEST_INFO  = "request_info"  //controller传来的请求信息
	CONTEXTKEY_RESPONSE_INFO = "response_Info" //传回给controller的返回信息

	CONTEXTKEY_ES_AOI_INFO		= "es_aoi_info"
	CONTEXTKEY_ES_GENERAL_INFO	= "es_general_info"
	CONTEXTKEY_ES_PRECISE_INFO	= "es_precise_info"
	CONTEXTKEY_HIST_RECALL_INFO	= "hist_recall_info"
	CONTEXTKEY_HOT_RECALL_INFO	= "hot_recall_info"
	CONTEXTKEY_QAC_RECALL_INFO	= "qac_recall_info"
	CONTEXTKEY_ROUGH_RANK_INFO	= "rough_rank_info"
	CONTEXTKEY_PRECISE_RANK_INFO= "precise_rank_info"
)
