package engine

import (
	"log"
	"zhenaiwang-crawler/fetcher"
	"zhenaiwang-crawler/model"
)

/**
  简单版爬虫引擎
		输入：seeds 种子
 */

type SimpleEngine struct {

}

//不断对输入的请求进行处理
func (e SimpleEngine) Run(seeds ... Request)  {

	//将 seeds 存入一个 slice 里面
	var requests []Request
	for _,r := range seeds{
		requests = append(requests, r)
	}
	//处理请求
	for len(requests) > 0 {
		//取出第一个请求进行处理
		r := requests[0]
		requests = requests[1:]

		//① 进行网页抓取
		parseResult, err := worker(r)
		if err != nil{
			continue
		}
		//② 抓取网页后进行解析
		//②-1：将递归进来的请求加入请求的 slice 中
		requests = append(requests, parseResult.Requests...)
		//②-2：将信息进行处理
		items := parseResult.Items
		for _,item := range items{
			switch item.(type) {
			case model.Profile:
				log.Println("=========== Got item ===========")
				profile := item.(model.Profile)
				profile.PrintProfile()
			default:
				log.Printf("Got item: %v\n", item)
			}
		}
	}
}

//将 Parser 和 Fetcher 合二为一
//输入 Request 进行 Fetch，并将 Parser 的结果返回出去
func worker(r Request) (ParseResult, error){
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		//出错就忽略这个请求
		log.Printf("Fetcher: error fetching url %s: %v\n", r.Url, err)
		return ParseResult{},err
	}
	return r.ParserFunc(body),nil
}

