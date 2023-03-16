package user

import (
	"gapi/controllers"
	"gapi/util"
	"gapi/util/db"
	"net/http"
)

func RegisterController(res http.ResponseWriter, username string, email string, password string) bool {
	_, ok := controllers.ValidMailAddress(email)

	if len(username) < 3 && len(username) > 20 {
		util.Write(res, "Not Valid Username, HTTP!\n")
		return false
	}
	if !ok {
		util.Write(res, "Not Valid Email, HTTP!\n")
		return false
	}

	r, ok := db.FindUserByEmail(email)
	if ok && r.Email != "" {
		util.Write(res, "Email already using, try another email, HTTP!\n")
		return false
	}
	if len(password) < 3 && len(password) > 20 {
		util.Write(res, "Not Valid Password, HTTP!\n")
		return false
	}
	return true
}
