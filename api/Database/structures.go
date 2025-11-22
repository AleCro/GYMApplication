package Db

import (
	"Svelgok-API/Environment"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type GroupType uint8

const (
	GroupUser  GroupType = 0x00
	GroupAdmin GroupType = 0xFF
)

// Defines the schema of a user.
type User struct {
	ID       *bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email    string         `json:"email" bson:"email"`
	Username string         `json:"username" bson:"username"`
	Password string         `json:"password,omitempty" bson:"password,omitempty"`
	Group    GroupType      `json:"group" bson:"group"`
}

// A record that holds an ID used to identify the
// specific session, and field `TargetUser` which
// the session aims to.
//
// `CreatedAt`: For MongoDB index to remove it after
//
//	a pre-defined amount of time.
type Session struct {
	ID         *bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	TargetUser *bson.ObjectID `json:"targetUser,omitempty" bson:"targetUser,omitempty"`
	CreatedAt  time.Time      `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}

type UserJWTClaim struct {
	SessionID string    `json:"sessionID,omitempty"`
	UserID    string    `json:"userID,omitempty"`
	Username  string    `json:"username"`
	Group     GroupType `json:"group"`
	jwt.RegisteredClaims
}

func (c *UserJWTClaim) Expired() bool {
	if Environment.STRICT_SESSION_CONSISTENCY {
		if time.Now().After(c.IssuedAt.Time.Add(Environment.JWT_TOKEN_LIFESPAN + Environment.STRICT_SESSION_CONSISTENCY_OFFSET)) {
			return true
		}
	}

	return time.Now().After(c.ExpiresAt.Time)
}

type Note struct {
	ID        *bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Owner     *bson.ObjectID `json:"owner,omitempty" bson:"owner,omitempty"`
	Title     string         `json:"title" bson:"title"`
	Content   string         `json:"content" bson:"content"`
	CreatedAt time.Time      `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt" bson:"updatedAt"`
}

type Event struct {
	ID          *bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Owner       *bson.ObjectID `json:"owner,omitempty" bson:"owner,omitempty"`
	Title       string         `json:"title" bson:"title"`
	Description string         `json:"description" bson:"description"`
	Date        time.Time      `json:"date" bson:"date"` // We can just store the date part or full time
	CreatedAt   time.Time      `json:"createdAt" bson:"createdAt"`
}

type Progress struct {
	ID          *bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Owner       *bson.ObjectID `json:"owner,omitempty" bson:"owner,omitempty"`
	Title       string         `json:"title" bson:"title"`
	Description string         `json:"description" bson:"description"`
	ImageData   string         `json:"imageData" bson:"imageData"`
	CreatedAt   time.Time      `json:"createdAt" bson:"createdAt"`
}

type SubGoal struct {
	ID        string `json:"id" bson:"id"` // Client-side generated ID or UUID
	Title     string `json:"title" bson:"title"`
	Completed bool   `json:"completed" bson:"completed"`
}

type Goal struct {
	ID          *bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Owner       *bson.ObjectID `json:"owner,omitempty" bson:"owner,omitempty"`
	Title       string         `json:"title" bson:"title"`
	Description string         `json:"description" bson:"description"`
	SubGoals    []SubGoal      `json:"subGoals" bson:"subGoals"`
	CreatedAt   time.Time      `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt" bson:"updatedAt"`
}
