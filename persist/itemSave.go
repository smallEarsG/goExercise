package persist

import (
	"context"
	"crypto/tls"
	"github.com/olivere/elastic/v7"
	"log"
	"net/http"
	"reptile/engine"
)

func ItemSaver(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	url := "https://127.0.0.1:9200"
	// Elasticsearch 集群用户名和密码（如果启用了基本身份验证）
	username := "elastic"
	password := "3URmDVmBBS14YG5=ObfO"
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetBasicAuth(username, password),
		elastic.SetHealthcheck(false), // 禁用健康检查
		elastic.SetScheme("https"),    // 使用 HTTPS 协议
		elastic.SetHttpClient(&http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}), // 跳过证书验证（仅用于测试）
		// 客户端维护集群的状态
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++
			err := Save(client, item, index) //
			if err != nil {
				log.Printf("Item saver error saving iteme %v:%v", item, err)
			}
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, item engine.Item, index string) error {

	indexService := client.Index().Index(index).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
