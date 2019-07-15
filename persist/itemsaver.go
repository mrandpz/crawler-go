package persist

import (
	"awesomeProject/crawler/engine"
	"context"
	"log"

	"github.com/olivere/elastic"
)

func ItemSaver(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)

	client, e := elastic.NewClient(
		// must turn off sniff in docker
		elastic.SetSniff(false))

	if e != nil {
		return nil, e
	}

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item : #%d,%s", itemCount, item)
			itemCount++
			err := save(client, index, item)
			if err != nil {
				log.Printf("Item saver:error %v: %v", item, err)
			}

		}
	}()
	return out, nil
}

func save(client *elastic.Client, index string, item engine.Item) (err error) {

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id == "" {
		indexService.Id(item.Id)
	}

	_, err = indexService.
		Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
