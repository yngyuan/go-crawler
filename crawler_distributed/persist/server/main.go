package main

import (
	"fmt"
	"github.com/olivere/elastic"
	"go-crawler/crawler_distributed/config"
	"go-crawler/crawler_distributed/persist"
	"go-crawler/crawler_distributed/rpcsupport"
	"log"
)

func main() {
	log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index: index,
		})
}