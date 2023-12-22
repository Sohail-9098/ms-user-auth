package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sohail-9098/ms-user-auth/user"
	"github.com/sohail-9098/ms-user-auth/util"
)

func connect() *sql.DB {
	config, err := loadConfig()
	util.HandleError("error loading config: ", err)
	config.Password = util.DecodePassword(config.Password)
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", config.User, config.Password, config.DbName, config.SslMode)
	db, err := sql.Open(config.Db, connStr)
	util.HandleError("error connecting to database: ", err)
	return db
}

func FetchUser(username string) user.User {
	db := connect()
	defer db.Close()

	query := "SELECT user_name, password FROM users WHERE user_name=$1"
	rows, err := db.Query(query, username)
	util.HandleError("query failed: ", err)
	defer rows.Close()

	var user user.User
	for rows.Next() {
		err := rows.Scan(&user.Username, &user.Password)
		util.HandleError("scan failed: ", err)
	}

	if err := rows.Err(); err != nil {
		util.HandleError("query failed: ", err)
	}

	return user
}
