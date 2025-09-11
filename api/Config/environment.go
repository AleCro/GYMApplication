package Config

import (
	"os"
)

var DATABASE_CONNECTION_STRING string

func Load() {
	GetEnv("DATABASE_CONNECTION_STRING", &DATABASE_CONNECTION_STRING, "mongodb://localhost:27017/")
}

func GetEnv(name string, into *string, def string) {
	value, exists := os.LookupEnv(name)
	if !exists {
		*into = def
		return
	}
	*into = value
}
