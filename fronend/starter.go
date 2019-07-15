package main

import (
	"awesomeProject/crawler/fronend/controller"
	"net/http"
)

func main() {
	//防止CSS等内容没有展示出来
	//因此使用http fileServer提供静态内容
	http.Handle("/", http.FileServer(
		http.Dir("crawler/fronend/view")))

	//当访问到/search,则创建对象,解析模板
	//注意,SearchResultHandler使用了http.Handle必须要实现ServeHTTP()
	http.Handle("/search",
		controller.CreateSearchResultHandler(
			"crawler/fronend/view/template.html"))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
