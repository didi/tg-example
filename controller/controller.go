/**
 *
 * Description: entrence of service
 * Author: 		dayunzhangyunfeng@didiglobal.com
 * Since: 		2020-07-14 11:51:20
 */

package controller

import (
"fmt"
"git.xiaojukeji.com/map-arch/tg-example/global/constants"
"git.xiaojukeji.com/map-arch/tg-example/idl"
"git.xiaojukeji.com/map-arch/tg-example/logic/dispatcher"
"git.xiaojukeji.com/map-arch/tg-example/models"
"github.com/didi/tg-flow/common/timeutils"
"github.com/didi/tg-flow/common/tlog"
"github.com/gin-gonic/gin"
"net/http"

)

type ExampleService struct {}

var dispatcherObj = &dispatcher.DispatcherObj{InterfaceName: "search"}

func (h *ExampleService) Query(c *gin.Context) {
	respInfo, err := h.doQuery(c)
	//fmt.Println("after h.doQuery respInfo, err===>", respInfo, err)
	if err != nil{
		c.JSON(http.StatusNoContent, err)
	}
	c.JSON(http.StatusOK, respInfo)
}

func (h *ExampleService) doQuery(c *gin.Context) (*idl.ResponseInfo, error) {
	ctx := c.Request.Context()
	if err := recover(); err != nil {
		errMsg := fmt.Sprintf("%v", err)
		tlog.Handler.ErrorCount(ctx, constants.DLTagSystemPanic, errMsg)
		return models.CreateErrorResponseInfo(constants.ErrrNoOther, errMsg), nil
	}

	tc := timeutils.NewTimeCosterUnit(timeutils.TimeUnitMillSecond)
	tc.StartCount()

	//1.参数校验
	tc.StartSectionCount("getRequestContext")
	rc, err := getRequestContext(c)
	tc.StopSectionCount("getRequestContext")
	if err != nil {
		errMsg := fmt.Sprintf("%v", err)
		tlog.Handler.ErrorCount(ctx, "getRequestContext_err", errMsg)
		return models.CreateErrorResponseInfo(constants.ErrrNoOther, errMsg), nil
	}

	//2.执行策略
	resultInfo, _ := dispatcher.DoStrategy(ctx, rc, dispatcherObj, tc)
	if resultInfo == nil {
		errMsg := "resultInfo is nil"
		tlog.Handler.ErrorCount(ctx, "DoStrategy_err", errMsg)
		return models.CreateErrorResponseInfo(constants.ErrrNoOther, errMsg), nil
	}

	//3.解析结果
	appResponse := resultInfo.(*idl.ResponseInfo)
	if appResponse == nil {
		errMsg := "resultInfo assert to *idl.ResponseInfo fail"
		tlog.Handler.ErrorCount(ctx, "response_assert_err", fmt.Sprintf("errMsg:%v, resultInfo:%v", errMsg, resultInfo))
		return models.CreateErrorResponseInfo(constants.ErrrNoOther, errMsg), nil
	}

	tc.StopCount()

	return appResponse, nil
}
