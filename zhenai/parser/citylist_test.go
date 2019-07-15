package parser

import (
	"awesomeProject/crawler/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, e := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	if e != nil {
		panic(e)
	}

	result := ParseCityList(contents, "http://www.zhenai.com/zhenghun")

	const resultSize = 470

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d but had %d", resultSize, len(result.Items))
	}
}
