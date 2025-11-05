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
	Notes    []string         `bson:"notes" json:"notes"`
	Calendar []*CalendarEvent `bson:"calendar" json:"calendar"`
	Progress []ProgressEntry  `bson:"progress" json:"progress"`
	Goals    []Goal           `bson:"goals" json:"goals"`
}

type CalendarEvent struct {
	Name string `json:"title" bson:"name"`
	Time uint64 `json:"time" bson:"time"`
	Timezone string `json:"timezone,omitempty"`
}

type Session struct {
	SessionID string `bson:"sID" json:"sID"`
	Target    string `bson:"sTarget" json:"sTarget"`
}

type Goal struct {
	Title     string   `bson:"title" json:"title"`
	Steps     []string `bson:"steps" json:"steps"`
	Completed bool     `bson:"completed" json:"completed"`
}

type ProgressEntry struct {
	Date     string  `bson:"date" json:"date"`
	Weight   float64 `bson:"weight" json:"weight"`
	Message  string  `bson:"message" json:"message"`
	PhotoURL string  `bson:"photo" json:"photo"`
}