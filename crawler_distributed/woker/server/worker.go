package main

import (
	"awesomeProject/crawler/crawler_distributed/config"
	"awesomeProject/crawler/crawler_distributed/rpcsupport"
	worker "awesomeProject/crawler/crawler_distributed/woker"
	"fmt"
	"log"
)

func main() {
	fmt.Println("worker")
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", config.WorkerPort0),
		worker.CrawlService{}))
}
