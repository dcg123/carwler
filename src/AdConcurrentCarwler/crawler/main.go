package main

import (
	"AdConcurrentCarwler/engine"
	"AdConcurrentCarwler/zhenai/parser"
	"AdConcurrentCarwler/scheduler"
)


func main() {
	e:=engine.ConcurrentEngine{
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:10,
	}
	//e:=engine.ConcurrentEngine{
	//	Scheduler:&scheduler.SimpleScheduler{},
	//	WorkerCount:10,
	//}
	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,
	})
}
