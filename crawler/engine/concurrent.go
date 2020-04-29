package engine

import (
	"log"
)

type ConcurrentEngine struct{
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
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

func (e  *ConcurrentEngine) Run(seeds...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i :=0; i<e.WorkerCount;  i++{
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url){
			log.Printf("Duplicate Request: %s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}

	for{
		result := <- out
		for _, item := range result.Items {
			go func() {
				e.ItemChan <- item
			}()
		}

		//URL Redup
		for _, request := range result.Requests {
			if isDuplicate(request.Url){
				log.Printf("Duplicate Request: %s", request.Url)
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, notifier ReadyNotifier) {
	go func() {
		for {
			// tell scheduler I'm ready
			notifier.WorkerReady(in)
			request := <- in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)
func isDuplicate(url string) bool {
	if(visitedUrls[url]) {
		return true
	}
	visitedUrls[url] = true
	return false
}
