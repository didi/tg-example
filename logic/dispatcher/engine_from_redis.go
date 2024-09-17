/**
Description : loader of workflow config info from redis
Author		: dayunzhangyunfeng@didiglobal.com
Date		: 2021-07-12
*/

package dispatcher

import (

"context"
"fmt"
"github.com/didi/tg-example/common/redis"
"github.com/didi/tg-flow/common/tlog"
"github.com/didi/tg-flow/wfengine"
)

const (
	RedisKeySceneModule = "scene_module_app_"
	RedisKeyWorkflow    = "workflow_app_"
	RedisKeyVersion     = "version_app_"
)

func GetLatestVersionFromRedis(appId int64) (string, error) {
	return redis.Handler.Get(context.TODO(), fmt.Sprintf("%v%v", RedisKeyVersion, appId))
}

func NewWorkflowEngineFromRedis(moduleObj wfengine.ModuleObjBase, appId int64) (*wfengine.WorkflowEngine, error) {
	//更新系统下全部场景的节点对象
	sceneModuleMapString, err := redis.Handler.Get(context.TODO(), fmt.Sprintf("%v%v", RedisKeySceneModule, appId))
	if err != nil {
		return nil, err
	}

	workflowMapStr, err := redis.Handler.Get(context.TODO(), fmt.Sprintf("%v%v", RedisKeyWorkflow, appId))
	if err != nil {
		return nil, err
	}

	//版本号不是必须，兼容吧
	version, err1 := GetLatestVersionFromRedis(appId)
	if err1 != nil {
		tlog.Handler.ErrorCount(context.TODO(), "GetLatestVersionFromRedis_err", fmt.Sprintf("appId=%v, err=%v", appId, err1))
	}

	return wfengine.NewWorkflowEngineFromKV(moduleObj, sceneModuleMapString, workflowMapStr, version)
}
