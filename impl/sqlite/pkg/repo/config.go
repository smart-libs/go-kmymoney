package repo

import (
	"os"
)

type (
	Config struct {
		SQLite3Source string `json:"sqlite3_db"`
	}
)

func NewConfig() Config {
	return Config{SQLite3Source: os.Getenv("SQLITE3_DB")}
}
