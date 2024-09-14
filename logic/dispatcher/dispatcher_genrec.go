/**
*	function：dispatcher of genrec
*	Author	:	dayunzhangyunfeng@didiglobal.com
*	Since	:	2018-12-03 19:45:54
 */
package dispatcher

import (
	"context"
	"encoding/json"
	"git.xiaojukeji.com/map-arch/tg-example/global/constants"
	"git.xiaojukeji.com/map-arch/tg-example/idl"
	"git.xiaojukeji.com/map-arch/tg-example/models"
	"github.com/didi/tg-flow/model"
	"strconv"
)

type DispatcherObj struct {
	Dispatcher
	InterfaceName string
}

func (d *DispatcherObj) BuildRequest(ctx context.Context, requestParam interface{}) *model.StrategyContext {
	r := requestParam.(*models.RequestContext)
	sc := model.NewStrategyContext(ctx)
	sc.SceneId = r.SceneId
	sc.FlowId = r.FlowId

	//如果用r.Pid分桶就赋值r.Pid给Userid, 如果用r.Phone分桶，就赋值r.Phone给UserId
	sc.UserId = strconv.FormatInt(r.Uid, 10)
	sc.Set(constants.CONTEXTKEY_REQUEST_INFO, r)

	//示例：如果有需要在工作流执行前初始化的参数，可在此初始化
	sc.Set("version", "6.0")

	return sc
}

func (d *DispatcherObj) BuildResponse(strategyContext *model.StrategyContext) interface{} {
	var btr *idl.ResponseInfo

	if strategyContext.IsLimited { //降级流量，走打底
		btr = doBackUpAction(strategyContext)
	} else {
		var ok bool
		itf := strategyContext.Get(constants.CONTEXTKEY_RESPONSE_INFO)
		if btr, ok = itf.(*idl.ResponseInfo); !ok { //策略异常，走打底
			btr = doBottomAction(strategyContext)
		}
	}

	strategyContext.Set(constants.CONTEXTKEY_RESPONSE_INFO, btr)
	return btr
}

func doBackUpAction(sc *model.StrategyContext) *idl.ResponseInfo {
	return doBottomAction(sc)
}

func doBottomAction(sc *model.StrategyContext) *idl.ResponseInfo {
	return models.CreateErrorResponseInfo(sc.ErrNo, sc.ErrMsg)
}

func (d *DispatcherObj) WriteLog(ctx context.Context, sc *model.StrategyContext) map[string]interface{} {
	var reqInfoStr string
	if itf := sc.Get(constants.CONTEXTKEY_REQUEST_INFO); itf != nil {
		reqInfo := itf.(*models.RequestContext)
		reqInfoBytes, err := json.Marshal(*reqInfo)
		if err == nil {
			reqInfoStr = string(reqInfoBytes)
		}
	}

	var respInfoStr string
	if itf := sc.Get(constants.CONTEXTKEY_RESPONSE_INFO); itf != nil {
		respInfo := itf.(*idl.ResponseInfo)
		respInfoBytes, err := json.Marshal(*respInfo)
		if err == nil {
			respInfoStr = string(respInfoBytes)
		}
	}

	pairs := make(map[string]interface{})
	//TODO ZYF interface的值需改为自己的接口名
	pairs["interface"] = d.GetInterfaceName()
	pairs["request"] = reqInfoStr
	pairs["response"] = respInfoStr
	return pairs
}

func (d *DispatcherObj) GetInterfaceName() string {
	return d.InterfaceName
}
