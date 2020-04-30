package persist

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"go-crawler/crawler/model"
	"testing"
)

func TestSave(t *testing.T) {
	expected := model.Profile{
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
	}

	id, err := save(expected)
	if err != nil{
		panic(err)
	}

	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err!= nil {
		panic(err)
	}

	// TODO: Try to start slastic search here using docker go client
	resp, err := client.Get().
		Index("dating_profile").
		Type("zhenai").
		Id(id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)
	var actual model.Profile
	err = json.Unmarshal(*resp.Source, &actual)

	if err != nil {
		panic(err)
	}

	if actual != expected {
		t.Errorf("got %v; expected %v",
			actual, expected)
	}
}
