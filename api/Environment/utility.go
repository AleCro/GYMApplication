package Environment

import (
	"os"
	"strconv"
	"strings"
)

func GetEnv(name, def string) string {
	s := os.Getenv(name)
	if s == "" || s == " " {
		return def
	}
	return strings.TrimSpace(s)
}

func GetEnvBool(name string, def bool) bool {
	s := os.Getenv(name)
	if s == "" || s == " " {
		return def
	}

	if strings.EqualFold(strings.TrimSpace(s), "true") {
		return true
	} else if strings.EqualFold(strings.TrimSpace(s), "false") {
		return false
	} else {
		return def
	}

	return def
}

func GetEnvInt(name string, def int) int {
	s := GetEnv(name, "")
	if s == "" || s == " " {
		return def
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return n
}
