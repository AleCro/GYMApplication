package Db

import (
	"Svelgok-API/Environment"
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

//      ___       _ _   _       _ _          _   _
//     |_ _|_ __ (_) |_(_) __ _| (_)______ _| |_(_) ___  _ __
//      | || '_ \| | __| |/ _` | | |_  / _` | __| |/ _ \| '_ \
//      | || | | | | |_| | (_| | | |/ / (_| | |_| | (_) | | | |
//     |___|_| |_|_|\__|_|\__,_|_|_/___\__,_|\__|_|\___/|_| |_|
//

// `Database` provides a wrapper around the MongoDB client and context.
// It acts as the primary data access layer (DAL), enforcing schema logic
// and abstracting database operations away from the HTTP handlers.
type Database struct {
	client *mongo.Client
	ctx    context.Context
}

// `Connect` establishes a new connection to the MongoDB server using the provided URL.
// On a successful connection, it pings the database to verify the connection.
// It also initializes the package-level `Connection` variable, implementing a
// method to easily access the database throughout the application.
//
// Returns a pointer to the initialized Database struct or an error if the
// connection or ping fails.
func Connect(URL string) (*Database, error) {
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

// Ping checks the liveness of the database connection by sending a ping
// to the primary node.
//
// Returns an error if the database cannot be reached.
func (db *Database) Ping() error {
	return db.client.Ping(db.ctx, readpref.Primary())
}

// `Initialize` ensures the database is ready for use by the application.
// It performs the following setup tasks:
//  1. Checks if the required 'users' and 'sessions' collections exist.
//  2. Creates the collections if they are missing.
//  3. Calls `ensureSessionsTTLIndex` to create or validate the TTL index
//     on the 'sessions' collection for session to expire automatically.
//
// Returns an error if any of these steps fail.
func (db *Database) Initialize() error {
	if db == nil || db.client == nil || db.ctx == nil {
		return fmt.Errorf("db.client or db.ctx is nil")
	}

	database := db.client.Database(Environment.DB_NAME)

	existing, err := database.ListCollectionNames(db.ctx, bson.D{})
	if err != nil {
		return err
	}
	hasUsers, hasSessions := false, false
	for _, n := range existing {
		if n == Environment.DB_USERS_COLLECTION {
			hasUsers = true
		}
		if n == Environment.DB_SESSIONS_COLLECTION {
			hasSessions = true
		}
	}
	if !hasUsers {
		if err := database.CreateCollection(db.ctx, Environment.DB_USERS_COLLECTION); err != nil {
			return err
		}
	}
	
	if !hasSessions {
		if err := database.CreateCollection(db.ctx, Environment.DB_SESSIONS_COLLECTION); err != nil {
			return err
		}
	}

	// Initialize new collections
	collections := []string{"notes", "events", "progress", "goals"}
	for _, name := range collections {
		found := false
		for _, n := range existing {
			if n == name {
				found = true
				break
			}
		}
		if !found {
			if err := database.CreateCollection(db.ctx, name); err != nil {
				return err
			}
		}
	}

	sessionsColl := database.Collection(Environment.DB_SESSIONS_COLLECTION)
	return ensureSessionsTTLIndex(db.ctx, sessionsColl)
}

// `ensureSessionsTTLIndex` is a helper function that guarantees a
// Time-To-Live (TTL) index exists on the 'sessions' collection.
//
// It targets the 'expiresAt' field, using the `SESSION_DURATION` from the
// environment to set the `expireAfterSeconds` value.
//
// Steps:
//  1. Lists all existing indexes on the collection.
//  2. Iterates them to find an index on 'expiresAt'.
//  3. If a *matching* TTL index is found (correct field, TTL value, and name), it does nothing.
//  4. If an *incorrect* index is found (e.g., wrong TTL, or a non-TTL index on the same field),
//     it is added to a list to be dropped.
//  5. After checking all indexes, it drops any conflicting ones.
//  6. If no matching index was found, it creates the correct TTL index.
//
// This ensures that MongoDB automatically deletes session documents when their
// `expiresAt` time is reached.
func ensureSessionsTTLIndex(ctx context.Context, collection *mongo.Collection) error {
	cur, err := collection.Indexes().List(ctx)
	if err != nil {
		return err
	}
	defer cur.Close(ctx)

	expSec := int32(Environment.SESSION_DURATION / time.Second)
	if expSec < 0 {
		return fmt.Errorf("SESSION_DURATION must be non-negative")
	}

	foundMatching := false
	var toDrop []string

	for cur.Next(ctx) {
		var spec bson.M
		if err := cur.Decode(&spec); err != nil {
			return err
		}

		nameAny, _ := spec["name"]
		name, _ := nameAny.(string)

		keyAny, _ := spec["key"]
		var usesExpiresAt bool
		switch k := keyAny.(type) {
		case bson.D:
			for _, v := range k {
				if v.Key == "expiresAt" {
					usesExpiresAt = true
					break
				}
			}
		case bson.M:
			if _, ok := k["expiresAt"]; ok {
				usesExpiresAt = true
			}
		case primitive.D:
			for _, kv := range k {
				if kv.Key == "expiresAt" {
					usesExpiresAt = true
					break
				}
			}
		}

		if !usesExpiresAt {
			continue
		}

		expAny, hasExpire := spec["expireAfterSeconds"]
		if hasExpire {
			var existing int64
			switch v := expAny.(type) {
			case int32:
				existing = int64(v)
			case int64:
				existing = v
			case float64:
				existing = int64(v)
			default:
				existing = -1
			}
			if existing == int64(expSec) && name == Environment.DB_SESSIONS_TLL_INDEX_NAME {
				foundMatching = true
				continue
			}
			if name != "" {
				toDrop = append(toDrop, name)
			}
			continue
		} else {
			if name != "" {
				toDrop = append(toDrop, name)
			}
			continue
		}
	}

	if err := cur.Err(); err != nil {
		return err
	}

	for _, nm := range toDrop {
		if nm == "" {
			continue
		}
		if err := collection.Indexes().DropOne(ctx, nm); err != nil {
			return err
		}
	}

	if !foundMatching {
		secs := expSec
		model := mongo.IndexModel{
			Keys:    bson.D{{Key: "expiresAt", Value: 1}},
			Options: options.Index().SetExpireAfterSeconds(secs).SetName(Environment.DB_SESSIONS_TLL_INDEX_NAME),
		}
		if _, err := collection.Indexes().CreateOne(ctx, model); err != nil {
			return err
		}
	}

	return nil
}

//      _   _
//     | | | |___  ___ _ __
//     | | | / __|/ _ \ '__|
//     | |_| \__ \  __/ |
//      \___/|___/\___|_|
//

// `GetUsers` retrieves a paginated and filtered list of users from the database.
//
// Parameters:
//   - `page`: The current page number (1-based).
//   - `limit`: The number of users to return per page.
//   - `query`: A string used to search against usernames (case-insensitive) and user IDs.
//
// Returns:
//   - A `PaginatedUsersResult` struct containing the list of users and pagination details.
//   - An error if the database query fails.
//
// Note: This function explicitly projects the 'password' field, removing it from all results.
func (db *Database) GetUsers(page, limit int64, query string) (*PaginatedUsersResult, error) {
	offset := (page - 1) * limit
	collection := db.client.Database(Environment.DB_NAME).Collection(Environment.DB_USERS_COLLECTION)

	filter := bson.D{}

	if query != "" {
		var orFilters bson.A
		orFilters = append(orFilters, bson.D{
			{"username", bson.D{{"$regex", query}, {"$options", "i"}}},
		})

		if id, err := strconv.ParseInt(query, 10, 64); err == nil {
			orFilters = append(orFilters, bson.D{{"_id", id}})
		}

		filter = bson.D{{"$or", orFilters}}
	}

	total, err := collection.CountDocuments(db.ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to count users: %v", err.Error())
	}

	cursor, err := collection.Find(db.ctx, filter, options.Find().
		SetLimit(limit).
		SetSkip(offset).
		SetSort(bson.D{{"createdAt", -1}}).
		SetProjection(bson.D{{"password", 0}}))
	if err != nil {
		return nil, fmt.Errorf("failed to find users: %v", err.Error())
	}
	defer cursor.Close(db.ctx)

	var users []*User
	if err := cursor.All(db.ctx, &users); err != nil {
		return nil, fmt.Errorf("failed to decode users: %v", err.Error())
	}

	return &PaginatedUsersResult{
		Users: users,
		Total: total,
		Limit: limit,
		Page:  page,
	}, nil
}

// `FilterOneUser` finds a single user document in the `users` collection
// that matches the provided BSON filter.
//
// Returns:
//   - A pointer to the decoded `User` struct.
//   - A boolean (true) if a user was found.
//   - An error if the query fails (but not if no document is found).
//
// If no user is found, it returns (nil, false, nil).
func (db *Database) FilterOneUser(filter bson.D) (*User, bool, error) {
	collection := db.client.Database(Environment.DB_NAME).Collection(Environment.DB_USERS_COLLECTION)
	var res *User = &User{}
	err := collection.FindOne(db.ctx, filter).Decode(res)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}
	return res, true, nil
}

// `FilterOneUserByID` is a helper function to find a user by their
// MongoDB ObjectID provided as a hex string.
//
// Returns:
//   - A pointer to the decoded `User` struct.
//   - A boolean (true) if a user was found.
//   - An error if the hex string is invalid or the query fails.
func (db *Database) FilterOneUserByID(hex string) (*User, bool, error) {
	objID, err := bson.ObjectIDFromHex(hex)
	if err != nil {
		return nil, false, err
	}

	return db.FilterOneUser(bson.D{{"_id", objID}})
}

// `InsertOneUser` adds a new user document to the `users` collection.
// It takes a pointer to a `User` struct, inserts it, and then populates
// the struct's `ID` field with the new `InsertedID` from the database.
//
// Returns:
//   - The same `User` pointer, now with the `ID` field populated.
//   - An error if the insertion fails.
func (db *Database) InsertOneUser(user *User) (*User, error) {
	collection := db.client.Database(Environment.DB_NAME).Collection(Environment.DB_USERS_COLLECTION)
	res, err := collection.InsertOne(db.ctx, user)
	if err != nil {
		return nil, err
	}
	id := res.InsertedID.(bson.ObjectID)
	user.ID = &id
	return user, nil
}

//      ____                _
//     / ___|  ___  ___ ___(_) ___  _ __
//     \___ \ / _ \/ __/ __| |/ _ \| '_ \
//      ___) |  __/\__ \__ \ | (_) | | | |
//     |____/ \___||___/___/_|\___/|_| |_|
//

// `CreateSession` inserts a new session document for a given user.
// The session's `TargetUser` field is set to the user's ID.
//
// Returns:
//   - A pointer to the newly created `Session` struct, with its `ID` field populated.
//   - An error if the insertion fails.
func (db *Database) CreateSession(user *User) (*Session, error) {
	var session *Session = &Session{
		TargetUser: user.ID,
		CreatedAt:  time.Now(),
	}
	collection := db.client.Database(Environment.DB_NAME).Collection(Environment.DB_SESSIONS_COLLECTION)
	res, err := collection.InsertOne(db.ctx, session)
	if err != nil {
		return nil, err
	}
	id := res.InsertedID.(bson.ObjectID)
	session.ID = &id
	return session, nil
}

// `UserFromSession` retrieves the `User` associated with a given `Session`
// by looking up the `session.TargetUser` ID.
//
// Returns:
//   - A pointer to the decoded `User` struct.
//   - A boolean (true) if the user was found.
//   - An error if the query fails.
func (db *Database) UserFromSession(session *Session) (*User, bool, error) {
	return db.FilterOneUser(bson.D{{"_id", session.TargetUser}})
}

// `GetSession` finds a single session document in the `sessions` collection
// that matches the provided BSON filter.
//
// Returns:
//   - A pointer to the decoded `Session` struct.
//   - A boolean (true) if a session was found.
//   - An error if the query fails (but not if no document is found).
//
// If no session is found, it returns (nil, false, nil).
func (db *Database) GetSession(filter bson.D) (*Session, bool, error) {
	collection := db.client.Database(Environment.DB_NAME).Collection(Environment.DB_SESSIONS_COLLECTION)
	var res *Session = &Session{}
	err := collection.FindOne(db.ctx, filter).Decode(res)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}
	return res, true, nil
}

