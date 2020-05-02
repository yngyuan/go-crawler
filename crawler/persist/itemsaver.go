package persist

import (
	"context"
	"errors"
	"github.com/olivere/elastic"
	"go-crawler/crawler/engine"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		// must turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil {
		return nil,err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for{
			item:= <- out
			log.Printf("Item saver: got item #%d: %v", itemCount, item)
			itemCount++
			
			err := Save(client, index, item)
			if err != nil{
				log.Printf("Item saver: error saving item %v: %v", item, err)
			}

		}
	}()
	return out, err
}

func Save(client *elastic.Client, index string,
	item engine.Item) (err error) {
	if item.Type == ""{
		return errors.New("must supply Type")
	}

	indexService := client.Index(). //database name
		Index(index).
		Type(item.Type).  // data sheet name
		Id(item.Id).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err = indexService.
		Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
