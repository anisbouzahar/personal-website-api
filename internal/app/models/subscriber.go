package models

import "time"

type Subscriber struct {
	ID           string    `json:"-" bson:"_id,omitempty"`
	Email        string    `json:"email" bson:"email"`
	SubscribedAt time.Time `json:"subscribed_at" bson:"subscribed_at"`
}
