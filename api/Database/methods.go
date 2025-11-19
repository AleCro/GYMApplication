package Db

import (
	"Svelgok-API/Environment"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/matthewhartstonge/argon2"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (u *User) Claims(session string) UserJWTClaim {
	return UserJWTClaim{
		UserID:    u.ID.Hex(),
		SessionID: session,
		Username:  u.Username,
		Group:     u.Group,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(Environment.JWT_TOKEN_LIFESPAN)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}

func (u *User) PasswordMatch(pwd string) bool {
	match, err := argon2.VerifyEncoded([]byte(pwd), []byte(u.Password))
	if err != nil {
		return false
	}
	return match
}

func (u *User) ChangePassword(new_pwd string) (*User, error) {
	encoded, err := Argon2.HashEncoded([]byte(new_pwd))
	if err != nil {
		return u, err
	}
	filter := bson.D{{"_id", u.ID}}
	update := bson.D{{"$set", bson.D{{"password", string(encoded)}}}}
	_, err = Connection.client.Database(Environment.DB_NAME).Collection(Environment.DB_USERS_COLLECTION).UpdateOne(Connection.ctx, filter, update)
	if err != nil {
		return u, err
	}

	usr, _, err := Connection.FilterOneUser(filter)
	if err != nil {
		return u, err
	}

	if usr != nil {
		u = usr
	}

	return u, nil
}

// `<User>.Update` pushes any changes to the User structure back to the database.
// It finds the user by its `ID` field and performs a `$set` operation
// with the current structure's values.
//
// This method uses a value receiver (`(u User)`) so that it operates on a copy.
// Setting `u.ID = nil` only affects this copy, preventing the immutable `_id`
// field from being included in the `$set` operation while leaving the original
// `User` structure unmodified.
//
// Returns:
//   - A `mongo.UpdateResult` containing information about the operation.
//   - An error if the user has no ID, the update fails, or no document is matched.
func (u User) Update() (*mongo.UpdateResult, error) {
	if u.ID == nil {
		return nil, errors.New("user has no ID")
	}

	collection := Connection.client.Database(Environment.DB_NAME).Collection(Environment.DB_USERS_COLLECTION)

	filter := bson.D{{"_id", u.ID}}
	u.ID = nil

	res, err := collection.UpdateOne(Connection.ctx, filter, bson.D{{"$set", u}})
	if err != nil {
		return nil, err
	}

	if res.MatchedCount == 0 {
		return res, mongo.ErrNoDocuments
	}

	return res, nil
}

func (u *User) ActiveSessions() ([]*Session, error) {
	if u.ID == nil {
		return nil, errors.New("user has no ID")
	}
	collection := Connection.client.Database(Environment.DB_NAME).Collection(Environment.DB_USERS_COLLECTION)
	cursor, err := collection.Find(Connection.ctx, bson.D{{"targetUser", u.ID}}, options.Find().SetSort(bson.D{{"createdAt", -1}}))
	if err != nil {
		return nil, fmt.Errorf("failed to find active sessions: %w", err)
	}
	defer cursor.Close(Connection.ctx)

	var sessions []*Session
	if err := cursor.All(Connection.ctx, &sessions); err != nil {
		return nil, fmt.Errorf("failed to decode active sessions: %v", err.Error())
	}
	return sessions, nil
}

func (u *User) InvalidateSessions() error {
	collection := Connection.client.Database(Environment.DB_NAME).Collection(Environment.DB_USERS_COLLECTION)
	_, err := collection.DeleteMany(Connection.ctx, bson.D{{"targetUser", u.ID}})
	if err != nil {
		return err
	}
	return nil
}
