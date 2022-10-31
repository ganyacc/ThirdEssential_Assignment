package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// To connect to database
func InitDb() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:Ganyayadav@123@tcp(127.0.0.1:3306)/Ecom?charset=utf8&parseTime=True&loc=Local")
	return db, err
}
