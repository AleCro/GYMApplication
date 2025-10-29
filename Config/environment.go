package Config

import (
	"os"
)

const DATABASE_NAME string = "gymapp"
const DATABASE_USER_COLLECTION string = "usersgym"
const DATABASE_SESSION_COLLECTION string = "sessionsgym"

var DATABASE_CONNECTION_STRING string

var SERVER_PORT string = "7284"

var UPLOAD_DIR string = "string"

func Load() {
	GetEnv("DATABASE_CONNECTION_STRING", &DATABASE_CONNECTION_STRING, "mongodb://localhost:27017/")
	GetEnv("UPLOAD_DIR", &UPLOAD_DIR, "uploads")
}

func GetEnv(name string, into *string, def string) {
	value, exists := os.LookupEnv(name)
	if !exists {
		*into = def
		return
	}
	*into = value
}
