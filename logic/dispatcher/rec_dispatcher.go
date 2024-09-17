/**
*	function：dispatcher of genrec
*	Author	:	dayunzhangyunfeng@didiglobal.com
*	Since	:	2018-12-03 19:45:54
 */
package dispatcher

import (
	"context"
	"encoding/json"
	"github.com/didi/tg-example/constants"
"github.com/didi/tg-example/models"
	"github.com/didi/tg-flow/model"
	"strconv"
)

type RecDispatcher struct {
	Name string
}

func (d *RecDispatcher) BuildRequest(ctx context.Context, requestParam interface{}) *model.StrategyContext {
	r := requestParam.(*models.RequestContext)
	sc := model.NewStrategyContext(ctx)
	sc.SceneId = r.SceneId
	sc.FlowId = r.FlowId

	//如果用r.Pid分桶就赋值r.Pid给Userid, 如果用r.Phone分桶，就赋值r.Phone给UserId
	sc.UserId = strconv.FormatInt(r.Uid, 10)
	sc.Set(constants.ContextkeyRequestInfo, r)

	//示例：如果有需要在工作流执行前初始化的参数，可在此初始化
	sc.Set("version", "6.0")

	return sc
}

func (d *RecDispatcher) BuildResponse(strategyContext *model.StrategyContext) interface{} {
	var btr *models.ResponseInfo

	if strategyContext.IsLimited { //降级流量，走打底
		btr = doBottomAction(strategyContext)
	} else {
		var ok bool
		itf := strategyContext.Get(constants.ContextkeyResponseInfo)
		if btr, ok = itf.(*models.ResponseInfo); !ok { //策略异常，走打底
			btr = doBottomAction(strategyContext)
		}
	}

	strategyContext.Set(constants.ContextkeyResponseInfo, btr)
	return btr
}

func doBottomAction(sc *model.StrategyContext) *models.ResponseInfo {
	return models.CreateErrorResponseInfo(sc.ErrNo, sc.ErrMsg)
}

func (d *RecDispatcher) WriteLog(ctx context.Context, sc *model.StrategyContext) map[string]interface{} {
	var reqInfoStr string
	if itf := sc.Get(constants.ContextkeyRequestInfo); itf != nil {
		reqInfo := itf.(*models.RequestContext)
		reqInfoBytes, err := json.Marshal(*reqInfo)
		if err == nil {
			reqInfoStr = string(reqInfoBytes)
		}
	}

	var respInfoStr string
	if itf := sc.Get(constants.ContextkeyResponseInfo); itf != nil {
		respInfo := itf.(*models.ResponseInfo)
		respInfoBytes, err := json.Marshal(*respInfo)
		if err == nil {
			respInfoStr = string(respInfoBytes)
		}
	}

	pairs := make(map[string]interface{})
	//TODO ZYF interface的值需改为自己的接口名
	pairs["interface"] = d.GetName()
	pairs["request"] = reqInfoStr
	pairs["response"] = respInfoStr
	return pairs
}

func (d *RecDispatcher) GetName() string {
	return d.Name
}
