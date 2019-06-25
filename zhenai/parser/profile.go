package parser

import (
	"awesomeProject/crawler/zhenai/engine"
	"fmt"
	"regexp"
)

const ProfileRe = `<div class="m-btn purple" [^>]*>([^<]+)</div>`

func ParseProfile(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(ProfileRe)
	all := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, value := range all {
		fmt.Printf("%s", value[1])
	}
	return result
}
