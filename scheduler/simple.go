package scheduler

import "reptile/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(requests chan engine.Request) {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) ConfigureMastWorkerChan(c chan engine.Request) {
	s.workerChan = c
}
func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.workerChan <- r }()
}
