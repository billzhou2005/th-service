package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"

//const DB_USERNAME = "bill"
const DB_PASSWORD = "zhoumb1202"
const DB_NAME = "gin_gorm"

//define for docker qmm-mysql
//const DB_HOST = "docker.for.mac.localhost"

// const DB_HOST = "192.168.5.14"
// const DB_PORT = "3306"

const DB_HOST = "140.143.149.188"
const DB_PORT = "3306"

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)
		return nil
	}

	return db
}
