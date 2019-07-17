package main

import (
	config2 "awesomeProject/crawler/crawler_distributed/config"
	persist2 "awesomeProject/crawler/crawler_distributed/persist"
	rpcsupport2 "awesomeProject/crawler/crawler_distributed/rpcsupport"
	"fmt"
	"log"

	"github.com/olivere/elastic"
)

func main() {
	log.Fatal(serveRpc(fmt.Sprintf(":%d", config2.ItemSaverPort), config2.ElasticIndex))

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
