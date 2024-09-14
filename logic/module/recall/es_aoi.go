/*
	Description: 5路召回
	Author: dayunzhangyunfeng@didiglobal.com
	Since:	2020-10-28 01:27:30
*/

package recall

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

type EsAoi struct {
	wfengine.ModelBase
}

func (e EsAoi) DoAction(ctx context.Context, sc *model.StrategyContext) interface{} {
	defer utl.Recover(ctx, constants.ErrTypeActionPanic)

	//1. 取数据
	reqInfo, err := utils.CheckRequestContext(sc)
	if reqInfo == nil || err != nil {
		errMsg := fmt.Sprintf("%v", err)
		tlog.Handler.ErrorCount(ctx, "utils_check_RequestInfo_err", errMsg)
		sc.Skip(constants.ErrrNoOther, errMsg)

		return err
	}

	//2. 取上游节点的处理结果
	esGeneralItems, err := utils.CheckItemsInfo(sc, constants.CONTEXTKEY_ES_GENERAL_INFO)
	if err != nil{
		errMsg := fmt.Sprintf("%v", err)
		tlog.Handler.ErrorCount(ctx, "utils.checkItemsInfo_err", errMsg)
		//验证超时功能，找不到就忽略,无需退出
		//sc.Skip(constants.ErrrNoOther, errMsg)
		//return err
	}
	//3. 召回
	items := mock.MockEsAoiRecall(esGeneralItems, 35)
	sc.Set(constants.CONTEXTKEY_ES_AOI_INFO, items)

	fmt.Println(fmt.Sprintf("完成时间=%v ,actionName=%v", time.Now(), e.GetName()))
	return items
}
