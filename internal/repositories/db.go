package repositories

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" //postgres driver
	"cart/helper"
)


func DataBaseConnection(mode, dbhost, user, dbname, password string) {
	conn, err := databaseSetUp(mode, dbhost, user, dbname, password)
	if err != nil {
		helper.PrintErrorMessage("404", err.Error())
		return
	}
	NewcartInfra(conn)
}

func databaseSetUp(mode, dbhost, user, dbname, password string) (*gorm.DB, error){
	dbURL := databaseCredentials(mode, dbhost, user, dbname, password)

	db, err := gorm.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	fmt.Println("Core database connection successful")

	autoMigrateTables(db)
	return db, nil
}

func databaseCredentials(mode, dbhost, user, dbname, password string) string {
	if mode == "dev" {

		dbURL := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbhost, user, dbname, password)

		return dbURL
	}

	return errors.New("error getting environment").Error()
}

func autoMigrateTables(conn *gorm.DB) {
	conn.AutoMigrate("")
}
