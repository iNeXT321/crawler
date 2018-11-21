package engine

type Concurrent struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item
	RequestProcessor Processor
}

type Processor func(Request) (ParseResult, error)

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerInChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *Concurrent) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()
	//创建worker
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerInChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		//打印items
		for _, item := range result.Items {
			if item.PayLoad == nil {
				continue
			}
			go func() { e.ItemChan <- item }()
		}
		//把Request给scheduler
		for _, r := range result.Requests {
			visitedUrl := isDuplicate(r.Url)
			if visitedUrl {
				continue
			}

			e.Scheduler.Submit(r)
		}
	}
}

func (e *Concurrent) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true
	return false
}
