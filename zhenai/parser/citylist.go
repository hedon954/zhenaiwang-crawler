package parser

import (
	"fmt"
	"regexp"
	"zhenaiwang-crawler/engine"
)

/**
珍爱网的城市列表解析器
*/

//<a href="http://www.zhenai.com/zhenghun/aba" data-v-1573aa7c>阿坝</a>
const cityListRegex string = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

//解析城市列表
//解析结果：城市名称 + 递归出来的请求
func ParseCityList(content []byte) engine.ParseResult {

	//提取出"url"和"city-name"
	re := regexp.MustCompile(cityListRegex)
	//示例：
	//[
	// <a href="http://www.zhenai.com/zhenghun/zunyi" data-v-1573aa7c>遵义</a>
	// http://www.zhenai.com/zhenghun/zunyi
	// 遵义
	//]
	matches := re.FindAllSubmatch(content, -1)
	//封装数据到 ParseResult
	result := engine.ParseResult{}
	for _, m := range matches {
		//城市的名字
		//result.Items = append(result.Items, "City " + string(m[2]))
		//后面需要做的请求
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				//城市列表下每一个链接需要做 ParseCity 操作
				ParserFunc: ParseCity,
			})
		fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
	}
	return result
}
