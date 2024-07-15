package database

import (
	"APP/Model"
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func DatabaseConnection() (*gorm.DB, error) {
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost", 3306, "student_db")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Errorf("failed to connect database")
		return nil, errors.New("Failed to Connect to DB:" + err.Error())
	}

	fmt.Println("Database connected")
	db.AutoMigrate(model.Student{})

	return db, nil
}

