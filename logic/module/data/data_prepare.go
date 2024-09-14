/**
*	function:	prepare some data for following modules
*	Author	:	dayunzhangyunfeng@didiglobal.com
*	Since	:	2020-07-07 19:35:00
 */

package data

import (
	"context"
	"fmt"
	"github.com/didi/tg-example/common/utils"
	"github.com/didi/tg-example/global/constants"
	"github.com/didi/tg-example/logic/module/mock"
	"github.com/didi/tg-flow/common/tlog"
utl "github.com/didi/tg-flow/common/utils"
	"github.com/didi/tg-flow/model"
	"github.com/didi/tg-flow/wfengine"
	"time"
)

type DataPrepare struct {
	wfengine.ModelBase
	ParamA string
	ParamB string
}

func (d *DataPrepare) DoAction(ctx context.Context, sc *model.StrategyContext) interface{} {
	defer utl.Recover(ctx, constants.ErrTypeActionPanic)

	reqInfo, err := utils.CheckRequestContext(sc)
	if reqInfo == nil || err != nil {
		errMsg := fmt.Sprintf("%v", err)
		tlog.Handler.ErrorCount(ctx, "utils_check_RequestInfo_err", errMsg)
		sc.Skip(constants.ErrrNoOther, errMsg)

		return err
	}

	// 演示在下游的条件节点之前设置条件所依赖的参数
	sc.Set("isRecallHot",true)

	reqInfo.UserProfile = mock.MockUserProfileFromRedis(reqInfo.Uid)
	sc.Set(constants.CONTEXTKEY_REQUEST_INFO, reqInfo)

	fmt.Println(fmt.Sprintf("完成时间=%v ,actionName=%v", time.Now(), "DataPrepare"))
	return reqInfo
}

