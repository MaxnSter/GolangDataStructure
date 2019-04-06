package EventBus

import (
	"container/list"
	"sync"
)

type event func()
type unSubscribe func()

type subscriber struct {
	topic string
	guard sync.Mutex
	fs    *list.List
}

func (s *subscriber) add(f event) (error, unSubscribe) {
	s.guard.Lock()
	l := s.fs.PushBack(f)
	s.guard.Unlock()

	one := sync.Once{}
	return nil, func() {
		one.Do(func() {
			s.guard.Lock()
			s.fs.Remove(l)
			s.guard.Unlock()
		})
	}
}

func (s *subscriber) fire() {
	var fs []event
	s.guard.Lock()
	for f := s.fs.Front(); f != nil; f = f.Next() {
		fs = append(fs, f.Value.(event))
	}
	s.guard.Unlock()

	for _, f := range fs {
		f()
	}
}

type eventBus struct {
	guard sync.Mutex
	m     map[string]*subscriber
}

func New() *eventBus {
	return &eventBus{
		m: map[string]*subscriber{},
	}
}

func (b *eventBus) Subscribe(topic string, f event) unSubscribe {
	b.guard.Lock()
	if _, ok := b.m[topic]; !ok {
		b.m[topic] = &subscriber{
			topic: topic,
			fs:    list.New(),
		}
	}
	b.guard.Unlock()

	_, cancel := b.m[topic].add(f)
	return cancel
}

func (b *eventBus) Fire(topic string) {
	b.guard.Lock()
	_, ok := b.m[topic]
	b.guard.Unlock()

	if !ok {
		return
	}
	b.m[topic].fire()
}
