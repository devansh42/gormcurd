package main

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//User, model for User
type User struct {
	gorm.Model
	Status                     bool
	Email                      string `gorm:"size:100"`
	EmailVerified              bool
	MobileNo                   string `gorm:"size:15"`
	MobileVerfied              bool
	UserName                   string `gorm:"size:20"`
	Password                   string `gorm:"size:32;type:varchar"` //SHA 256 Salt
	FullName                   string `gorm:"size:100"`
	NameInitials               string
	LastPasswordUpdateDatetime time.Time
}

//BasicUserProfile, model for Basic User Info
type BasicUserProfile struct {
	gorm.Model
	User               User
	UserId             string
	Bio                string `gorm:"size:255"`
	Hometown           string `gorm:"size:50"`
	University         string `gorm:"size:50"`
	Work               string `gorm:"size:50"`
	RelationshipStatus string `gorm:"size:50"`
	Music              string `gorm:"size:50"`
	Link               string `gorm:"size:150"`
}

func getDataBase() (*gorm.DB, error) {
	//host=localhost sslmode=disable port=5432 user=postgres  password=root1234  dbname=demo_database
	return gorm.Open("sqlite3", "test.db")

}

//createDefaultSchema, is a default method to create default database schema
func createDefaultSchema(db *gorm.DB) {
	var err error
	if db == nil {
		db, err = getDataBase()
		if err != nil {
			//Handle Error
			log.Fatal("Couldn't create default schema due to ", err.Error())
		}

	}
	for _, v := range db.CreateTable(User{}).GetErrors() {
		log.Fatal("Error Occured during default schema creation, ", v.Error())
	}
}
