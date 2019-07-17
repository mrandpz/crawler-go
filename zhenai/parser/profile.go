package parser

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/model"
	"regexp"
	"strings"
)

var singleRe = regexp.MustCompile(`<div class="m-btn [0-9a-z]+" [^>]*>([^<]+)</div>`)

//var singleRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^<]+)</div>`)

//type ParseResult struct {
//	Requests []Request
//	Items    []interface{}
//}
// 解析用户信息
func parseProfile(contents []byte, name string, url string) engine.ParseResult {
	profile := model.Profile{}

	profile.Name = name

	// 找到所有匹配的个人资料
	single := getUser(singleRe, contents)

	profile.Single = single

	var reqs []engine.Request

	str := strings.Replace(url,
		"http://album.zhenai.com/u/", "", -1)

	result := engine.ParseResult{
		Requests: reqs,
		Items: []engine.Item{
			{
				Type:    "zhenai",
				Url:     url,
				Id:      str,
				Payload: profile,
			},
		},
	}

	return result

}

func getUser(re *regexp.Regexp, contents []byte) []string {
	// 找到所有匹配的
	all := re.FindAllSubmatch(contents, -1)

	var data []string
	if all != nil {
		//
		for _, tags := range all {
			allSub := re.FindSubmatch(tags[0])
			data = append(data, string(allSub[1]))
		}
	}
	return data
}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contens []byte, url string) engine.ParseResult {
	return parseProfile(contens, url, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ProfileParser", p.userName
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}
