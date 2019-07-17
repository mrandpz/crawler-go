package parser

import (
	"awesomeProject/crawler/crawler_distributed/config"
	"awesomeProject/crawler/engine"
	"fmt"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// 获取城市列表
func ParseCityList(contents []byte, _ string) engine.ParseResult {
	fmt.Println("Re")
	re := regexp.MustCompile(cityListRe)
	all := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	fmt.Println("city")

	for _, m := range all {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}
	return result
}
