package main

import (
	"fmt"
	"time"
)

type (
	Event struct {
		event int
	}
	Observer interface {
		whenNotifying(Event)
	}
	Notifier interface {
		Register(Observer)
		Unregister(Observer)
		Notifying(Event)
	}
)
type (
	observerEvent struct {
		id int
	}

	eventNotifier struct {
		observers map[Observer]struct{}
	}
)

func (a *observerEvent) whenNotifying(b Event) {
	fmt.Printf("We got %d observer: %d\n", a.id, b.event)
}

func (o *eventNotifier) Register(l Observer) {
	o.observers[l] = struct{}{}
}

func (o *eventNotifier) Unregister(l Observer) {
	delete(o.observers, l)
}

func (p *eventNotifier) Notifying(e Event) {
	for o := range p.observers {
		o.whenNotifying(e)
	}
}

func main() {

	n := eventNotifier{
		observers: map[Observer]struct{}{},
	}

	n.Register(&observerEvent{id: 1})
	exit := time.NewTimer(10 * time.Second).C
	next := time.NewTicker(time.Second).C
	for {
		select {
		case <-exit:
			return
		case t := <-next:
			n.Notifying(Event{event: int(t.UnixNano())})
		}
	}
}
