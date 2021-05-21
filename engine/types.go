package engine

/**
  定义一些封装类
 */

//请求
type Request struct {
	Url string
	ParserFunc func([]byte) ParseResult
}

//解析结果
type ParseResult struct {
	Requests []Request
	Items []interface{}
}

func NilParser([]byte) ParseResult{
	return ParseResult{}
}