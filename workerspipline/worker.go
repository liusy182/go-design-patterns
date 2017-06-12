package worker

import (
	"fmt"
	"strings"
)

type WorkerLauncher interface {
	LaunchWorker(in chan Request)
}

type PreffixSuffixWorker struct {
	id      int
	prefixS string
	suffixS string
}

func (w *PreffixSuffixWorker) LaunchWorker(in chan Request) {
	r := <-w.prefix(w.append(w.uppercase(in)))
	r.Handler(r.Data)
}

func (w *PreffixSuffixWorker) uppercase(in <-chan Request) <-chan Request {
	out := make(chan Request)

	go func() {
		for msg := range in {
			s, ok := msg.Data.(string)

			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Data = strings.ToUpper(s)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (w *PreffixSuffixWorker) append(in <-chan Request) <-chan Request {
	out := make(chan Request)
	go func() {
		for msg := range in {
			uppercaseString, ok := msg.Data.(string)

			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Data = fmt.Sprintf("%s%s", uppercaseString, w.suffixS)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (w *PreffixSuffixWorker) prefix(in <-chan Request) <-chan Request {
	out := make(chan Request)
	go func() {
		for msg := range in {
			uppercasedStringWithSuffix, ok := msg.Data.(string)

			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Data = fmt.Sprintf("%s%s", w.prefixS, uppercasedStringWithSuffix)
			out <- msg
		}
		close(out)
	}()
	return out
}
