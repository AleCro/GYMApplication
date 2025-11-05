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

// --- Inferred/Required Structs (Added for completeness and ID support) ---
type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
	Notes []string `bson:"notes"`
    Calendar []CalendarEvent `bson:"calendar"`
    Progress []ProgressEntry `bson:"progress"`
    Goals []Goal `bson:"goals"`
}

type Session struct {
	SessionID string `bson:"sID"`
	Target string `bson:"target"`
}

type CalendarEvent struct {
	Name string `bson:"name"`
	Time uint64 `bson:"time"`
	Timezone string `bson:"timezone,omitempty"`
}

type ProgressEntry struct {
    ID string `bson:"id"` // Added ID field for robust deletion
	Date string `bson:"date"`
	Weight float64 `bson:"weight"`
	Message string `bson:"message"`
	PhotoURL string `bson:"photo"`
}

type Goal struct {
	Title string `bson:"title"`
	Steps []string `bson:"steps"`
	Completed bool `bson:"completed"`
}

type Database struct {
	client *mongo.Client
	ctx context.Context
}

var Connection *Database // Global connection instance

func NewDatabase(URL string) (*Database, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(URL))

	if err != nil {
		return nil, err
	}

	var res *Database = &Database{
		client: client,
		ctx: 	context.Background(),
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
		Target: 	target.Username,
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

func (db *Database) DeleteGoal(username string, index int) error {
	collection := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)
	filter := bson.M{"username": username}

	// 1. Fetch the user document to get the current goals array
	var user struct {
		Goals []Goal `bson:"goals"`
	}
	err := collection.FindOne(db.ctx, filter).Decode(&user)
	if err != nil {
		return fmt.Errorf("user not found or decode error: %v", err)
	}

	// 2. Validate the index
	if index < 0 || index >= len(user.Goals) {
		return fmt.Errorf("goal index %d out of range (max %d)", index, len(user.Goals)-1)
	}

	// 3. Remove the goal at the specified index using slice manipulation
	user.Goals = append(user.Goals[:index], user.Goals[index+1:]...)

	// 4. Update the document by saving the modified goals array back to MongoDB
	update := bson.M{"$set": bson.M{"goals": user.Goals}}
	_, err = collection.UpdateOne(db.ctx, filter, update)
	return err
}

func (db *Database) DeleteCalendarEvent(username string, index int) error {
	collection := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)
	filter := bson.M{"username": username}

	var user User
	err := collection.FindOne(db.ctx, filter).Decode(&user)
	if err != nil {
		return fmt.Errorf("user not found or decode error: %v", err)
	}

	if len(user.Calendar) == 0 {
		return fmt.Errorf("no events to delete")
	}
	if index < 0 || index >= len(user.Calendar) {
		return fmt.Errorf("index %d out of range (max %d)", index, len(user.Calendar)-1)
	}

	// remove event at index
	user.Calendar = append(user.Calendar[:index], user.Calendar[index+1:]...)

	update := bson.M{"$set": bson.M{"calendar": user.Calendar}}
	_, err = collection.UpdateOne(db.ctx, filter, update)
	return err
}
// -----------------------
// PROGRESS
// -----------------------
func (db *Database) AddProgress(username string, entry ProgressEntry) error {
	collection := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)
	filter := bson.M{"username": username}
	update := bson.M{"$push": bson.M{"progress": entry}}
	_, err := collection.UpdateOne(db.ctx, filter, update)
	return err
}

func (db *Database) GetProgress(username string) ([]ProgressEntry, error) {
	collection := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)
	filter := bson.M{"username": username}

	var result struct {
		Progress []ProgressEntry `bson:"progress"`
	}
	err := collection.FindOne(db.ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result.Progress, nil
}

// **NEW FUNCTION: DeleteProgress by ID**
func (db *Database) DeleteProgress(username string, entryID string) error {
    collection := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)
    filter := bson.M{"username": username}

    // Use $pull to remove the element from the 'progress' array that matches the ID
    update := bson.M{"$pull": bson.M{"progress": bson.M{"id": entryID}}}

    result, err := collection.UpdateOne(db.ctx, filter, update)
    if err != nil {
        return fmt.Errorf("failed to update user document: %v", err)
    }

    if result.ModifiedCount == 0 {
        // This case covers: user not found, or entryID not found in progress array
        return errors.New("progress entry not found or already deleted")
    }
    
    return nil
}

// -----------------------
// GOALS
// -----------------------
func (db *Database) AddGoal(username string, goal Goal) error {
	collection := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)
	filter := bson.M{"username": username}
	update := bson.M{"$push": bson.M{"goals": goal}}
	_, err := collection.UpdateOne(db.ctx, filter, update)
	return err
}

func (db *Database) GetGoals(username string) ([]Goal, error) {
	collection := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)
	filter := bson.M{"username": username}

	var result struct {
		Goals []Goal `bson:"goals"`
	}
	err := collection.FindOne(db.ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result.Goals, nil
}

func (db *Database) UpdateGoal(username string, index int, done bool) error {
	collection := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)
	updatePath := fmt.Sprintf("goals.%d.completed", index)
	filter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{updatePath: done}}
	_, err := collection.UpdateOne(db.ctx, filter, update)
	return err
}