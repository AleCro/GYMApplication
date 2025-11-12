//Mehtods of the DB.

package Database

import (
	"GYMAppAPI/Config"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// --- Connection ---
func NewDatabase(URL string) (*Database, error) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URL)) // <-- ctx first
	if err != nil {
		return nil, err
	}
	db := &Database{client: client, ctx: ctx}
	Connection = db
	return db, db.Ping()
}

func (db *Database) Ping() error           { return db.client.Ping(db.ctx, readpref.Primary()) }
func (db *Database) Client() *mongo.Client { return db.client }

// --- Sessions & Users ---

func (db *Database) SessionFind(id string) (*Session, bool, error) {
	coll := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_SESSION_COLLECTION)
	var out Session
	err := coll.FindOne(db.ctx, bson.D{{Key: "sID", Value: id}}).Decode(&out)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}
	return &out, true, nil
}

func (db *Database) SessionCreate(target *User, id string) (*Session, error) {
	coll := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_SESSION_COLLECTION)
	s := &Session{SessionID: id, Target: target.Username}
	_, err := coll.InsertOne(db.ctx, s)
	return s, err
}

func (db *Database) UserFindUsername(username string) (*User, bool, error) {
	coll := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)
	var out User
	err := coll.FindOne(db.ctx, bson.D{{Key: "username", Value: username}}).Decode(&out)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}
	return &out, true, nil
}

func (db *Database) UserFromSession(sessionID string) (*User, bool, error) {
	sess, ok, err := db.SessionFind(sessionID)
	if err != nil || !ok {
		return nil, false, err
	}
	return db.UserFindUsername(sess.Target)
}

// --- Notes ---

func (db *Database) UpdateNote(username string, i int, newNote string) error {
	coll := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)
	_, err := coll.UpdateOne(db.ctx,
		bson.M{"username": username},
		bson.M{"$set": bson.M{fmt.Sprintf("notes.%d", i): newNote}},
	)
	return err
}

func (db *Database) DeleteNote(username string, index int) error {
	coll := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)

	var doc struct {
		Notes []string `bson:"notes"`
	}
	if err := coll.FindOne(db.ctx, bson.M{"username": username}).Decode(&doc); err != nil {
		return err
	}
	if index < 0 || index >= len(doc.Notes) {
		return fmt.Errorf("index %d out of range", index)
	}
	doc.Notes = append(doc.Notes[:index], doc.Notes[index+1:]...)
	_, err := coll.UpdateOne(db.ctx, bson.M{"username": username}, bson.M{"$set": bson.M{"notes": doc.Notes}})
	return err
}

func (db *Database) AddNoteToUser(username, note string) error {
	coll := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)
	_, err := coll.UpdateOne(db.ctx, bson.M{"username": username}, bson.M{"$push": bson.M{"notes": note}})
	return err
}

// --- Calendar ---

func (db *Database) AddCalendarEvent(username string, cal *CalendarEvent) error {
	coll := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)
	_, err := coll.UpdateOne(db.ctx, bson.M{"username": username}, bson.M{"$push": bson.M{"calendar": cal}})
	return err
}

func (db *Database) DeleteCalendarEvent(username string, index int) error {
	coll := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)

	var user User
	if err := coll.FindOne(db.ctx, bson.M{"username": username}).Decode(&user); err != nil {
		return fmt.Errorf("user not found or decode error: %v", err)
	}
	if index < 0 || index >= len(user.Calendar) {
		return fmt.Errorf("index %d out of range (max %d)", index, len(user.Calendar)-1)
	}
	user.Calendar = append(user.Calendar[:index], user.Calendar[index+1:]...)
	_, err := coll.UpdateOne(db.ctx, bson.M{"username": username}, bson.M{"$set": bson.M{"calendar": user.Calendar}})
	return err
}

// --- Progress ---

func (db *Database) AddProgress(username string, entry ProgressEntry) error {
	coll := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)
	if entry.ID.IsZero() {
		entry.ID = primitive.NewObjectID()
	}
	_, err := coll.UpdateOne(db.ctx, bson.M{"username": username}, bson.M{"$push": bson.M{"progress": entry}})
	if err != nil {
		return fmt.Errorf("failed to add progress: %v", err)
	}
	return nil
}

