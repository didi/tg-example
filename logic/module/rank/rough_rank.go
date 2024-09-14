/*
Description: rank
Author:		dayunzhangyunfeng@didiglobal.com
Since:      2020-07-14 17:36:06
*/
package rank

import (
	"context"
	"fmt"
	"git.xiaojukeji.com/map-arch/tg-example/common/utils"
	"git.xiaojukeji.com/map-arch/tg-example/global/constants"
	"git.xiaojukeji.com/map-arch/tg-example/logic/module/mock"
	"github.com/didi/tg-flow/common/tlog"
utl "github.com/didi/tg-flow/common/utils"
	"github.com/didi/tg-flow/model"
	"github.com/didi/tg-flow/wfengine"
	"time"
	)

type RoughRank struct {
	wfengine.ModelBase
}

func (r RoughRank) DoAction(ctx context.Context, sc *model.StrategyContext) interface{} {
	defer utl.Recover(ctx, constants.ErrTypeActionPanic)

	reqInfo, err := utils.CheckRequestContext(sc)
	if reqInfo == nil || err != nil {
		errMsg := fmt.Sprintf("%v", err)
		tlog.Handler.ErrorCount(ctx, "utils_check_RequestInfo_err", errMsg)
		sc.Skip(constants.ErrrNoOther, errMsg)

		return err
	}

	//TODO do somthing about precise rank here
	esPreciseItems, err := utils.CheckItemsInfo(sc, constants.CONTEXTKEY_ES_PRECISE_INFO)
	histRecallItems, err := utils.CheckItemsInfo(sc, constants.CONTEXTKEY_HIST_RECALL_INFO)
	hotRecallItems, err := utils.CheckItemsInfo(sc, constants.CONTEXTKEY_HOT_RECALL_INFO)
	qacRecallItems, err := utils.CheckItemsInfo(sc, constants.CONTEXTKEY_QAC_RECALL_INFO)
	items := mock.MockRoughRank(esPreciseItems, histRecallItems, hotRecallItems, qacRecallItems)
	sc.Set(constants.CONTEXTKEY_ROUGH_RANK_INFO, items)

	fmt.Println(fmt.Sprintf("完成时间=%v ,actionName=%v", time.Now(), r.GetName()))

	return items
}
