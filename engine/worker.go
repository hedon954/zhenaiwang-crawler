package engine

import (
	"log"
	"zhenaiwang-crawler/fetcher"
)

//将 Parser 和 Fetcher 合二为一
//输入 Request 进行 Fetch，并将 Parser 的结果返回出去
func worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		//出错就忽略这个请求
		log.Printf("Fetcher: error fetching url %s: %v\n", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
