package Database

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Database struct {
	client *mongo.Client
	ctx    context.Context
}

type User struct {
	Username string `bson:"username" json:"username"`
	// "If deserialized to JSON it's empty for a reason"
	Password string           `bson:"password" json:"password,omitempty"`
	Notes    []string         `json:"notes"`
	Calendar []*CalendarEvent `json:"calendar"`
}

type CalendarEvent struct {
	Name string `json:"title" bson:"name"`
	Time uint64 `json:"time" bson:"time"`
}

type Session struct {
	SessionID string `bson:"sID" json:"sID"`
	Target    string `bson:"sTarget" json:"sTarget"`
}
