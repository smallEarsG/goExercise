package engine

import (
	"log"
)

type ConcurrentEngin struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	//ConfigureMastWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngin) Run(seeds ...Request) {

	// 创建chennl通道
	//in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler) // 相当于开了几个携程 也就goroutine
	}
	for _, r := range seeds {
		if isDuplicante(r.Url) {
			log.Printf("Duplicate request %s", r.Url)
			continue //跳出循环  去下一个循环
		}
		e.Scheduler.Submit(r)
	}
	//itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			//log.Printf("#%d Got item: %v ", itemCount, item)
			//itemCount++
			go func() { e.ItemChan <- item }()

		}
		// url dedup
		for _, request := range result.Requests {
			if isDuplicante(request.Url) {
				log.Printf("Duplicate request %s", request.Url)
				continue //跳出循环  去下一个循环
			}
			e.Scheduler.Submit(request) // qk
		}
	}
}

var visitedUrls = make(map[string]bool)

func isDuplicante(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	//in := make(chan Request)
	go func() {
		for {

			// tell scheduler i`m ready
			// request : <-in  // <- in 从那里来 ？ 是表示scheduler中的workerChan选择了你这个worker 才会给你发送数据
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
