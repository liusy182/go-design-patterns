package worker

import (
	"errors"
	"time"
)

type Dispatcher interface {
	LaunchWorker(w WorkerLauncher)
	MakeRequest(Request) error
	Stop()
}

type dispatcher struct {
	inCh chan Request
}

func NewDispatcher(b int) Dispatcher {
	return &dispatcher{
		inCh: make(chan Request, b),
	}
}

func (d *dispatcher) LaunchWorker(w WorkerLauncher) {
	w.LaunchWorker(d.inCh)
}

func (d *dispatcher) Stop() {
	close(d.inCh)
}

func (d *dispatcher) MakeRequest(r Request) error {
	select {
	case d.inCh <- r:
		return nil
	case <-time.After(time.Second * 5):
		return errors.New("timeout after 5 seconds")
	}
}
