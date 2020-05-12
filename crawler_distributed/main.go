package main

import (
	"flag"
	"go-crawler/crawler/engine"
	"go-crawler/crawler/scheduler"
	"go-crawler/crawler/zhenai/parser"
	"go-crawler/crawler_distributed/config"
	itemsaver "go-crawler/crawler_distributed/persist/client"
	"go-crawler/crawler_distributed/rpcsupport"
	worker "go-crawler/crawler_distributed/worker/client"
	"log"
	"net/rpc"
	"strings"
)

const url = "http://localhost:8080/mock/www.zhenai.com/zhenghun"

var (
	itemSaverHost = flag.String(
		"itemsaver_host", "", "itemsaver host")
	workerHosts = flag.String(
		"worker_hosts", "",
		"worker hosts(comma seperated)")
)

func main()  {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		// scheduler.SimpleScheduler
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:        url,
		Parser: engine.NewFuncParser(
			parser.ParseCityList, config.ParseCityList),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("connected to %s", h)
		} else {
			log.Printf("error connecting to %s: %v", h, err)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for{
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