// `GetSessionFromID` is a helper function to find a session by its
// MongoDB ObjectID provided as a hex string.
//
// Returns:
//   - A pointer to the decoded `Session` struct.
//   - A boolean (true) if a session was found.
//   - An error if the hex string is invalid or the query fails.
func (db *Database) GetSessionFromID(hex string) (*Session, bool, error) {
	objID, err := bson.ObjectIDFromHex(hex)
	if err != nil {
		return nil, false, err
	}

	return db.GetSession(bson.D{{"_id", objID}})
}

// `RemoveSession` deletes a single session from the database based on its `bson.ObjectID`.
//
// Returns:
//   - The count of deleted documents (0 or 1).
//   - An error if the deletion operation fails.
func (db *Database) RemoveSession(sessionID bson.ObjectID) (int64, error) {
	collection := db.client.Database(Environment.DB_NAME).Collection(Environment.DB_SESSIONS_COLLECTION)

	filter := bson.D{{"_id", sessionID}}

	res, err := collection.DeleteOne(db.ctx, filter)
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}

// `RemoveSessionFromID` deletes a single session from the database based on its
// hex string ID. This is a convenience wrapper around `RemoveSession`.
//
// Returns:
//   - The count of deleted documents (0 or 1).
//   - An error if the hex string is invalid or the deletion fails.
func (db *Database) RemoveSessionFromID(hex string) (int64, error) {
	objID, err := bson.ObjectIDFromHex(hex)
	if err != nil {
		return 0, fmt.Errorf("invalid session ID format: %w", err)
	}

	return db.RemoveSession(objID)
}

