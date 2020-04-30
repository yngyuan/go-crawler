package view

import (
	"go-crawler/crawler/engine"
	"go-crawler/crawler/frontend/model"
	common "go-crawler/crawler/model"
	"os"
	"testing"

)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView(
		"template.html")

	out, err := os.Create("template.test.html")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url: "http://localhost:8080/mock/album.zhenai.com/u/1490323100484795963",
		Type:"zhenai",
		Id: "1490323100484795963",
		Payload: common.Profile{
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
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)
	if err != nil {
		t.Error(err)
	}

	// TODO: verify contents in template.test.html
}
