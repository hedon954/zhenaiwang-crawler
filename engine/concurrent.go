package engine

import (
	"log"
	"zhenaiwang-crawler/model"
)

/**
并发版爬虫引擎
*/

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier //接口组合
	Submit(Request)
	WorkChan() chan Request //问 Scheduler：我有一个 Worker，请问要给我哪个 channel
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

//并发执行爬虫
func (e *ConcurrentEngine) Run(seeds ...Request) {

	//创建输入通道
	//in := make(chan Request)
	e.Scheduler.Run()

	//由 Scheduler 调度 Worker 去执行 Fetch 和 Parse， 输出 ParseResult
	out := make(chan ParseResult)

	//创建 Worker 去执行爬虫操作
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkChan(), out, e.Scheduler)
	}

	//提交任务
	for _, r := range seeds {
		//对 URL 进行去重
		if isDuplicate(r) {
			log.Printf("Duplicate request: %s\n", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}

	//处理结果
	itemIndex := 1
	for {
		//从 channel 拿到 Worker 处理的结果
		result := <-out
		//输出爬取到的信息
		for _, item := range result.Items {
			switch item.(type) {
			case model.Profile:
				log.Printf("=========== Got item %d ===========\n", itemIndex)
				profile := item.(model.Profile)
				profile.PrintProfile()
			default:
				log.Printf("Got item %d: %v\n", itemIndex, item)
			}
			itemIndex++
		}

		//继续处理新的请求
		for _, request := range result.Requests {
			//对 URL 进行去重
			if isDuplicate(request) {
				//log.Printf("Duplicate request: %s\n", request.Url)
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

//判断 Request 是否重复
var visitedUrls = make(map[string]bool)

func isDuplicate(request Request) bool {
	if visitedUrls[request.Url] {
		return true
	} else {
		visitedUrls[request.Url] = true
		return false
	}
}

//创建 Worker
//输入一个 Request
//由 Scheduler 协调 Request 和 Worker 之间的联系
//交给 Worker 去 Fetch 和 Parse，最终生成一个 ParseResult
func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	//每个 worker 有一个自己的 channel
	go func() {
		for {
			//在将 in 传进 channel 之前要告诉它我已经准备好了
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
