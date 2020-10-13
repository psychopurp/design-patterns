package observer

import (
	"sync"
	"testing"
	"time"
)

func TestObserver(t *testing.T) {
	sub := eventSubject{
		observers: sync.Map{},
	}

	obs1 := eventObserver{name: "obs1", time: time.Now()}
	obs2 := eventObserver{name: "obs2", time: time.Now()}
	obs3 := eventObserver{name: "obs3", time: time.Now()}

	sub.AddListener(&obs1)
	sub.AddListener(&obs2)
	sub.AddListener(&obs3)

	for i := 0; i < 10; i++ {
		sub.Notify(Event{data: i})
	}

}
