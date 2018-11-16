package scheduler

import "test/crawler/engine"
/*
每个worker分配一"in"通道，但共用一个“out”通道
*/
type QueuedScheduler struct{
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func (s *QueuedScheduler) WorkerInChan() chan engine.Request {
	return  make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request){
	s.workerChan <- w
}

//总控
func (s *QueuedScheduler) Run(){
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		//request队列
		var requestQ []engine.Request
		//worker队列
		var workerQ [] chan engine.Request

		for{
			var activeRequest engine.Request
			var activeWorker chan engine.Request

			if len(requestQ)>0 && len(workerQ)>0{
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <- s.requestChan:
				requestQ = append(requestQ, r)
			case w := <- s.workerChan:
				workerQ = append(workerQ, w)
				//
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
