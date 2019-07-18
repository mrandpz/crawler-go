package client

import (
	"awesomeProject/crawler/crawler_distributed/config"
	worker "awesomeProject/crawler/crawler_distributed/woker"
	"awesomeProject/crawler/engine"
	"net/rpc"
)

func CreateProcessor(
	clientChan chan *rpc.Client) engine.Processor {

	return func(
		req engine.Request) (
		engine.ParseResult, error) {

		sReq := worker.SerializeRequest(req)

		var sResult worker.ParseResult

		c := <-clientChan
		e := c.Call(config.CrawlServiceRpc, sReq, &sResult)

		if e != nil {
			return engine.ParseResult{}, e
		}

		return worker.DeserializeResult(sResult), nil
	}
}
