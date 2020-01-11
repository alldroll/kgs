package event

import (
	"sync/atomic"
	"testing"

	"github.com/asaskevich/EventBus"

	"github.com/short-d/app/mdtest"
)

func TestEventDispatcher_Dispatch(t *testing.T) {
	testCases := []struct {
		name      string
		events    []Event
		listeners []Listener
	}{
		{
			name: "one event one listener",
			events: []Event{
				fakeEvent("event1"),
			},
			listeners: []Listener{
				&fakeListener{
					name:          "event1",
					expectedCalls: 1,
				},
			},
		},
		{
			name: "one event two listener",
			events: []Event{
				fakeEvent("event1"),
			},
			listeners: []Listener{
				&fakeListener{
					name:          "event1",
					expectedCalls: 1,
				},
				&fakeListener{
					name:          "event1",
					expectedCalls: 1,
				},
			},
		},
		{
			name: "one of the listeners listens to another event",
			events: []Event{
				fakeEvent("event1"),
			},
			listeners: []Listener{
				&fakeListener{
					name:          "event1",
					expectedCalls: 1,
				},
				&fakeListener{
					name:          "event2",
					expectedCalls: 0,
				},
			},
		},
		{
			name: "two events two listeners",
			events: []Event{
				fakeEvent("event1"),
				fakeEvent("event2"),
			},
			listeners: []Listener{
				&fakeListener{
					name:          "event2",
					expectedCalls: 1,
				},
				&fakeListener{
					name:          "event1",
					expectedCalls: 1,
				},
			},
		},
		{
			name: "multiple events calls",
			events: []Event{
				fakeEvent("event1"),
				fakeEvent("event1"),
				fakeEvent("event2"),
				fakeEvent("event1"),
				fakeEvent("event1"),
				fakeEvent("event2"),
			},
			listeners: []Listener{
				&fakeListener{
					name:          "event2",
					expectedCalls: 2,
				},
				&fakeListener{
					name:          "event1",
					expectedCalls: 4,
				},
				&fakeListener{
					name:          "event1",
					expectedCalls: 4,
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			eventDispatcher := NewEventDispatcher(EventBus.New())

			err := BindListeners(eventDispatcher, testCase.listeners)
			mdtest.Equal(t, nil, err)

			for _, e := range testCase.events {
				err = eventDispatcher.Dispatch(e)
				mdtest.Equal(t, nil, err)
			}

			err = eventDispatcher.Close()
			mdtest.Equal(t, nil, err)

			for _, obj := range testCase.listeners {
				listener, ok := obj.(*fakeListener)

				mdtest.Equal(t, true, ok)
				mdtest.Equal(t, listener.expectedCalls, listener.actualCalls)
			}
		})
	}
}

func TestEventDispatcher_Close(t *testing.T) {
	ev := fakeEvent("event1")
	eventDispatcher := NewEventDispatcher(EventBus.New())
	listener := &fakeListener{name: ev.GetName()}

	_ = BindListeners(eventDispatcher, []Listener{listener})

	err := eventDispatcher.Dispatch(ev)
	mdtest.Equal(t, nil, err)

	err = eventDispatcher.Close()
	mdtest.Equal(t, nil, err)

	err = eventDispatcher.Dispatch(ev)
	mdtest.Equal(t, ErrDispatcherIsClosed, err)

	// there was no any listener call, because we have unsubscribed the listener
	mdtest.Equal(t, int32(1), listener.actualCalls)
	mdtest.Equal(t, ErrDispatcherIsClosed, eventDispatcher.Close())
}

type fakeListener struct {
	name          string
	expectedCalls int32
	actualCalls   int32
}

func (f *fakeListener) Handle(e Event) {
	// atomically increase actualCalls number, because of concurrent access to this variable
	atomic.AddInt32(&f.actualCalls, 1)
}

func (f *fakeListener) GetSubscribedEvent() string {
	return f.name
}

type fakeEvent string

func (f fakeEvent) GetName() string {
	return string(f)
}
