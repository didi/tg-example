/*
	Description: 3路召回
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

type QacRecall struct {
	wfengine.ModelBase
}

func (q QacRecall) DoAction(ctx context.Context, sc *model.StrategyContext) interface{} {
	defer utl.Recover(ctx, constants.ErrTypeActionPanic)
	fmt.Println(fmt.Sprintf("actionName:%v,\tstart time:%v", q.GetName(), time.Now()))

	//1. 提取请求上下文信息
	reqInfo, err := utils.CheckRequestContext(sc)
	if reqInfo == nil || err != nil {
		errMsg := fmt.Sprintf("%v", err)
		tlog.Handler.ErrorCount(ctx, "utils_check_RequestInfo_err", errMsg)
		sc.Skip(constants.ErrNoOther, errMsg)

		return err
	}

	//2. 执行qac召回
	items := mock.MockQacRecall(10)
	sc.Set(constants.ContextkeyQacRecallInfo, items)

	fmt.Println(fmt.Sprintf("actionName:%v,\tfinish time:%v", q.GetName(), time.Now()))
	return items
}

func (q QacRecall) OnTimeout(ctx context.Context, sc *model.StrategyContext) {
	fmt.Println("execute timeout callback of " + q.GetName())
}
