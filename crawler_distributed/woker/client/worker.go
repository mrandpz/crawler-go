package client

import (
	"awesomeProject/crawler/crawler_distributed/config"
	"awesomeProject/crawler/crawler_distributed/rpcsupport"
	worker "awesomeProject/crawler/crawler_distributed/woker"
	"awesomeProject/crawler/engine"
	"fmt"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(
		fmt.Sprintf(":%d", config.WorkerPort0))

	if err != nil {
		return nil, err
	}

	return func(
		req engine.Request) (
		engine.ParseResult, error) {

		sReq := worker.SerializeRequest(req)

		var sResult worker.ParseResult

		e := client.Call(config.CrawlServiceRpc, sReq, &sResult)

		if e != nil {
			fmt.Println(e)
			return engine.ParseResult{}, err
		}

		return worker.DeserializeResult(sResult), nil
	}, nil
}