//          ___        _______  __     __    _ _     _       _   _
//         | \ \      / /_   _| \ \   / /_ _| (_) __| | __ _| |_(_) ___  _ __
//      _  | |\ \ /\ / /  | |    \ \ / / _` | | |/ _` |/ _` | __| |/ _ \| '_ \
//     | |_| | \ V  V /   | |     \ V / (_| | | | (_| | (_| | |_| | (_) | | | |
//      \___/   \_/\_/    |_|      \_/ \__,_|_|_|\__,_|\__,_|\__|_|\___/|_| |_|
//

// `VerifyJWTSignature` parses a JWT string, validates its signing method (HMAC)
// and signature against the `JWT_SECRET`.
//
// Parameters:
//   - `tokenStr`: The raw JWT token string.
//   - `ignoreJWTExpiration`: If true, the function will return the claims
//     even if the token is expired (`jwt.ErrTokenExpired`). This is useful to
//     refresh the token or when a session needs to be removed from the database
//     on logout.
//
// Returns:
//   - A pointer to the populated `UserJWTClaim` struct.
//   - An error if parsing, signature validation, or (if not ignored) expiration fails.
func VerifyJWTSignature(tokenStr string, ignoreJWTExpiration bool) (*UserJWTClaim, error) {
	var claims *UserJWTClaim = &UserJWTClaim{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenUnverifiable
		}
		return []byte(Environment.JWT_SECRET), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) && ignoreJWTExpiration {
			return claims, nil
		}

		return nil, err
	}

	if token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrTokenInvalidClaims
}

