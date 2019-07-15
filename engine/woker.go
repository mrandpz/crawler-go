package engine

import (
	"awesomeProject/crawler/fetcher"
	"log"
)

// type Request struct {
//	Url        string
//	ParserFunc func([]byte) ParseResult
//}

//type ParseResult struct {
//	Requests []Request
//	Items    []interface{}
//}
func worker(r Request) (ParseResult, error) {
	// 发起请求，获得返回的内容
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error"+"Fetchering url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	// 解析body返回parseResult
	return r.ParserFunc(body, r.Url), nil
}
