package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)
	//初始化调度器chan
	e.Scheduler.ConfigureMasterWorkerChan(in)

	//创建多个worker
	for i := 0; i <= e.WorkerCount; i++ {
		createWorker(in, out)
	}
	//吧所有的请求提交到chan中
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item: &%d,%v", itemCount, item)
			itemCount++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

/**
每个worker开启一个协程 等待chan的数据返回
 */
func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()

}
