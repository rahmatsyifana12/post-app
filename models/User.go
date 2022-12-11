package models

import "time"

type User struct {
	ID 			string `gorm:"size:255;primaryKey"`
	Username 	string `gorm:"size:255"`
	Password 	string `gorm:"size:255;not null"`
	Name 		string `gorm:"size:255;not null"`
	Phone 		string `gorm:"size:63"`
	CreatedAt 	time.Time `gorm:"not null"`
  	UpdatedAt 	time.Time `gorm:"not null"`
	Posts		[]Post
}