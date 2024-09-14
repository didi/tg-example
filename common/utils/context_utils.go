package utils

import (
	"fmt"
	"git.xiaojukeji.com/map-arch/tg-example/global/constants"
	"git.xiaojukeji.com/map-arch/tg-example/idl"
	"git.xiaojukeji.com/map-arch/tg-example/models"
	"github.com/didi/tg-flow/model"
)

func CheckRequestContext(sc *model.StrategyContext) (*models.RequestContext, error) {
	requestContextItf := sc.Get(constants.CONTEXTKEY_REQUEST_INFO)
	var err error
	if requestContextItf == nil {
		err = fmt.Errorf("key:%v not found", constants.CONTEXTKEY_REQUEST_INFO)
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

func CheckItemsInfo(sc *model.StrategyContext, key string) ([]*idl.ItemType, error) {
	itf := sc.Get(key)
	var err error
	if itf == nil {
		err = fmt.Errorf("key:%v not found", key)
		return nil, err
	}

	var items []*idl.ItemType
	var ok bool
	if items, ok = itf.([]*idl.ItemType); !ok {
		err = fmt.Errorf("recResp:%v type must be []*idl.ItemType", itf)
		return nil, err
	}

	return items, nil
}

func CheckResponseInfo(sc *model.StrategyContext) (*idl.ResponseInfo, error) {
	recRespItf := sc.Get(constants.CONTEXTKEY_RESPONSE_INFO)
	var err error
	if recRespItf == nil {
		err = fmt.Errorf("key:%v not found", constants.CONTEXTKEY_RESPONSE_INFO)
		return nil, err
	}

	var recResp *idl.ResponseInfo
	var ok bool
	if recResp, ok = recRespItf.(*idl.ResponseInfo); !ok {
		err = fmt.Errorf("recResp:%v type must be *models.ResponseInfo", recRespItf)
		return nil, err
	}

	return recResp, nil
}
