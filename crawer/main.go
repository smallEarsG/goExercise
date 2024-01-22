package main

import (
	"reptile/engine"
	"reptile/persist"
	"reptile/scheduler"
	"reptile/zhenai/parser"
)

func main() {
	//regexp.Compile("res")
	//resp, err := http.Get("http://localhost:8080/mock/www.zhenai.com/zhenghun")
	//engine.SimpleEngin{}.Run(engine.Request{
	//	Url:       "http://localhost:8080/mock/www.zhenai.com/zhenghun",
	//	ParserFun: parser.ParserCityList,
	//})
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngin{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    itemChan}
	e.Run(engine.Request{
		Url:       "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParserFun: parser.ParserCityList,
	})
}
