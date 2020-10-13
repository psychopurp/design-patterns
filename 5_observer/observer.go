package observer

import (
	"fmt"
	"sync"
	"time"
)

type (
	Event struct {
		data int
	}

	//observer 都必须实现此接口
	Observer interface {
		NotifyCallback(Event)
	}
	//subject 实现此接口
	Subject interface {
		AddListener(Observer)
		RemoveListener(Observer)
		Notify(Event)
	}

	eventObserver struct {
		subject Subject
		name    string
		time    time.Time
	}

	eventSubject struct {
		observers sync.Map
	}
)

//实现observer 接口
func (e *eventObserver) NotifyCallback(event Event) {
	fmt.Printf("Observer: %v Received: %v after: %v\n", e.name, event.data, time.Since(e.time))
}

//实现subject接口
func (s *eventSubject) AddListener(obs Observer) {
	s.observers.Store(obs, struct{}{})
}
func (s *eventSubject) RemoveListener(obs Observer) {
	s.observers.Delete(obs)
}
func (s *eventSubject) Notify(event Event) {
	s.observers.Range(func(key, value interface{}) bool {
		if key == nil || value == nil {
			return false
		}
		key.(Observer).NotifyCallback(event)
		return true
	})
}
