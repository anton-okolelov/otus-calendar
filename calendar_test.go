package calendar

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	calendar := New()

	event := Event{
		Start:   time.Time{},
		End:     time.Time{},
		Payload: "aaa",
	}
	eventId := calendar.AddEvent(event)
	savedEvent, err := calendar.Event(eventId)

	assert.Equal(t, event, savedEvent)
	assert.Nil(t, err)
}

func TestDelete(t *testing.T) {
	calendar := New()

	event := Event{
		Start:   time.Time{},
		End:     time.Time{},
		Payload: "aaa",
	}
	eventId := calendar.AddEvent(event)
	calendar.DeleteEvent(eventId)

	_, err := calendar.Event(eventId)
	assert.NotNil(t, err)
}


func TestUpdate(t *testing.T) {
	calendar := New()

	event := Event{
		Start:   time.Time{},
		End:     time.Time{},
		Payload: "aaa",
	}
	eventId := calendar.AddEvent(event)
	
	calendar.UpdateEvent(eventId, Event{
		Start:   time.Time{},
		End:     time.Time{},
		Payload: "bbb",
	})
	
	savedEvent, err := calendar.Event(eventId)

	assert.Equal(t, "bbb", savedEvent.Payload)
	assert.Nil(t, err)
}