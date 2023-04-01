package user

import (
	"gapi/controllers"
	"gapi/util"
	"gapi/util/db"
	"net/http"
)

func RegisterController(res http.ResponseWriter, username string, email string, password string) bool {
	_, ok := controllers.ValidMailAddress(email)
	resp := make(map[string]string)

	if len(username) < 3 && len(username) > 20 {
		resp["error"] = "Not Valid Username"
		util.JsonWrite(res, resp)
		return false
	}
	if !ok {
		resp["error"] = "Not Valid Email"
		util.JsonWrite(res, resp)
		return false
	}

	r, ok := db.FindUserByEmail(email)
	if ok && r.Email != "" {
		resp["error"] = "Email already using try another email."
		util.JsonWrite(res, resp)
		return false
	}
	if len(password) < 3 && len(password) > 20 {
		resp["error"] = "Not Valid Password"
		util.JsonWrite(res, resp)
		return false
	}
	return true
}
