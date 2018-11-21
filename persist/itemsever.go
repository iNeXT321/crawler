package persist

import (
	"context"
	"github.com/pkg/errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"test/crawler/engine"
)

func ItemSever(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver: got item #%d %v", itemCount, item)
			err := Save(client, index, item)
			if err != nil {
				log.Printf("Item Saver : error saving item %v :%v", item, err)
			}
			itemCount++
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) (err error) {

	if item.Type == "" {
		return errors.New("Must supply Type")
	}
	indexSever := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexSever.Id(item.Id)
	}
	_, err = indexSever.Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
