package main

import (
	"awesomeProject/crawler/crawler_distributed/rpcsupport"
	worker "awesomeProject/crawler/crawler_distributed/woker"
	"flag"
	"fmt"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

//  go run worker.go --port=9000 执行当前文件
func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", *port),
		worker.CrawlService{}))
}
