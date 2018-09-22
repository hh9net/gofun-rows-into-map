package main

import (
	"encoding/json"
	log "log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	dbaccess "../pkg/db"
)

func main() {

	db, err := sqlx.Connect("mysql", "root:@(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := db.Queryx(`select * from test_table t1 where t1.version = 
		(select max(t2.version) from test_table t2 where t1.id = t2.id)`)

	ms, err := dbaccess.ScanIntoMaps(rows)

	log.Println("Displaying maps as json:")

	for i, m := range ms {
		jsonString, _ := json.Marshal(m)
		log.Printf("%d. %s", i, jsonString)
	}

}
