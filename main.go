package main

import (
	"zhenaiwang-crawler/engine"
	"zhenaiwang-crawler/scheduler"
	"zhenaiwang-crawler/zhenai/parser"
)

const zhenAiUrl string = "https://www.zhenai.com/zhenghun"

const shanghaiUrl string = "https://www.zhenai.com/zhenghun/shanghai"

func main() {

	//concurrentEngine := engine.ConcurrentEngine{
	//	Scheduler: &scheduler.QueuedScheduler{},
	//	WorkerCount: 10,
	//}

	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}

	concurrentEngine.Run(engine.Request{
		Url:        zhenAiUrl,
		ParserFunc: parser.ParseCityList,
	})

	//concurrentEngine.Run(engine.Request{
	//	Url:		 shanghaiUrl,
	//	ParserFunc:  parser.ParseCity,
	//})
}
