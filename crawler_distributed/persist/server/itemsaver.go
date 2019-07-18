package main

import (
	config2 "awesomeProject/crawler/crawler_distributed/config"
	persist2 "awesomeProject/crawler/crawler_distributed/persist"
	rpcsupport2 "awesomeProject/crawler/crawler_distributed/rpcsupport"
	"flag"
	"fmt"
	"log"

	"github.com/olivere/elastic"
)

var port = flag.Int("port", 0, "the port for me to listen on")

//  go run worker.go --port=9000 执行当前文件
func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config2.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		return err
	}

	fmt.Println("itemsaver")
	// 创建一个RPC
	return rpcsupport2.ServeRpc(host,
		&persist2.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
