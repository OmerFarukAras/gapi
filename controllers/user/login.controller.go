package user

import (
	"gapi/controllers"
	"gapi/util"
	"net/http"
)

func LoginController(res http.ResponseWriter, email string, password string) bool {
	_, ok := controllers.ValidMailAddress(email)
	resp := make(map[string]string)
	if !ok {
		resp["error"] = "Not Valid Email"
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
