package Db

import (
	"github.com/matthewhartstonge/argon2"
)

var Connection *Database = nil
var Argon2 argon2.Config = argon2.DefaultConfig()
