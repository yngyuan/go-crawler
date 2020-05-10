package main

import (
	"fmt"
	"go-crawler/crawler_distributed/config"
	"go-crawler/crawler_distributed/rpcsupport"
	"go-crawler/crawler_distributed/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T)  {
	const host = ":9000"
	go rpcsupport.ServeRpc(
		host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://localhost:8080/mock/album.zhenai.com/u/1490323100484795963",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "故我初心",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)

	if err != nil{
		t.Error(err)
	} else {
		fmt.Println(result)
	}

}