// `VerifyJWTSignatureEx` parses claims from a JWT token, *ignoring*
// JWT expiration errors.
//
// Useful for endpoints like refresh token, where you need to identify
// the user from an expired access token to issue a new one.
//
// It *only* validates the signature, not expiration. It does *not*
// check the database, can be paired up with `ValidateJWTSessionFromClaims`
// for a full database check.
//
// Returns:
//   - A pointer to the `UserJWTClaim` struct (even if expired).
//   - An error if the signature is invalid or the token is malformed.
func VerifyJWTSignatureEx(tokenStr string) (*UserJWTClaim, error) {
	claims := &UserJWTClaim{}
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenUnverifiable
		}
		return []byte(Environment.JWT_SECRET), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return claims, nil
		}

		return nil, err
	}

	return claims, nil
}

// `ValidateJWTSession` performs a full validation of a JWT for an active session.
// It chains signature verification with database checks.
//
// Logic:
// 1. Calls `VerifyJWTSignature` to validate the token's signature and claims.
// 2. Calls `ValidateJWTSessionFromClaims` to check the database.
//
// This is the primary function to use in middleware to protect authenticated routes.
//
// Returns:
//   - A pointer to the `User` associated with the token.
//   - A pointer to the `Session` associated with the token.
//   - An error if the JWT is invalid, the session doesn't exist, or the user doesn't exist.
func ValidateJWTSession(tokenStr string, ignoreJWTExpiration bool) (*User, *Session, error) {
	claims, err := VerifyJWTSignature(tokenStr, ignoreJWTExpiration)
	if err != nil {
		return nil, nil, err
	}
	return ValidateJWTSessionFromClaims(claims, ignoreJWTExpiration)
}

