package scheduler

import "test/crawler/engine"

/*
所有worker共用一个“in”和“out”通道
*/
type SimpleScheduler struct{
	workerInChan chan engine.Request
}

func (s *SimpleScheduler) WorkerInChan() chan engine.Request {
	return s.workerInChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimpleScheduler) Run() {
	s.workerInChan = make(chan engine.Request)
}


func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {s.workerInChan <- request}()
}




