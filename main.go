package main

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/persist"
	"awesomeProject/crawler/scheduler"
	"awesomeProject/crawler/zhenai/parser"
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

	itemChan, err := persist.ItemSaver("dating_profile1")
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
