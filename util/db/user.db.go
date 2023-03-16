package db

import (
	"gapi/models"
	nanoid "github.com/matoous/go-nanoid/v2"
)

func CreateUser(user models.User) {
	driver := Database()
	id, _ := nanoid.New()
	user.CID = id
	user.Role = "0"
	err := driver.Insert(user)
	if err != nil {
		panic(err)
	}
	AddLog("User Created. - Username: " + user.Username + ", Email: " + user.Email + ", ID:" + user.CID)
}

func FindUserByEmail(email string) (*models.User, bool) {
	driver := Database()
	var user *models.User
	err := driver.Open(models.User{}).Where("email", "=", email).First().AsEntity(&user)
	if err != nil {
		return nil, false
	}
	return user, true
}

func FindUserByID(id string) (*models.User, bool) {
	driver := Database()
	var user *models.User
	err := driver.Open(models.User{}).Where("id", "=", id).First().AsEntity(&user)
	if err != nil {
		return nil, false
	}
	return user, true
}
