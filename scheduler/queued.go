package scheduler

import "reptile/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request //用来接收 request
	// 每个worker / request有自己的chan
	workerChan chan chan engine.Request //worker chan 只要看worker 对位的类型  chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	//TODO implement me
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	// 告诉外界我有一个workerChan可以接收 request
	s.workerChan <- w
}

//	func (s *QueuedScheduler) ConfigureMastWorkerChan(c chan engine.Request) {
//		//TODO implement me
//		panic("implement me")
//	}
func (s *QueuedScheduler) Run() {
	// RUN 方法用来总控
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() { // 开协程的作用是为了独立运行且持续运行
		var requestQ []engine.Request
		// workerQ 控制request的派发
		var workerQ []chan engine.Request // 队列里存储的是发送request的chan
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			// len(requestQ) > 0 表示队列中有需要处理的请求
			if len(requestQ) > 0 && len(workerQ) > 0 { //len(workerQ) > 0 表示有空闲的workerChan 可以工作
				activeRequest = requestQ[0] // engine.Request
				activeWorker = workerQ[0]   //chan engine.Request
				// 在这里发送channel如果卡住可能会导致 使select中的requestChan 或者 workerChan发不进来
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest: // 当拿到队列中的一个请求 和 第一个空闲的workerChan时 chan engine.Request  将  Request 发给worker
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}

	}()
}
