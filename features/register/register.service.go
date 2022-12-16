package register

import (
	"post-app/configs"
	"post-app/entities"
)

type RegisterService struct {}

func (rs* RegisterService) CreateUser(u* entities.User) {
	db := configs.DBConfig()

	user := db.First(&u)
	println(user)
}