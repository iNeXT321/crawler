package main

import (
	"test/crawler/engine"
	"test/crawler/persist"
	"test/crawler/scheduler"
	"test/crawler/zhenai/parser"
	"test/crawler_distributed/config"
)

//数据库名
const index = "dating_profile"

func main() {
	//配置一个Item save 服务
	itemChan, err := persist.ItemSever(index)
	if err != nil {
		panic(err)
	}
	//配置 Scheduler（调度器）、 worker的个数
	e := engine.Concurrent{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}
	//启动程序
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParse(parser.ParseCityList, config.ParseCityList),
	})
}
