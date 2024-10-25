package connection

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5438
	dbname="test-db-name"
	user = "test-user"
	password="test-password"
)

func Connect() (dbConnection *sql.DB) {
	psqlInfo := fmt.Sprintf("host=" + host + " port=" + strconv.Itoa(port) + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping() 
	if err != nil {
		panic(err)
	}

	fmt.Println("DB Connection successful")

	return db
}