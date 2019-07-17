package engine

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item
	RequestProcessor Processor
}

type Processor func(Request) (ParseResult, error)

type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	ReadyNotifier
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(Seeds ...Request) {

	out := make(chan ParseResult)

	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		e.creatWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range Seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			// 拿到item 就扔掉

			go func() {
				e.ItemChan <- item
			}()
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}

}

func (e *ConcurrentEngine) creatWorker(
	in chan Request, out chan ParseResult, r ReadyNotifier) {
	go func() {
		for {
			r.WorkerReady(in)
			request := <-in
			result, e := e.RequestProcessor(request)
			if e != nil {
				continue
			}
			out <- result
		}
	}()
}

//用于url去重
var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
