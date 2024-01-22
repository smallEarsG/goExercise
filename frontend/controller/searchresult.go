package controller

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/olivere/elastic/v7"
	"net/http"
	"reflect"
	"regexp"
	"reptile/engine"
	"reptile/frontend/model"
	"reptile/frontend/view"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(
	template string) SearchResultHandler {
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

	return SearchResultHandler{
		view: view.CreateSearchResultView(
			template),
		client: client,
	}
}

func (h SearchResultHandler) ServeHTTP(
	w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))

	from, err := strconv.Atoi(
		req.FormValue("from"))
	if err != nil {
		from = 0
	}
	fmt.Fprintf(w, "q=%s,from=%d", q, from)
	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusBadRequest)
		return
	}

	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusBadRequest)
		return
	}
}

const pageSize = 10

func (h SearchResultHandler) getSearchResult(
	q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	//result.Query = q

	resp, err := h.client.
		Search("dating_profile").              //config.ElasticIndex
		Query(elastic.NewQueryStringQuery(q)). //rewriteQueryString(q)
		From(from).
		Do(context.Background())

	if err != nil {
		return result, err
	}

	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(
		reflect.TypeOf(engine.Item{}))
	if result.Start == 0 {
		result.PrevFrom = -1
	} else {
		result.PrevFrom =
			(result.Start - 1) /
				pageSize * pageSize
	}
	result.NextFrom =
		result.Start + len(result.Items)

	return result, nil
}

// Rewrites query string. Replaces field names
// like "Age" to "Payload.Age"
func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}
