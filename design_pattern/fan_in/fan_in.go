package fan_in

import (
	"github.com/MaxnSter/GolangDataStructure/try"
	"sort"
	"sync"
)

type Out = <-chan interface{}

func FanIn(cs ...<-chan interface{}) Out {
	out := make(chan interface{})
	wg := sync.WaitGroup{}
	transfer := func(ch <-chan interface{}) {
		for data := range ch {
			out <- data
		}
		wg.Done()
	}
	done := func() {
		wg.Wait()
		close(out)
	}

	for _, ch := range cs {
		wg.Add(1)
		go transfer(ch)
	}
	go done()

	return out
}

type PriorityCh struct {
	Ch       <-chan interface{}
	Priority int
}

func FanInWithPriority(ps ...PriorityCh) Out {
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].Priority > ps[j].Priority
	})

	var cs []<-chan interface{}
	for _, p := range ps {
		cs = append(cs, p.Ch)
	}
	return FanInWithPicker(func(idx []int) []int {
		return idx
	}, cs...)
}

type token struct {
	t chan interface{}
	done chan struct{}
}

type OnPick func(available []int) []int

func FanInWithPicker(picker OnPick, cs ...<-chan interface{}) Out {
	if picker == nil {
		return FanIn(cs...)
	}

	ts := make([]token, len(cs))
	idx := make([]int, len(cs))
	for i := range cs {
		idx[i] = i
		ts[i] = token{
			t:make(chan interface{}),
			done:make(chan struct{}),
		}
	}
	closeToken := func(i int) {
		close(ts[i].done)
	}
	isTokenClosed := func(i int) bool {
		select {
		case <-ts[i].done:
			return true
		default:
			return false
		}
	}
	put := func(i int, data interface{}) {
		ts[i].t <- data
	}
	take := func(i int) interface{}{
		select {
		case <-ts[i].done:
			return nil
		case data := <-ts[i].t:
			return data
		}
	}
	receive := func(i int, ch <-chan interface{}) {
		for data := range ch {
			put(i, data)
		}
		closeToken(i)
	}

	out := make(chan interface{})
	// start work
	for i, ch := range cs {
		go receive(i, ch)
	}
	// picking
	tryF := func() error{
		for {
			var available []int
			for _, i := range idx {
				if !isTokenClosed(i) {
					available = append(available, i)
				}
			}
			if len(available) <= 0 {
				break
			}

			idx = picker(available)
			for _, i := range idx {
				data := take(i)
				if data != nil {
					out <- data
				}
			}
		}
		return nil
	}
	final := func(_ error) error{
		close(out)
		return nil
	}
	go try.Try(tryF).Final(final).Do()
	return out
}
