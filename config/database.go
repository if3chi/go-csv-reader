package config

import (
	"database/sql"
	logger "read-csv/services"
)

func DB_Connect() {
	db, err := sql.Open("sqlite3", "./people.db")
	logger.CheckErr(err)
	defer db.Close()
}
