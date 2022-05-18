package main

import (
	"database/sql"
	"log"

	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 9000
	user     = "postgres"
	password = "Jee@2020"
	dbname   = "gostore"
)

func main() {
	sqlconnect := fmt.Sprint("host=%s port=%d password=%s dbname=%s sslmode=disable", host, port, password, dbname)

	db, err := sql.Open("postgres", sqlconnect)

	CheckError(err)

	defer db.Close()

	insrt := `insert into "coustomers"("UID","Name")" values(121,"Mit")`

	_, e := db.Exec(insrt)
	CheckError(e)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
