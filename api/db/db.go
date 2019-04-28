package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // postgres driver, dickheads
)

const DB_URL_VAR = "DATABASE_URL"

func Open() (*sql.DB, error) {
	dbURL, ok := os.LookupEnv(DB_URL_VAR)
	if !ok {
		return nil, fmt.Errorf("%s not set", DB_URL_VAR)
	}
	return sql.Open("postgres", dbURL)
}
