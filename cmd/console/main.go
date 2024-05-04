package main

import (
	"apiAcademy/internal/database"
	"apiAcademy/internal/examples"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost" // 127.0.0.1
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "academy_db"
)

func main() {
	db, err := database.UsualConnect(host, port, user, password, dbname)
	if err != nil {
		log.Fatal("cannot connect to DB:", err.Error())
	}
	defer db.Close()

	//examples.Example1(db)
	//examples.Example2SqlInjection(db)

	//examples.StaticQuery(db)
	//examples.DynamicQuery(db)
	//examples.DynamicQueryPrepared(db)
	examples.DynamicQueryPreparedUsingStruct(db)
}
