package main

import (
"context"
"fmt"
"github.com/cihub/seelog"
"github.com/didi/tg-example/controller"
"github.com/didi/tg-example/logic/cron"
"github.com/didi/tg-example/logic/dispatcher"
"github.com/didi/tg-flow/common/tlog"
"github.com/gin-gonic/gin"
"log"
)

func main(){
	// 初始化seelog
	seeLog, err := seelog.LoggerFromConfigAsFile("conf/tlog.xml")
	if err != nil {
		log.Fatalf("Error initializing seelog: %v", err)
	}

	// 为tg-flow提供CountLogger初始化参数
	tlog.InitCountLoggerFromSeelog(seeLog, nil)
	defer seeLog.Flush()

	//如有redis访问的需要（如采用从redis加载工作流配置的方式），可在此处初始化,对应的redis配置文件为:conf/app.conf
	//redis.InitRedisHandlers([]string{"redis"})

	//工作流引擎初始化
	dispatcher.InitWorkflowEngine(context.TODO())

	//定时任务
	cron.StartCronTask()

	r := gin.New()
	es := &controller.ExampleService{}
	r.GET("/api/recommend",  es.Recommend)

	tlog.Handler.Info("tg-example server start finished !!!")

	// 运行应用，监听在端口8080
	if err := r.Run(":8080"); err != nil {
		tlog.Handler.Errorf(context.TODO(), "interface_query", fmt.Sprintf("Error starting server: %v", err))
	}
}
