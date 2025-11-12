/*
Defines the schema of the DB.
*/
package Database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	client *mongo.Client
	ctx    context.Context
}

type User struct {
	Username string          `bson:"username" json:"username"`
	Password string          `bson:"password" json:"password,omitempty"`
	Notes    []string        `bson:"notes" json:"notes"`
	Calendar []CalendarEvent `bson:"calendar" json:"calendar"`
	Progress []ProgressEntry `bson:"progress" json:"progress"`
	Goals    []Goal          `bson:"goals" json:"goals"`
}

type Session struct {
	SessionID string `bson:"sID" json:"sID"`
	Target    string `bson:"target" json:"target"`
}

type CalendarEvent struct {
	Name     string `bson:"name" json:"title"`
	Time     uint64 `bson:"time" json:"time"`
	Timezone string `bson:"timezone,omitempty" json:"timezone,omitempty"`
}

type Goal struct {
	Title     string   `bson:"title" json:"title"`
	Steps     []string `bson:"steps" json:"steps"`
	Completed bool     `bson:"completed" json:"completed"`
}

type ProgressEntry struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Date    string             `bson:"date" json:"date"`
	Weight  float64            `bson:"weight" json:"weight"`
	Message string             `bson:"message" json:"message"`
	Photo   string             `bson:"photo,omitempty" json:"photo,omitempty"`
}
