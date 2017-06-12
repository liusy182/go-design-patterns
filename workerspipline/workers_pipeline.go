package worker

import (
	"fmt"
	"log"
	"sync"
)

type Request struct {
	Data    interface{}
	Handler RequestHandler
}

type RequestHandler func(interface{})

func NewStringRequest(s string, id int, wg *sync.WaitGroup) Request {
	return Request{
		Data: "Hellow",
		Handler: func(i interface{}) {
			defer wg.Done()
			s, ok := i.(string)
			if !ok {
				log.Fatal("expecting a string")
			}
			fmt.Println(s)
		},
	}
}

func main() {
	bufferSize := 100
	var dispatcher Dispatcher = NewDispatcher(bufferSize)
	workers := 3
	for i := 0; i < workers; i++ {
		w := &PreffixSuffixWorker{
			prefixS: fmt.Sprintf("worker id: %d -> ", i),
			suffixS: " World",
			id:      i,
		}
		dispatcher.LaunchWorker(w)
	}

	requests := 10
	var wg sync.WaitGroup
	wg.Add(requests)
	for i := 0; i < requests; i++ {
		req := NewStringRequest("(Msg_id: %d) -> Hello", i, &wg)
		dispatcher.MakeRequest(req)
	}

	dispatcher.Stop()

	wg.Wait()
}
