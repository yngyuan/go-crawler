package engine

import (
	"log"
)

type ConcurrentEngine struct{
	Scheduler Scheduler
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

func (e  *ConcurrentEngine) Run(seeds...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i :=0; i<e.WorkerCount;  i++{
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	itemCount := 0
	for{
		result := <- out
		for _, item := range result.Items {
			log.Printf("Got item #%d: %v", itemCount,item)
			itemCount++
		}

		for _, request := range result.Requests {
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
