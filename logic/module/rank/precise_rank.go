/*
Description: rank
Author:		dayunzhangyunfeng@didiglobal.com
Since:      2020-07-14 17:36:06
*/
package rank

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

type PreciseRank struct {
	wfengine.ModelBase
}

func (r PreciseRank) DoAction(ctx context.Context, sc *model.StrategyContext) interface{} {
	defer utl.Recover(ctx, constants.ErrTypeActionPanic)

	fmt.Println(fmt.Sprintf("actionName:%v,\tstart time:%v", r.GetName(), time.Now()))

	reqInfo, err := utils.CheckRequestContext(sc)
	if reqInfo == nil || err != nil {
		errMsg := fmt.Sprintf("%v", err)
		tlog.Handler.ErrorCount(ctx, "utils_check_RequestInfo_err", errMsg)
		sc.Skip(constants.ErrNoOther, errMsg)

		return err
	}

	items, err := utils.CheckItemsInfo(sc, constants.ContextkeyRoughRankInfo)
	if items == nil || err != nil {
		errMsg := fmt.Sprintf("%v", err)
		tlog.Handler.ErrorCount(ctx, "utils_CheckResponseInfo_err", errMsg)
		sc.Skip(constants.ErrNoOther, errMsg)

		return err
	}

	items = mock.MockPreciseRank(items)
	sc.Set(constants.ContextkeyPreciseRankInfo, items)
	fmt.Println(fmt.Sprintf("actionName:%v,\tfinish time:%v", r.GetName(), time.Now()))

	return items
}
