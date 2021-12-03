package models

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Key struct {
	gorm.Model
	Name        string
	Description string
	Status      string
}

type User struct {
	gorm.Model
	ID           int
	Username     string
	Email        string
	DisplayName  string
	PasswordHash string
	CanLogin     bool
}

type Assignment struct {
	gorm.Model
	ID      int
	User    string
	Key     string
	DateOut time.Time
	DateIn  time.Time
}

var DB *gorm.DB

func DBConnect() {
	db, err := gorm.Open(sqlite.Open("keymaster.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db.")
	}

	db.AutoMigrate(&Key{})

	// newWorkout := Workout{user: "Mike", name: "A", workout_start: time.Now()}
	// db.Create(&newWorkout)
	// db.Create(&Key{User: "Mike", Name: "A", Workout_start: time.Now()})

	// var firstWorkout Workout
	// db.First(&firstWorkout, 1)
	// fmt.Println(firstWorkout)
	DB = db
}
