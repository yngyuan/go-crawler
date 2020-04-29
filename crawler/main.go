package main

import (
	"go-crawler/crawler/engine"
	"go-crawler/crawler/scheduler"
	"go-crawler/crawler/zhenai/parser"
)

const url = "http://localhost:8080/mock/www.zhenai.com/zhenghun"

func main()  {
	e := engine.ConcurrentEngine{
		// scheduler.SimpleScheduler
		Scheduler: &scheduler.QueueScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url: url,
		ParserFunc: parser.ParseCityList,
	})
}

