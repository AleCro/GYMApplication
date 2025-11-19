package Environment

import "time"

var (
	// Where the database will connect to
	DATABASE_URL  string = GetEnv("DATABASE_URL", "mongodb://127.0.0.1:27017/")
	// The name of the MongoDB database.
	DB_NAME string = GetEnv("DB_NAME", "svelgok")
	// The name of the collection in which users will be stored.
	DB_USERS_COLLECTION string = GetEnv("DB_USERS_COLLECTION", "users")
	// The name of the collection in which sessions will be stored.
	DB_SESSIONS_COLLECTION string = GetEnv("DB_SESSIONS_COLLECTION", "sessions")
	// The name of the TTL index (to make it concise and prevent multiple indexes).
	DB_SESSIONS_TLL_INDEX_NAME string = GetEnv("DB_SESSIONS_TLL_INDEX_NAME", "session-expiration")

	// The string that will be used as the secret for signing and verifying JWT tokens
	//
	// WARNING: This should be a strong, random string in a production environment.
	JWT_SECRET string = GetEnv("JWT_SECRET", "your-super-strong-random-secret-key")
	// How long a JWT token can last
	JWT_TOKEN_LIFESPAN time.Duration = time.Minute * time.Duration(GetEnvInt("JWT_TOKEN_LIFESPAN", 15))
	// How long after a token expires can it be renewed
	JWT_RENEWAL_GRACE_PERIOD time.Duration = time.Minute * time.Duration(GetEnvInt("JWT_RENEWAL_GRACE_PERIOD", 15))
	// How long a session (in the database) can last.
	SESSION_DURATION time.Duration = time.Minute * time.Duration(GetEnvInt("SESSION_DURATION", 15))

	// From what host(s) will request(s) be accepted
	//
	// WARNING: A value of `*` is insecure and should only be used for development.
	API_CORS_ORIGIN      string = GetEnv("API_CORS_ORIGIN", "*")
	API_CORS_CREDENTIALS string = GetEnv("API_CORS_CREDENTIALS", "true")
	API_CORS_HEADERS     string = GetEnv("API_CORS_HEADERS", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	API_CORS_METHODS     string = GetEnv("API_CORS_METHODS", "GET, POST, PATCH, PUT, DELETE, OPTIONS")

	// Enables extra, stricter checks on session and token validity. When true,
	// handlers will manually check if a session or token has expired relative to
	// its creation time + offset, in addition to standard expiration checks.
	STRICT_SESSION_CONSISTENCY        bool          = GetEnvBool("STRICT_SESSION_CONSISTENCY", false)
	STRICT_SESSION_CONSISTENCY_OFFSET time.Duration = time.Minute * time.Duration(GetEnvInt("STRICT_SESSION_CONSISTENCY_OFFSET", 2))
)
