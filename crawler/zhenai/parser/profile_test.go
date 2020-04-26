package parser

import (
	"go-crawler/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfilet(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents, "故我初心")

	if len(result.Items) != 1 {
		t.Errorf("result shoud contain 1 element, but was %v", result.Items)
	}

	actual := result.Items[0].(model.Profile)

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

	if actual != expected {
		t.Errorf("expected %v; but got %v", expected, actual)
	}
}
