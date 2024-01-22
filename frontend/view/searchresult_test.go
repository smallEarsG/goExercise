package view

import (
	"os"
	"reptile/engine"
	"reptile/frontend/model"
	common "reptile/model"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("template.html")
	out, err := os.Create("template.test.html")
	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url: "http",
		Id:  "108906739",
		Payload: common.Profile{
			Name:       "test",
			Gender:     "女",
			Age:        0,
			Height:     0,
			Weight:     0,
			Income:     "财务自由",
			Marriage:   "离异",
			Education:  "初中",
			Occupation: "金融",
			Hokou:      "南京市",
			Xinzuo:     "狮子座",
			House:      "无房",
			Car:        "无车",
		},
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}
	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}