// `ValidateJWTSessionFromClaims` performs the database-level validation of JWT claims.
// It ensures that the session and user referenced in the claims still exist
// and are correctly linked.
//
// Logic:
// 1. Fetches the session from the DB using `claims.SessionID`.
// 2. Fetches the user from the DB using the `session.TargetUser` ID.
// 3. Verifies that the fetched `user.ID` matches the `claims.UserID`.
//
// Returns:
//   - A pointer to the `User`.
//   - A pointer to the `Session`.
//   - An error if the session is not found ("session expired"), the user is not found,
//     or there is a user/session mismatch.
func ValidateJWTSessionFromClaims(claims *UserJWTClaim, ignoreJWTExpiration bool) (*User, *Session, error) {
	session, found, err := Connection.GetSessionFromID(claims.SessionID)
	if err != nil {
		return nil, nil, err
	}

	if !found {
		return nil, nil, fmt.Errorf("session expired")
	}

	user, found, err := Connection.UserFromSession(session)
	if err != nil {
		return nil, nil, err
	}

	if !found {
		return nil, nil, fmt.Errorf("user not found")
	}

	if user.ID.Hex() != claims.UserID {
		return nil, nil, fmt.Errorf("target user does not match")
	}

	return user, session, nil
}

// `CreateJWTToken` generates a new JWT token string, signed with the HS256
// method and the application's `JWT_SECRET`.
//
// Parameters:
//   - claims: Any struct that satisfies the `jwt.Claims` interface
//     (e.g., `UserJWTClaim`).
//
// Returns:
//   - The signed token string.
//   - An error if signing fails.
func CreateJWTToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(Environment.JWT_SECRET))
}

// Notes CRUD

// CreateNote inserts a new note into the database.
// It sets the CreatedAt and UpdatedAt timestamps to the current time.
// Returns the created note with its new ID, or an error.
func (db *Database) CreateNote(note *Note) (*Note, error) {
	collection := db.client.Database(Environment.DB_NAME).Collection("notes")
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()
	res, err := collection.InsertOne(db.ctx, note)
	if err != nil {
		return nil, err
	}
	id := res.InsertedID.(bson.ObjectID)
	note.ID = &id
	return note, nil
}

// GetNotes retrieves all notes for a specific owner.
// Results are sorted by UpdatedAt in descending order (newest first).
func (db *Database) GetNotes(ownerID *bson.ObjectID) ([]*Note, error) {
	collection := db.client.Database(Environment.DB_NAME).Collection("notes")
	filter := bson.D{{"owner", ownerID}}
	cursor, err := collection.Find(db.ctx, filter, options.Find().SetSort(bson.D{{"updatedAt", -1}}))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(db.ctx)
	var notes []*Note
	if err := cursor.All(db.ctx, &notes); err != nil {
		return nil, err
	}
	return notes, nil
}

