package Database

import (
	"GYMAppAPI/Config"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func NewDatabase(URL string) (*Database, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(URL))

	if err != nil {
		return nil, err
	}

	var res *Database = &Database{
		client: client,
		ctx:    context.Background(),
	}

	Connection = res

	return res, res.Ping()

}

// Make sure your connection is alive
func (db *Database) Ping() error {
	return db.client.Ping(db.ctx, readpref.Primary())
}

func (db *Database) Disconnect() error {
	return db.client.Disconnect(db.ctx)
}

func (db *Database) SessionFind(id string) (*Session, bool, error) {
	collection := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_SESSION_COLLECTION)
	filter := bson.D{{Key: "sID", Value: id}}
	var res *Session = &Session{}
	err := collection.FindOne(db.ctx, filter).Decode(res)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}
	return res, true, nil
}

func (db *Database) SessionCreate(target *User, id string) (*Session, error) {
	collection := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_SESSION_COLLECTION)
	var res *Session = &Session{
		SessionID: id,
		Target:    target.Username,
	}
	_, err := collection.InsertOne(db.ctx, res)
	return res, err
}

func (db *Database) UserFindUsername(username string) (*User, bool, error) {
	collection := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)
	filter := bson.D{{Key: "username", Value: username}}
	var res *User = &User{}
	err := collection.FindOne(db.ctx, filter).Decode(res)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}
	return res, true, nil
}

func (db *Database) UserFromSession(sessionID string) (*User, bool, error) {
	session, found, err := db.SessionFind(sessionID)
	if err != nil {
		return nil, false, err
	}

	if !found {
		return nil, false, nil
	}

	return db.UserFindUsername(session.Target)
}

func (db *Database) UpdateNote(username string, i int, newNotes string) error {
	collection := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)
	filter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{fmt.Sprintf("notes.%v", i): newNotes}}
	_, err := collection.UpdateOne(db.ctx, filter, update)
	return err
}

func (db *Database) DeleteNote(username string, index int) error {
	collection := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)
	filter := bson.M{"username": username}
	var result struct {
		Notes []string `bson:"notes"`
	}
	err := collection.FindOne(db.ctx, filter).Decode(&result)
	if err != nil {
		return err
	}

	if index < 0 || index >= len(result.Notes) {
		return fmt.Errorf("index %d out of range", index)
	}

	result.Notes = append(result.Notes[:index], result.Notes[index+1:]...)
	update := bson.M{"$set": bson.M{"notes": result.Notes}}
	_, err = collection.UpdateOne(db.ctx, filter, update)
	return err
}

func (db *Database) AddNoteToUser(username string, note string) error {
	collection := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)
	filter := bson.M{"username": username}
	update := bson.M{"$push": bson.M{"notes": note}}
	_, err := collection.UpdateOne(db.ctx, filter, update)
	return err
}

func (db *Database) AddCalendarEvent(username string, cal *CalendarEvent) error {
	collection := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)
	filter := bson.M{"username": username}
	update := bson.M{"$push": bson.M{"calendar": cal}}
	_, err := collection.UpdateOne(db.ctx, filter, update)
	return err
}
