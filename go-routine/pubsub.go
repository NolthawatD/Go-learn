package main

import (
	"sync"
)

// Message represents a simple message structure
type Message struct {
	Content string
}

// Subscriber represents a subscriber to a topic
type Subscriber struct {
	ID      uint `gorm:"primaryKey"`
	Channel chan Message
}

// PubSub represents the publish/subscribe mechanism
type PubSub struct {
	mu           sync.RWMutex
	Subscribers  map[uint]*Subscriber
	subscriberID uint
}

func (ps *PubSub) Subscribe() *Subscriber {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ps.subscriberID++
	subscriber := &Subscriber{
		ID:      ps.subscriberID,
		Channel: make(chan Message),
	}
	ps.Subscribers[subscriber.ID] = subscriber
	return subscriber
}

func (ps *PubSub) Unsubscribe(id uint) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	close(ps.Subscribers[id].Channel)
	delete(ps.Subscribers, id)
}

func (ps *PubSub) Publish(message Message) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	for _, subscriber := range ps.Subscribers {
		subscriber.Channel <- message
	}
}
