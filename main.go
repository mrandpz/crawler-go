package main

import (
	"awesomeProject/crawler/zhenai/engine"
	"awesomeProject/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		//Url:        "http://www.zhenai.com/zhenghun",
		//ParserFunc: parser.ParseCityList,

		Url:        "http://www.zhenai.com/zhenghun/foshan",
		ParserFunc: parser.ParseCity,

		//Url:        "http://album.zhenai.com/u/70034915",
		//ParserFunc: parser.ParseProfile,
	})

}
