package persist

import (
	"go-crawler/crawler/model"
	"testing"
)

func TestSave(t *testing.T) {
	profile := model.Profile{
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

	save(profile)
}
