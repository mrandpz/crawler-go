package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

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
		creatWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range Seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf(
				"GOT ITEM #%d: %v", itemCount, item)
			itemCount++

		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}

}

func creatWorker(in chan Request, out chan ParseResult, r ReadyNotifier) {
	go func() {
		for {
			r.WorkerReady(in)
			request := <-in
			result, e := worker(request)
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
