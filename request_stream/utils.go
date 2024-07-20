package request_stream

import (
	"container/heap"
	"math/rand"
	"sync"
	"time"
)

type RequestPayload struct {
	Uuid string
}

type Request struct {
	UserId  string
	Payload RequestPayload
}

type HeapPriorityQueue []*Request

func (pq HeapPriorityQueue) Len() int           { return len(pq) }
func (pq HeapPriorityQueue) Less(i, j int) bool { return pq[i].Payload.Uuid < pq[j].Payload.Uuid }
func (pq HeapPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *HeapPriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Request))
}
func (pq *HeapPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func GetRandomItemFromServer(payload *RequestPayload) {
	delay := time.Duration(2000+(100*rand.Intn(10))) * time.Millisecond
	time.Sleep(delay)
	// log.Printf("Successfully processed payload with id: %s", payload.Uuid)
}

func ProcessIncomingRequests(pq *HeapPriorityQueue, request *Request, wg *sync.WaitGroup) {
	defer wg.Done()

	Mu.Lock()
	heap.Push(pq, request)
	QueuedRequests++
	Mu.Unlock()

	GetRandomItemFromServer(&request.Payload)

	Mu.Lock()
	QueuedRequests--
	ProcessedRequests++
	Mu.Unlock()
}

func GetMincalculation(a, b int) int {
	if a < b {
		return a
	}
	return b
}
