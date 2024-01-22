package engine

import (
	"log"
	"reptile/fetcher"
)

type SimpleEngin struct {
}

func (e SimpleEngin) Run(seeds ...Request) { //表示可以接受多个Requestd的参数
	//elastic
	//elastic.NewClient()
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:] // 取一个删一个
		//log.Printf("Ferthing %s", r.Url)
		//body, err := fetcher.Fecth(r.Url) // 取出获取到的内容
		//if err != nil {
		//	log.Printf("Fetcher:error ferther url %s: %v", r.Url, err)
		//}
		//parserResult := r.ParserFun(body) // 掉用自己相对应的方法进行解析
		parserResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parserResult.Requests...)
		for _, item := range parserResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fecth(r.Url) // 取出获取到的内容
	if err != nil {
		log.Printf("Fetcher:error ferther url %s: %v", r.Url, err)
		return ParseResult{}, err

	}
	return r.ParserFun(body, r.Url), nil // 掉用自己相对应的方法进行解析

}
