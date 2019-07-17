package parser

import (
	"awesomeProject/crawler/crawler_distributed/config"
	"awesomeProject/crawler/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a></th>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

// 从第一次获取的城市列表中获得的链接，点击进入第一页用户列表，再执行解析用户信息
func ParseCity(contents []byte, _ string) engine.ParseResult {
	// 从正则中获取 信息
	all := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	// 获得解析后的内容 ParseResult
	for _, m := range all {
		name := string(m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: NewProfileParser(name),
		})
	}

	submatch := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range submatch {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}

	return result
}