func (db *Database) GetProgress(username string) ([]ProgressEntry, error) {
	user, ok, err := db.UserFindUsername(username)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	if user.Progress == nil {
		return []ProgressEntry{}, nil
	}

	patched := false
	for i := range user.Progress {
		if user.Progress[i].ID.IsZero() {
			user.Progress[i].ID = primitive.NewObjectID()
			patched = true
		}
	}
	if patched {
		_, _ = db.client.Database(Config.DATABASE_NAME).
			Collection(Config.DATABASE_USER_COLLECTION).
			UpdateOne(db.ctx,
				bson.M{"username": username},
				bson.M{"$set": bson.M{"progress": user.Progress}},
			)
	}
	return user.Progress, nil
}

func (db *Database) DeleteProgress(username string, progressID primitive.ObjectID) error {
	coll := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)

	var u struct {
		Progress []ProgressEntry `bson:"progress"`
	}
	if err := coll.FindOne(db.ctx, bson.M{"username": username}).Decode(&u); err != nil {
		return fmt.Errorf("user not found or decode error: %v", err)
	}

	out := make([]ProgressEntry, 0, len(u.Progress))
	removed := false
	for _, p := range u.Progress {
		if p.ID != progressID {
			out = append(out, p)
		} else {
			removed = true
		}
	}
	if !removed {
		return fmt.Errorf("progress entry %s not found", progressID.Hex())
	}

	_, err := coll.UpdateOne(db.ctx, bson.M{"username": username}, bson.M{"$set": bson.M{"progress": out}})
	return err
}

// --- Goals ---

func (db *Database) AddGoal(username string, goal Goal) error {
	coll := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)
	_, err := coll.UpdateOne(db.ctx, bson.M{"username": username}, bson.M{"$push": bson.M{"goals": goal}})
	return err
}

func (db *Database) GetGoals(username string) ([]Goal, error) {
	coll := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)
	var out struct {
		Goals []Goal `bson:"goals"`
	}
	if err := coll.FindOne(db.ctx, bson.M{"username": username}).Decode(&out); err != nil {
		return nil, err
	}
	return out.Goals, nil
}

func (db *Database) UpdateGoalSteps(username string, index int, steps []string) error {
	coll := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)

	// optional: validate index
	var doc struct {
		Goals []Goal `bson:"goals"`
	}
	if err := coll.FindOne(db.ctx, bson.M{"username": username}).Decode(&doc); err != nil {
		return fmt.Errorf("user not found or decode error: %v", err)
	}
	if index < 0 || index >= len(doc.Goals) {
		return fmt.Errorf("goal index %d out of range (max %d)", index, len(doc.Goals)-1)
	}

	path := fmt.Sprintf("goals.%d.steps", index)
	_, err := coll.UpdateOne(db.ctx, bson.M{"username": username}, bson.M{"$set": bson.M{path: steps}})
	return err
}

// ⬅️ Missing earlier — now added
func (db *Database) UpdateGoal(username string, index int, completed bool) error {
	coll := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)

	var doc struct {
		Goals []Goal `bson:"goals"`
	}
	if err := coll.FindOne(db.ctx, bson.M{"username": username}).Decode(&doc); err != nil {
		return fmt.Errorf("user not found or decode error: %v", err)
	}
	if index < 0 || index >= len(doc.Goals) {
		return fmt.Errorf("goal index %d out of range (max %d)", index, len(doc.Goals)-1)
	}

	path := fmt.Sprintf("goals.%d.completed", index)
	_, err := coll.UpdateOne(db.ctx, bson.M{"username": username}, bson.M{"$set": bson.M{path: completed}})
	return err
}

func (db *Database) DeleteGoal(username string, index int) error {
	coll := db.client.Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)

	var doc struct {
		Goals []Goal `bson:"goals"`
	}
	if err := coll.FindOne(db.ctx, bson.M{"username": username}).Decode(&doc); err != nil {
		return fmt.Errorf("user not found or decode error: %v", err)
	}
	if index < 0 || index >= len(doc.Goals) {
		return fmt.Errorf("goal index %d out of range (max %d)", index, len(doc.Goals)-1)
	}
	doc.Goals = append(doc.Goals[:index], doc.Goals[index+1:]...)
	_, err := coll.UpdateOne(db.ctx, bson.M{"username": username}, bson.M{"$set": bson.M{"goals": doc.Goals}})
	return err
}
