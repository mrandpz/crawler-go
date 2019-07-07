package parser

import (
	"awesomeProject/crawler/model"
	"awesomeProject/crawler/zhenai/engine"
	"fmt"
	"regexp"
)

var singleRe = regexp.MustCompile(`<div class="m-btn [0-9a-z]+" [^>]*>([^<]+)</div>`)

//var singleRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^<]+)</div>`)

// 解析用户信息
func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}

	profile.Name = name

	// 找到所有匹配的个人资料
	single := getUser(singleRe, contents)

	profile.Single = single
	fmt.Printf("%s", profile.Name)

	return engine.ParseResult{}

}

func getUser(re *regexp.Regexp, contents []byte) []interface{} {
	// 找到所有匹配的
	all := re.FindAllSubmatch(contents, -1)

	var data []interface{}
	if all != nil {
		//fmt.Printf("%s", all)
		//
		for _, tags := range all {
			//fmt.Printf("tags[0]= %s", tags[0])
			//fmt.Println()
			// 找到匹配的
			allSub := re.FindSubmatch(tags[0])
			//
			data = append(data, allSub[1])
			//fmt.Println()
		}
	}
	return data
}
