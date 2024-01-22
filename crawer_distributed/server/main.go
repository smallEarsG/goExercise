package main

import (
	"crypto/tls"
	"github.com/olivere/elastic/v7"
	"net/http"
	"reptile/crawer_distributed/persist"
	"reptile/crawer_distributed/rpcsupport"
)

func main() {
	url := "https://127.0.0.1:9200"
	// Elasticsearch 集群用户名和密码（如果启用了基本身份验证）
	username := "elastic"
	password := "3URmDVmBBS14YG5=ObfO"
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
	if err != nil {
		panic(err)
	}
	rpcsupport.ServeRpc(":1234", persist.ItemSaverService{
		Client: client,
		Index:  "dating_profile",
	})
}
