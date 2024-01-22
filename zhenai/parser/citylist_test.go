package parser

import (
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {
	//contents, err := fetcher.Fecth("http://localhost:8080/mock/www.zhenai.com/zhenghun")
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		//panic(err)
		t.Errorf("%s\n", contents)
	}

	//fmt.Printf("%s\n", contents)
	//fmt.Println("长度/:", len(contents))
	result := ParserCityList(contents)
	const resultSize = 470
	expectedUrls := []string{
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/aba",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/akesu",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/alashanmeng",
		//"阿坝", "", "",
	}
	//expectedCites := []string{
	//	"City阿坝", "City阿克苏", "City阿拉善盟",
	//}
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d request;but had %d", resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s ;but was %s", i, url, result.Requests[i].Url)
		}
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d request;but had %d", resultSize, len(result.Items))
	}

	//for i, city := range expectedCites {
	//
	//	if result.Items[i].(string) != city {
	//		t.Errorf("expected city #%d: %s ; but was %s", i, city, result.Items[i])
	//
	//		//t.Errorf("expected city #%d: %s ;but was %s", i, city, result.Items[i].(string))
	//		//t.Errorf("expected city #%d: %s ;but was ", i, result.Items[i].(string))
	//	}
	//}
}
