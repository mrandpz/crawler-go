package engine

import (
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
		parseResult, e := Worker(r)
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
