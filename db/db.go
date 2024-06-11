package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sohail-9098/ms-user-auth/user"
	"github.com/sohail-9098/ms-user-auth/util"
)

func connect() (*sql.DB, error) {
	// load DB config from file
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	// decode password
	config.Password, err = util.DecodePassword(config.Password)
	if err != nil {
		return nil, err
	}

	// prepare connection string and
	// connect to database
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", config.User, config.Password, config.DbName, config.SslMode)
	db, err := sql.Open(config.Db, connStr)
	if err != nil {
		return nil, err
	}

	//return db connection instance
	return db, nil
}

func FetchUser(username string) (user.Credentials, error) {
	var user user.Credentials
	// connect to db
	db, err := connect()
	if err != nil {
		return user, err
	}
	defer db.Close()

	// fetch user details
	// from db
	query := "SELECT user_name, password FROM users WHERE user_name=$1"
	rows, err := db.Query(query, username)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	// write to user struct
	for rows.Next() {
		err := rows.Scan(&user.Username, &user.Password)
		if err != nil {
			return user, err
		}
	}

	// check for errors during iteration
	if err := rows.Err(); err != nil {
		return user, err
	}

	return user, nil
}