// UpdateNote modifies the title and content of an existing note.
// It also updates the UpdatedAt timestamp.
// Returns an error if the note is not found or the user is not the owner.
func (db *Database) UpdateNote(id string, ownerID *bson.ObjectID, title, content string) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	collection := db.client.Database(Environment.DB_NAME).Collection("notes")
	filter := bson.D{{"_id", objID}, {"owner", ownerID}}
	update := bson.D{{"$set", bson.D{
		{"title", title},
		{"content", content},
		{"updatedAt", time.Now()},
	}}}
	res, err := collection.UpdateOne(db.ctx, filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return fmt.Errorf("note not found or unauthorized")
	}
	return nil
}

// DeleteNote removes a note from the database.
// Returns an error if the note is not found or the user is not the owner.
func (db *Database) DeleteNote(id string, ownerID *bson.ObjectID) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	collection := db.client.Database(Environment.DB_NAME).Collection("notes")
	filter := bson.D{{"_id", objID}, {"owner", ownerID}}
	res, err := collection.DeleteOne(db.ctx, filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("note not found or unauthorized")
	}
	return nil
}

// Events CRUD

// CreateEvent inserts a new calendar event into the database.
// It sets the CreatedAt timestamp to the current time.
func (db *Database) CreateEvent(event *Event) (*Event, error) {
	collection := db.client.Database(Environment.DB_NAME).Collection("events")
	event.CreatedAt = time.Now()
	res, err := collection.InsertOne(db.ctx, event)
	if err != nil {
		return nil, err
	}
	id := res.InsertedID.(bson.ObjectID)
	event.ID = &id
	return event, nil
}

// GetEvents retrieves all events for a specific owner.
// Results are sorted by Date in ascending order (oldest first).
func (db *Database) GetEvents(ownerID *bson.ObjectID) ([]*Event, error) {
	collection := db.client.Database(Environment.DB_NAME).Collection("events")
	filter := bson.D{{"owner", ownerID}}
	cursor, err := collection.Find(db.ctx, filter, options.Find().SetSort(bson.D{{"date", 1}}))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(db.ctx)
	var events []*Event
	if err := cursor.All(db.ctx, &events); err != nil {
		return nil, err
	}
	return events, nil
}

// UpdateEvent modifies the details of an existing event.
// Returns an error if the event is not found or the user is not the owner.
func (db *Database) UpdateEvent(id string, ownerID *bson.ObjectID, title, description string, date time.Time) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	collection := db.client.Database(Environment.DB_NAME).Collection("events")
	filter := bson.D{{"_id", objID}, {"owner", ownerID}}
	update := bson.D{{"$set", bson.D{
		{"title", title},
		{"description", description},
		{"date", date},
	}}}
	res, err := collection.UpdateOne(db.ctx, filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return fmt.Errorf("event not found or unauthorized")
	}
	return nil
}

// DeleteEvent removes an event from the database.
// Returns an error if the event is not found or the user is not the owner.
func (db *Database) DeleteEvent(id string, ownerID *bson.ObjectID) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	collection := db.client.Database(Environment.DB_NAME).Collection("events")
	filter := bson.D{{"_id", objID}, {"owner", ownerID}}
	res, err := collection.DeleteOne(db.ctx, filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("event not found or unauthorized")
	}
	return nil
}

// Progress CRUD

// CreateProgress inserts a new progress entry into the database.
// It sets the CreatedAt timestamp to the current time.
func (db *Database) CreateProgress(progress *Progress) (*Progress, error) {
	collection := db.client.Database(Environment.DB_NAME).Collection("progress")
	progress.CreatedAt = time.Now()
	res, err := collection.InsertOne(db.ctx, progress)
	if err != nil {
		return nil, err
	}
	id := res.InsertedID.(bson.ObjectID)
	progress.ID = &id
	return progress, nil
}

