package main

import (
	"zhenaiwang-crawler/engine"
	"zhenaiwang-crawler/persist"
	"zhenaiwang-crawler/scheduler"
	"zhenaiwang-crawler/zhenai/parser"
)

const zhenAiUrl string = "https://www.zhenai.com/zhenghun"
const shanghaiUrl string = "https://www.zhenai.com/zhenghun/shanghai"
const esIndex string = "dating_profile"
const esType string = "zhenai"

func main() {

	saver, err := persist.ItemSaver(esIndex, esType)
	if err != nil {
		panic(err)
	}

	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    saver,
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
