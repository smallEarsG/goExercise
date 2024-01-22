package parser

import (
	"io/ioutil"
	"reptile/engine"
	"reptile/model"
	"testing"
)

func TestParseProFile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProFile(contents, "test", "http://localhost:8080/mock/album.zhenai.com/u/108906739")
	if len(result.Items) != 1 {
		t.Errorf("Result should contain 1 element;but was %v", result.Items)
	}
	actual := result.Items[0]
	expected := engine.Item{
		Url: "http",
		Id:  "108906739",
		Payload: model.Profile{
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
	if expected != actual {
		t.Errorf("expected %v ;but was %v", expected, actual)
	}
}
