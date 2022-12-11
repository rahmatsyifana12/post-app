package models

import "time"

type Post struct {
	ID 			string `gorm:"size:255;primaryKey"`
	UserID		string `gorm:"size:255;not null"`
	Title		string `gorm:"size:255;not null"`
	Content		string `gorm:"size:1023;not null"`
	CreatedAt 	time.Time `gorm:"not null"`
  	UpdatedAt 	time.Time `gorm:"not null"`
}