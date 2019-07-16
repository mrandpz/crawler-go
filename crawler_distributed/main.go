package main

import (
	config2 "awesomeProject/crawler/crawler_distributed/config"
	client2 "awesomeProject/crawler/crawler_distributed/persist/client"
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/scheduler"
	"awesomeProject/crawler/zhenai/parser"
	"fmt"
)

func main() {
	//engine.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,

	//Url:        "http://www.zhenai.com/zhenghun/foshan",
	//ParserFunc: parser.ParseCity,

	//Url:        "http://album.zhenai.com/u/70034915",
	//ParserFunc: parser.ParseProfile,
	//})

	itemChan, err := client2.ItemSaver(fmt.Sprintf(":%d", config2.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	var e = engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
