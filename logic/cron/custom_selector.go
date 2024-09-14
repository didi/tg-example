package cron

import (
	"errors"
	"fmt"
	"github.com/didi/tg-flow/model"
	"github.com/didi/tg-flow/wfengine"
	"hash/crc32"
	"strconv"
)

type CustomSelector struct {
}

/**
	本FlowSelector的实现仅为演示分桶逻辑的示例，具体使用时需要根据场景来自行定义
 */
func (r *CustomSelector) SelectWorkflowId(sc *model.StrategyContext, sceneModule *wfengine.SceneModule) (int64, string, error) {
	slotId := int(crc32.ChecksumIEEE([]byte(sc.UserId)))
	if slotId < 0 {
		slotId = -1 * slotId
	}
	slotId = slotId % 100
	//TODO 此处仅演示如何根据userid进行分桶，使用方需根据实际业务场景来实现具体的分桶逻辑。
	var groupName string
	if slotId < 95 {
		groupName = "control_group"
	}else if slotId<96 {
		groupName = "treatment_group1"
	}else if slotId<97{
		groupName = "treatment_group2"
	}else {
		groupName = "treatment_group3"
	}

	wfId, ok := sceneModule.GroupWorkflowMap[groupName]
	if !ok{
		return -1, groupName, errors.New(fmt.Sprintf("no matched workflow_id, slotId=%v, groupName=%v", slotId, groupName))
	}
	return wfId, strconv.Itoa(slotId), nil
}