// GetProgress retrieves all progress entries for a specific owner.
// Results are sorted by CreatedAt in descending order (newest first).
func (db *Database) GetProgress(ownerID *bson.ObjectID) ([]*Progress, error) {
	collection := db.client.Database(Environment.DB_NAME).Collection("progress")
	filter := bson.D{{"owner", ownerID}}
	cursor, err := collection.Find(db.ctx, filter, options.Find().SetSort(bson.D{{"createdAt", -1}}))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(db.ctx)
	var progress []*Progress
	if err := cursor.All(db.ctx, &progress); err != nil {
		return nil, err
	}
	return progress, nil
}

// UpdateProgress modifies the title and description of a progress entry.
// Note: Image data is currently not updatable via this method.
func (db *Database) UpdateProgress(id string, ownerID *bson.ObjectID, title, description string) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	collection := db.client.Database(Environment.DB_NAME).Collection("progress")
	filter := bson.D{{"_id", objID}, {"owner", ownerID}}
	update := bson.D{{"$set", bson.D{
		{"title", title},
		{"description", description},
	}}}
	res, err := collection.UpdateOne(db.ctx, filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return fmt.Errorf("progress not found or unauthorized")
	}
	return nil
}

// DeleteProgress removes a progress entry from the database.
// Returns an error if the entry is not found or the user is not the owner.
func (db *Database) DeleteProgress(id string, ownerID *bson.ObjectID) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	collection := db.client.Database(Environment.DB_NAME).Collection("progress")
	filter := bson.D{{"_id", objID}, {"owner", ownerID}}
	res, err := collection.DeleteOne(db.ctx, filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("progress not found or unauthorized")
	}
	return nil
}
// Goals CRUD

// CreateGoal inserts a new goal into the database.
// It sets the CreatedAt and UpdatedAt timestamps to the current time.
func (db *Database) CreateGoal(goal *Goal) (*Goal, error) {
	collection := db.client.Database(Environment.DB_NAME).Collection("goals")
	goal.CreatedAt = time.Now()
	goal.UpdatedAt = time.Now()
	res, err := collection.InsertOne(db.ctx, goal)
	if err != nil {
		return nil, err
	}
	id := res.InsertedID.(bson.ObjectID)
	goal.ID = &id
	return goal, nil
}

// GetGoals retrieves all goals for a specific owner.
// Results are sorted by UpdatedAt in descending order (newest first).
func (db *Database) GetGoals(ownerID *bson.ObjectID) ([]*Goal, error) {
	collection := db.client.Database(Environment.DB_NAME).Collection("goals")
	filter := bson.D{{"owner", ownerID}}
	cursor, err := collection.Find(db.ctx, filter, options.Find().SetSort(bson.D{{"updatedAt", -1}}))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(db.ctx)
	var goals []*Goal
	if err := cursor.All(db.ctx, &goals); err != nil {
		return nil, err
	}
	return goals, nil
}

// UpdateGoal modifies the title, description, and subgoals of an existing goal.
// It also updates the UpdatedAt timestamp.
func (db *Database) UpdateGoal(id string, ownerID *bson.ObjectID, title, description string, subGoals []SubGoal) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	collection := db.client.Database(Environment.DB_NAME).Collection("goals")
	filter := bson.D{{"_id", objID}, {"owner", ownerID}}
	update := bson.D{{"$set", bson.D{
		{"title", title},
		{"description", description},
		{"subGoals", subGoals},
		{"updatedAt", time.Now()},
	}}}
	res, err := collection.UpdateOne(db.ctx, filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return fmt.Errorf("goal not found or unauthorized")
	}
	return nil
}

// DeleteGoal removes a goal from the database.
func (db *Database) DeleteGoal(id string, ownerID *bson.ObjectID) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	collection := db.client.Database(Environment.DB_NAME).Collection("goals")
	filter := bson.D{{"_id", objID}, {"owner", ownerID}}
	res, err := collection.DeleteOne(db.ctx, filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("goal not found or unauthorized")
	}
	return nil
}
