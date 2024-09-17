package cron

import (

"context"
"github.com/didi/tg-example/logic/dispatcher"
)

func StartCronTask() {
	//定时任务，可在此添加
	go dispatcher.StartWorkflowEngine(context.TODO())

}
