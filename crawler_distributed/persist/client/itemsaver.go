package client

import (
	config2 "awesomeProject/crawler/crawler_distributed/config"
	rpcsupport2 "awesomeProject/crawler/crawler_distributed/rpcsupport"
	"awesomeProject/crawler/engine"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	// 连接RPC
	client, e := rpcsupport2.NewClient(host)
	if e != nil {
		return nil, e
	}
	out := make(chan engine.Item)

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item : #%d,%s", itemCount, item)
			itemCount++

			// Call RPC TO save item
			result := ""
			err := client.Call(config2.ItemSaverRpc, item, &result)

			if err != nil {
				log.Printf("Item saver:error %v: %v", item, err)
			}

		}
	}()
	return out, nil
}
