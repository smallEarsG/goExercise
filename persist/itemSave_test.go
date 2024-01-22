package persist

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"net/http"
	"reptile/engine"
	"reptile/model"
	"testing"
)

func TestSaver(t *testing.T) {
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

	url := "https://127.0.0.1:9200"
	// Elasticsearch 集群用户名和密码（如果启用了基本身份验证）
	username := "elastic"
	password := "S1DpTCl3Ml5bKhoW1eik"
	//TODO: try to start up elastic search
	//here using docker go client
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetBasicAuth(username, password),
		elastic.SetHealthcheck(false), // 禁用健康检查
		elastic.SetScheme("https"),    // 使用 HTTPS 协议
		elastic.SetHttpClient(&http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}), // 跳过证书验证（仅用于测试）
		// 客户端维护集群的状态
		elastic.SetSniff(false))
	const index = "dating_test"
	err = Save(client, expected, index)
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().Index("dating_profile").Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", resp.Source)
	var actual engine.Item
	err = json.Unmarshal(resp.Source, &actual)
	if err != nil {
		panic(err)
	}
	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	if actual != expected {
		log.Println("数据有问题")
	}
}
