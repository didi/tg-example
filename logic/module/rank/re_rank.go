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
	"git.xiaojukeji.com/map-arch/tg-example/idl"
	"git.xiaojukeji.com/map-arch/tg-example/logic/module/mock"
	"github.com/didi/tg-flow/common/tlog"
utl "github.com/didi/tg-flow/common/utils"
	"github.com/didi/tg-flow/model"
	"github.com/didi/tg-flow/wfengine"
	"time"
)

type ReRank struct {
	wfengine.ModelBase
}

func (r ReRank) DoAction(ctx context.Context, sc *model.StrategyContext) interface{} {
	defer utl.Recover(ctx, constants.ErrTypeActionPanic)

	reqInfo, err := utils.CheckRequestContext(sc)
	if reqInfo == nil || err != nil {
		errMsg := fmt.Sprintf("%v", err)
		tlog.Handler.ErrorCount(ctx, "utils_check_RequestInfo_err", errMsg)
		sc.Skip(constants.ErrrNoOther, errMsg)

		return err
	}

	items, err := utils.CheckItemsInfo(sc, constants.CONTEXTKEY_PRECISE_RANK_INFO)
	if items == nil || err != nil {
		errMsg := fmt.Sprintf("%v", err)
		tlog.Handler.ErrorCount(ctx, "utils_CheckItemsInfo_err", errMsg)
		sc.Skip(constants.ErrrNoOther, errMsg)

		return err
	}
	mock.MockReRank(items)

	data := &idl.DataType{
		Items: items,
		Version: "1.0",
	}
	responseInfo := &idl.ResponseInfo{
		Errno: 0,
		Errmsg: "",
		Data: data,
	}
	sc.Set(constants.CONTEXTKEY_RESPONSE_INFO, responseInfo)
	fmt.Println(fmt.Sprintf("完成时间=%v ,actionName=%v", time.Now(), r.GetName()))

	return responseInfo
}
