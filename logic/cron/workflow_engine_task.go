package cron

import (
"context"
"fmt"
"github.com/didi/tg-example/global/constants"
"github.com/didi/tg-example/logic/module"
"github.com/didi/tg-flow/common/tlog"
"github.com/didi/tg-flow/common/utils"
"github.com/didi/tg-flow/wfengine"
"time"
)

var WorkflowEngine *wfengine.WorkflowEngine

const (
	workflowPath  = "conf/workflow"
    DLTagCronTask = "crontab_task"
)

func StartWorkflowEngine(ctx context.Context){
	for{
		defer utils.Recover(ctx, "init_workflow_engine_err")
		InitWorkflowEngine(ctx)

		time.Sleep(time.Duration(2* time.Minute))
	}
}

func InitWorkflowEngine(ctx context.Context) {
	moduleObj := module.ModuleObject{}

	onlineVersion := "0"
	if WorkflowEngine != nil {
		onlineVersion = WorkflowEngine.GetVersion()
	}

	//第1种初始化方式：load from redis
	//latestVersion, err := GetLatestVersionFromRedis(constants.CurrentAppId)
	//if err != nil || latestVersion == "" {
	//	tlog.Handler.ErrorCount(ctx, "cron_task_workflow_err", fmt.Sprintf("get latest version from redis fail, appId=%v, err=%v", constants.CurrentAppId, err))
	//	return
	//}
	//if onlineVersion == latestVersion {
	//	tlog.Handler.Infof(ctx, "cron_task_workflow", "Current workflow engine version is the latest version:%v, need not update!", latestVersion)
	//	return
	//}
	//workflowEngine, err := NewWorkflowEngineFromRedis(moduleObj, constants.CurrentAppId)
	//if err != nil {
	//	tlog.Handler.ErrorCount(ctx, "cron_task_workflow_err", fmt.Sprintf("workflow engine init fail, AppId=%v, error=%v", constants.CurrentAppId, err))
	//	return
	//}

	//第2种初始化方式：load from file
	latestVersion, err := wfengine.GetLatestVersionFromFile(workflowPath)
	if err != nil || latestVersion == "" {
		tlog.Handler.ErrorCount(ctx, "cron_task_workflow_err", fmt.Sprintf("get latest version from path:%v fail, err=%v", workflowPath, err))
		return
	}
	if onlineVersion == latestVersion {
		tlog.Handler.Infof(ctx, "cron_task_workflow_err", fmt.Sprintf("Current workflow engine version is the latest version:%v, need not update!", latestVersion))
		return
	}
	workflowEngine, err := wfengine.NewWorkflowEngineFromFile(moduleObj, workflowPath)
	if err != nil {
		tlog.Handler.ErrorCount(ctx, "cron_task_workflow_err", fmt.Sprintf("workflow engine init fail, AppId=%v, error=%v", constants.CurrentAppId, err))
		return
	}

	//scene.json中的flow_type属性，取值有2种：
	//flow_type = 0, 内置随机分桶，根据用户id来分
	//flow_type = 1，需自定义FlowSelector
	//如果有flow_type=1的情况，需自行实现FlowSelector;如没有，则不需实现
	workflowEngine.SetCustomFlowSelector(&CustomSelector{})
	WorkflowEngine = workflowEngine
	tlog.Handler.Infof(ctx, DLTagCronTask, fmt.Sprintf("workflow engine update successful, appId=%v, workflowEngine=%v", constants.CurrentAppId, WorkflowEngine))
}
