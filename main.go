package main

import (
	"awesomeProject/crawler/zhenai/engine"
	"awesomeProject/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		//Url:        "http://www.zhenai.com/zhenghun",
		// ParserFunc: parser.ParseCityList,

		Url:        "http://album.zhenai.com/u/70034915",
		ParserFunc: parser.ParseProfile,
	})

}
