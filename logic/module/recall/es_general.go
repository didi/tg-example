/*
	Description: 4路召回
	Author: dayunzhangyunfeng@didiglobal.com
	Since:	2020-10-28 01:27:30
*/

package recall

import (
	"context"
	"fmt"
	"github.com/didi/tg-example/common/utils"
	"github.com/didi/tg-example/constants"
"github.com/didi/tg-example/logic/module/mock"
	"github.com/didi/tg-flow/common/tlog"
utl "github.com/didi/tg-flow/common/utils"
	"github.com/didi/tg-flow/model"
	"github.com/didi/tg-flow/wfengine"
	"time"
)

type EsGeneral struct {
	wfengine.ModelBase
}

func (e EsGeneral) DoAction(ctx context.Context, sc *model.StrategyContext) interface{} {
	defer utl.Recover(ctx, constants.ErrTypeActionPanic)
	fmt.Println(fmt.Sprintf("actionName:%v,\tstart time:%v", e.GetName(), time.Now()))

	//1. 取数据
	reqInfo, err := utils.CheckRequestContext(sc)
	if reqInfo == nil || err != nil {
		errMsg := fmt.Sprintf("%v", err)
		tlog.Handler.ErrorCount(ctx, "utils_check_RequestInfo_err", errMsg)
		sc.Skip(constants.ErrNoOther, errMsg)

		return err
	}

	//2. 取上游结果
	//userInfo, err := utils.CheckItemsInfo(sc, constants.CONTEXTKEY_REQUEST_INFO)
	//if userInfo == nil || err != nil {
	//	errMsg := fmt.Sprintf("%v", err)
	//	logger.Handler.ErrorCount(ctx, "utils_checkItemsInfo_err", errMsg)
	//	sc.Skip(constants.ErrNoOther, errMsg)
	//	return err
	//}

	//TODO do somthing about es general recall
	items := mock.MockEsGeneralRecall(reqInfo, 40)
	sc.Set(constants.ContextkeyEsGeneralInfo, items)

	fmt.Println(fmt.Sprintf("actionName:%v,\tfinish time:%v", e.GetName(), time.Now()))
	return items
}

func (e EsGeneral) OnTimeout(ctx context.Context, sc *model.StrategyContext) {
	fmt.Println("execute timeout callback of :"+e.GetName())
}
