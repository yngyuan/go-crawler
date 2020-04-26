package main

import (
	"go-crawler/crawler/engine"
	"go-crawler/crawler/zhenai/parser"
)

const url = "http://localhost:8080/mock/www.zhenai.com/zhenghun"

func main()  {
	engine.Run(engine.Request{
		Url: url,
		ParserFunc: parser.ParseCityList,
	})
}

