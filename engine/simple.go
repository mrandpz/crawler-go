package engine

import (
	"awesomeProject/crawler/fetcher"
	"log"
)

type SimleEngine struct {
}

func (e SimleEngine) Run(seeds ...Request) {
	var requests []Request
	// 拿到seeds 之后遍历seeds    推到requests中
	for _, r := range seeds {
		requests = append(requests, r)
	}

	//
	for len(requests) > 0 {
		// 拿出requests的第一项递归
		r := requests[0]
		requests = requests[1:]

		// 一个worker发起一个请求返回解析后的数据 parseResult
		parseResult, e := worker(r)
		if e != nil {
			continue
		}
		// 再从parseResult 中拿出Requests的所有request再push到requests的队列中
		requests = append(requests, parseResult.Requests...)

		// 打印出parseResult的所有item信息
		for _, item := range parseResult.Items {
			log.Printf("Got item %s", item)
		}

	}
}

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
	return r.ParserFunc(body), nil
}
