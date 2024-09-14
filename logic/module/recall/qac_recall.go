/*
	Description: 3路召回
	Author: dayunzhangyunfeng@didiglobal.com
	Since:	2020-10-28 01:27:30
*/

package recall

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

type QacRecall struct {
	wfengine.ModelBase
}

func (q QacRecall) DoAction(ctx context.Context, sc *model.StrategyContext) interface{} {
	defer utl.Recover(ctx, constants.ErrTypeActionPanic)

	reqInfo, err := utils.CheckRequestContext(sc)
	if reqInfo == nil || err != nil {
		errMsg := fmt.Sprintf("%v", err)
		tlog.Handler.ErrorCount(ctx, "utils_check_RequestInfo_err", errMsg)
		sc.Skip(constants.ErrrNoOther, errMsg)

		return err
	}

	//TODO do somthing about es general recall
	items := mock.MockQacRecall(300)
	sc.Set(constants.CONTEXTKEY_QAC_RECALL_INFO, items)

	fmt.Println(fmt.Sprintf("完成时间=%v ,actionName=%v", time.Now(), q.GetName()))
	return items
}

func (q QacRecall) OnTimeout(ctx context.Context, sc *model.StrategyContext) {
	fmt.Println("execute timeout callback of " + q.GetName())
}
