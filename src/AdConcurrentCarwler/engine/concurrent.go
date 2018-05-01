package engine

import (
	"log"
)

type ConcurrentEngine struct{
	Scheduler Scheduler
	WorkerCount int
}
type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request)  {
	//吧请求提交到request chan中
	for _,r:=range seeds{
		e.Scheduler.Submit(r)
	}
	//in:=make(chan Request)
	out:=make(chan ParseResult)
	//e.Scheduler.ConfigureMasterWorkerChan(in)
	//分发任务
	e.Scheduler.Run()
	//for i:=0;i<e.WorkerCount;i++{
	//	createWorker(in,out)
	//}
	//创建多个work
	for i:=0;i<e.WorkerCount;i++{
		createWorker(out,e.Scheduler)
	}
	for {
		result:=<-out
		for _,item:=range result.Items{
			log.Printf("Got item: %v",item)
		}
		for _,request:=range result.Requests{
			e.Scheduler.Submit(request)
		}
	}
}
/**
每个work开启一个goroutine
 */
func createWorker(out chan ParseResult,s Scheduler)  {
	in:=make(chan Request)
	go func() {
		for{
			//每个work 都有一个channel
			s.WorkerReady(in)
			request:=<-in
			result,err:=worker(request)
			if err!=nil{
				continue
			}
			out<-result
		}
	}()
}
//func createWorker(in chan Request,out chan ParseResult)  {
//	go func() {
//		for{
//			request:=<-in
//			result,err:=worker(request)
//			if err!=nil{
//				continue
//			}
//			out<-result
//		}
//	}()
//}
