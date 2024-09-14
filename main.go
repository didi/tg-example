package main

import (
"context"
"fmt"
"github.com/cihub/seelog"
"github.com/didi/tg-example/controller"
"github.com/didi/tg-example/logic/cron"
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

	//如有redis访问的需要（如从redis加载工作流配置），可在此处初始化
	//redis.InitRedisHandlers([]string{"redis", "redis_native"})

	r := gin.New()

	cron.StartCronTask()
	tlog.Handler.Info("conf.StartCronTask finished")

	es := &controller.ExampleService{}
	r.GET("/search/query",  es.Query)

	// 运行应用，监听在端口8080
	if err := r.Run(":8080"); err != nil {
		tlog.Handler.Errorf(context.TODO(), "interface_query", fmt.Sprintf("Error starting server: %v", err))
	}
}
