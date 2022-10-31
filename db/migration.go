package db

import (
	"fmt"
	"log"
)

func MigrateDB() error {

	db, err := InitDb()
	if err != nil {
		log.Println("Could not establish connection with underlying Database")

	}

	fmt.Println("successful connection!!!")

	err = db.AutoMigrate(&Users{}).Error
	if err != nil {
		log.Println("Error while migrating Database Scheme", err.Error())

	}

	err = db.AutoMigrate(&Product{}).Error
	if err != nil {
		log.Println("Error while migrating Database Scheme", err.Error())

	}

	err = db.AutoMigrate(&UserProdActivity{}).Error
	if err != nil {
		log.Println("Error while migrating Database Scheme", err.Error())

	}

	err = db.AutoMigrate(&UserLoginActivity{}).Error
	if err != nil {
		log.Println("Error while migrating Database Scheme", err.Error())

	}

	err = db.AutoMigrate(&Token{}).Error
	if err != nil {
		log.Println("Error while migrating Database Scheme", err.Error())

	}

	return nil
}
