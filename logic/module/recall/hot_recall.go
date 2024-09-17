/*
	Description: 1路召回
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

type HotRecall struct {
	wfengine.ModelBase
}

func (h HotRecall) DoAction(ctx context.Context, sc *model.StrategyContext) interface{} {
	defer utl.Recover(ctx, constants.ErrTypeActionPanic)
	fmt.Println(fmt.Sprintf("actionName:%v,\tstart time:%v", h.GetName(), time.Now()))

	//1. 取数据
	reqInfo, err := utils.CheckRequestContext(sc)
	if reqInfo == nil || err != nil {
		errMsg := fmt.Sprintf("%v", err)
		tlog.Handler.ErrorCount(ctx, "utils_check_RequestInfo_err", errMsg)
		sc.Skip(constants.ErrNoOther, errMsg)

		return err
	}

	//TODO do somthing about es general recall
	items := mock.MockHotRecall(20)
	sc.Set(constants.ContextkeyHotRecallInfo, items)

	fmt.Println(fmt.Sprintf("actionName:%v,\tfinish time:%v", h.GetName(), time.Now()))
	return items
}
