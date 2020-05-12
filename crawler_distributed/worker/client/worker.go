package client

import (
	"go-crawler/crawler/engine"
	"go-crawler/crawler_distributed/config"
	"go-crawler/crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) (engine.Processor) {

	return func(req engine.Request) (engine.ParseResult, error) {

		sReq := worker.SerializeRequest(req)

		var sResult worker.ParseResult
		cli := <- clientChan
		err := cli.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err !=nil {
			return engine.ParseResult{}, err
		}

		return worker.DeserializeResult(sResult), nil
	}
}
