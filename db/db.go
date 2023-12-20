package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/sohail-9098/ms-user-auth/user"
)

func connect() *sql.DB {
	connStr := "user=postgres password=pg@123 dbname=ms_user_auth sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}
	return db
}

func FetchUser(username string) user.User {
	db := connect()
	defer db.Close()

	query := "SELECT user_name, password FROM users WHERE user_name=$1"
	rows, err := db.Query(query, username)
	if err != nil {
		log.Fatal("query failed: ", err)
	}
	defer rows.Close()

	var user user.User
	for rows.Next() {
		err := rows.Scan(&user.Username, &user.Password)
		if err != nil {
			log.Fatal("scan failed: ", err)
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatal("error in rows: ", err)
	}

	return user
}
