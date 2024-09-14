package cron

import "context"

const (
	//2分钟，定时任务周期
	BASE_TASK_TIME = "0 0/2 * * * ?"
)

func StartCronTask() {
	go StartWorkflowEngine(context.TODO())
	//如有其它定时任务，可在此添加
}
