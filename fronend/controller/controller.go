package controller

import (
	"awesomeProject/crawler/engine"
	model2 "awesomeProject/crawler/fronend/model"
	"awesomeProject/crawler/fronend/view"
	"context"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/olivere/elastic"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, e := elastic.NewClient(
		elastic.SetSniff(false))
	if e != nil {
		panic(e)
	}

	return SearchResultHandler{
		view:   view.CreatSearchResultView(template),
		client: client,
	}
}

// localhost:8888/search?q=男
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))

	from, err := strconv.Atoi(req.FormValue("from"))

	if err != nil {
		from = 0
	}

	page, err := h.getSearchResult(q, from)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	e := h.view.Render(w, page)

	if e != nil {
		fmt.Println("ERROR")
		http.Error(w, e.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(q string, from int) (model2.SearchResult, error) {
	var result model2.SearchResult
	result.Query = q
	resp, err := h.client.Search().
		Index("dating_profile1").
		Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).
		From(from).Size(10).
		Do(context.Background())

	if err != nil {
		fmt.Println("err", err)
		return result, err
	}

	result.Hits = resp.TotalHits() //命中个数
	result.Start = from            //第几条数据开始

	//使用Each函数,遍历匹配的数据
	//使用反射机制,解析成engine.Item{}格式
	result.Items = resp.Each(
		reflect.TypeOf(engine.Item{}))
	//fmt.Println(result.Hits)

	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)
	return result, nil
}

// Payload.Single:(已购房) 解决Payload字段太长 直接用 Single:(已购房搜索)
func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}
