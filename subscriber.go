package main

import (
	"time"
)

// Subscriber Represents user
type Subscriber struct {
	ID        int            // primary key
	Email     string         `storm:"unique,index"` // this field will be indexed with a unique constraint
	Name      string         // this field will not be indexed
	Payments  []*Transaction // Transaction made by user as historical data
	ValidUpTo time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (*Subscriber) createNewSubscriber(e *Subscriber) *Subscriber {
	
	return e
}
