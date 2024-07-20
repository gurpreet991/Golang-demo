package main

import (
	"container/heap"
	"fmt"
	"sync"
	"time"
	"yunaio-multiplexer-worker/request_stream"
)

func main() {
	totalRequests := request_stream.U * request_stream.RequestsPerUser
	minuteRequestLimit := request_stream.K
	totalRequestsToProcess := totalRequests

	totalRuntime := time.Duration((totalRequests+minuteRequestLimit-1)/minuteRequestLimit*60) * time.Second

	fmt.Printf("Total Requests: %d\n", totalRequests)
	fmt.Printf("Request Limit per Minute: %d\n", minuteRequestLimit)
	fmt.Printf("Total Runtime: %d seconds\n", int(totalRuntime.Seconds()))

	pq := &request_stream.HeapPriorityQueue{}
	heap.Init(pq)

	var wg sync.WaitGroup

	start := time.Now()
	endTime := start.Add(totalRuntime)

	for time.Now().Before(endTime) {

		pendingRequests := totalRequestsToProcess - request_stream.ProcessedRequests
		if pendingRequests <= 0 {
			break
		}

		requestsThisInterval := request_stream.GetMincalculation(request_stream.K, pendingRequests)

		for i := 0; i < requestsThisInterval; i++ {
			wg.Add(1)
			payload := request_stream.RequestPayload{Uuid: fmt.Sprintf("uuid-%d", i)}
			request := request_stream.Request{UserId: fmt.Sprintf("user%d", i), Payload: payload}
			go request_stream.ProcessIncomingRequests(pq, &request, &wg)
		}

		if pendingRequests > request_stream.K {
			time.Sleep(60 * time.Second)
		} else {
			remainingimeInterval := endTime.Sub(time.Now())
			if remainingimeInterval > 0 {
				time.Sleep(remainingimeInterval)
			}
		}

		fmt.Printf("|--------------seconds --------------- 60 seconds|\n")
		fmt.Printf("|              %d Total Processed requests             |\n", request_stream.ProcessedRequests)
		fmt.Printf("|              %d Processed requests             |\n", requestsThisInterval)
		fmt.Printf("|               %d queued               |\n", totalRequests-request_stream.ProcessedRequests)
	}

	wg.Wait()
}
