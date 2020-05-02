package main

import (
	"go-crawler/crawler/engine"
	"go-crawler/crawler/model"
	"go-crawler/crawler_distributed/config"
	"go-crawler/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host =":1234"
	expected := engine.Item{
		Url: "http://localhost:8080/mock/album.zhenai.com/u/1490323100484795963",
		Type:"zhenai",
		Id: "1490323100484795963",
		Payload: model.Profile{
			Age:        82,
			Height:     143,
			Weight:     96,
			Income:     "10001-20000元",
			Gender:     "男",
			Name:       "故我初心",
			Xinzuo:     "狮子座",
			Occupation: "产品经理",
			Marriage:   "离异",
			House:      "有房",
			Hukou:      "上海市",
			Education:  "硕士",
			Car:        "无车",
		},
	}

	// start ItemSaverServer
	go serveRpc(host, "test1")
	time.Sleep(5*time.Second)
	// start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// Call save service
	result := ""
	err = client.Call(config.ItemSaverRpc, expected, &result)
	if err != nil || result != "ok"{
		t.Errorf("result: %s; err %s", result, err)
	}
}
