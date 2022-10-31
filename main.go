package main

import (
	"ThirdEssentials/db"
	"ThirdEssentials/server"
	"log"
)

func main() {

	err := db.MigrateDB()
	if err != nil {
		log.Fatal(err)
	}

	r := server.InitServer()
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Cound not start server ", err)
	}

}
