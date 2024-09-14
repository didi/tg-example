/*
	Description: 6路召回
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

type EsPrecise struct {
	wfengine.ModelBase
}

func (e EsPrecise) DoAction(ctx context.Context, sc *model.StrategyContext) interface{} {
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
	esAoiItems, err := utils.CheckItemsInfo(sc, constants.CONTEXTKEY_ES_AOI_INFO)
	if esAoiItems == nil || err != nil{
		errMsg := fmt.Sprintf("err=%v, key=%v", err, constants.CONTEXTKEY_ES_AOI_INFO)
		tlog.Handler.ErrorCount(ctx, "utils.checkItemsInfo_err", errMsg)
		//验证超时功能，找不到就忽略,无需退出
		//sc.Skip(constants.ErrrNoOther, errMsg)
		//return err
	}

	//3. 召回
	items := mock.MockEsPreciseRecall(esAoiItems, 405)
	sc.Set(constants.CONTEXTKEY_ES_PRECISE_INFO, items)

	fmt.Println(fmt.Sprintf("完成时间=%v ,actionName=%v", time.Now(), e.GetName()))
	return items
}
