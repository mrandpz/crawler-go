package persist

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/persist"
	"log"

	"github.com/olivere/elastic"
)

// 提供RPC 给 Itemsaver
type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(
	item engine.Item, result *string) error {

	err := persist.Save(s.Client, s.Index, item)

	log.Printf("Item %v saved.", item)

	if err != nil {
		*result = "ok"
	} else {
		log.Printf("Error saveing item %v:%v", item, err)
	}
	return err
}
