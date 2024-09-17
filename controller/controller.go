/**
 *
 * Description: entrence of service
 * Author: 		dayunzhangyunfeng@didiglobal.com
 * Since: 		2020-07-14 11:51:20
 */

package controller

import (
"fmt"
"github.com/didi/tg-example/constants"
"github.com/didi/tg-example/logic/dispatcher"
"github.com/didi/tg-example/models"
"github.com/didi/tg-flow/common/timeutils"
"github.com/didi/tg-flow/common/tlog"
"github.com/gin-gonic/gin"
"net/http"

)

type ExampleService struct {}

var RecDispatcher = &dispatcher.RecDispatcher{Name: "rec"}

func (h *ExampleService) Recommend(c *gin.Context) {
	respInfo, err := h.doRecommend(c)
	//fmt.Println("after h.doRecommend respInfo, err===>", respInfo, err)
	if err != nil{
		c.JSON(http.StatusNoContent, err)
	}
	c.JSON(http.StatusOK, respInfo)
}

func (h *ExampleService) doRecommend(c *gin.Context) (*models.ResponseInfo, error) {
	ctx := c.Request.Context()
	if err := recover(); err != nil {
		errMsg := fmt.Sprintf("%v", err)
		tlog.Handler.ErrorCount(ctx, constants.DLTagSystemPanic, errMsg)
		return models.CreateErrorResponseInfo(constants.ErrNoOther, errMsg), nil
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
		return models.CreateErrorResponseInfo(constants.ErrNoOther, errMsg), nil
	}

	//2.执行策略
	resultInfo, _ := dispatcher.DoStrategy(ctx, rc, RecDispatcher, tc)
	if resultInfo == nil {
		errMsg := "resultInfo is nil"
		tlog.Handler.ErrorCount(ctx, "DoStrategy_err", errMsg)
		return models.CreateErrorResponseInfo(constants.ErrNoOther, errMsg), nil
	}

	//3.解析结果
	appResponse := resultInfo.(*models.ResponseInfo)
	if appResponse == nil {
		errMsg := "resultInfo assert to *models.ResponseInfo fail"
		tlog.Handler.ErrorCount(ctx, "response_assert_err", fmt.Sprintf("errMsg:%v, resultInfo:%v", errMsg, resultInfo))
		return models.CreateErrorResponseInfo(constants.ErrNoOther, errMsg), nil
	}

	tc.StopCount()

	return appResponse, nil
}
