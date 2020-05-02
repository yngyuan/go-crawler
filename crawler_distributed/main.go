package main

import (
	"fmt"
	"go-crawler/crawler/engine"
	"go-crawler/crawler/scheduler"
	"go-crawler/crawler/zhenai/parser"
	"go-crawler/crawler_distributed/config"
	"go-crawler/crawler_distributed/persist/client"
)

const url = "http://localhost:8080/mock/www.zhenai.com/zhenghun"

func main()  {
	itemChan, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		// scheduler.SimpleScheduler
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})
}

