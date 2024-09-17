package utils

import (
	"fmt"
	"github.com/didi/tg-example/constants"
"github.com/didi/tg-example/models"
	"github.com/didi/tg-flow/model"
)

func CheckRequestContext(sc *model.StrategyContext) (*models.RequestContext, error) {
	requestContextItf := sc.Get(constants.ContextkeyRequestInfo)
	var err error
	if requestContextItf == nil {
		err = fmt.Errorf("key:%v not found", constants.ContextkeyRequestInfo)
		return nil, err
	}

	var requestContext *models.RequestContext
	var ok bool
	if requestContext, ok = requestContextItf.(*models.RequestContext); !ok {
		err = fmt.Errorf("requestRequest:%v type must be *models.RequestContext", requestContextItf)
		return nil, err
	}

	return requestContext, nil
}

func CheckItemsInfo(sc *model.StrategyContext, key string) ([]*models.ItemType, error) {
	itf := sc.Get(key)
	var err error
	if itf == nil {
		err = fmt.Errorf("key:%v not found", key)
		return nil, err
	}

	var items []*models.ItemType
	var ok bool
	if items, ok = itf.([]*models.ItemType); !ok {
		err = fmt.Errorf("recResp:%v type must be []*models.ItemType", itf)
		return nil, err
	}

	return items, nil
}

func CheckResponseInfo(sc *model.StrategyContext) (*models.ResponseInfo, error) {
	recRespItf := sc.Get(constants.ContextkeyResponseInfo)
	var err error
	if recRespItf == nil {
		err = fmt.Errorf("key:%v not found", constants.ContextkeyResponseInfo)
		return nil, err
	}

	var recResp *models.ResponseInfo
	var ok bool
	if recResp, ok = recRespItf.(*models.ResponseInfo); !ok {
		err = fmt.Errorf("recResp:%v type must be *models.ResponseInfo", recRespItf)
		return nil, err
	}

	return recResp, nil
}
