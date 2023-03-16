package user

import (
	"gapi/controllers"
	"gapi/util"
	"net/http"
)

func LoginController(res http.ResponseWriter, email string, password string) bool {
	_, ok := controllers.ValidMailAddress(email)

	if !ok {
		util.Write(res, "Not Valid Email, HTTP!\n")
		return false
	}

	if len(password) < 3 && len(password) > 20 {
		util.Write(res, "Not Valid Password, HTTP!\n")
		return false
	}
	return true
}
