package user

import (
	"gapi/controllers/user"
	"gapi/util"
	"gapi/util/db"
	"net/http"
)

func LoginRoute(res http.ResponseWriter, req *http.Request) {
	util.Write(res, "Hello, HTTP!\n")
	util.Write(res, "login, HTTP!\n")

	req.ParseForm()

	password := req.Form.Get("password")
	email := req.Form.Get("email")

	cr := user.LoginController(res, email, password)
	if !cr {
		util.Write(res, "Incorrect form data, HTTP!\n")
		return
	}

	user, ok := db.FindUserByEmail(email)
	if ok {
		if user.Password == util.ShaHash(password) {
			util.Write(res, "valid user, HTTP!\n")
			token := util.CreateToken(email)
			util.Write(res, token+", HTTP!\n")
		} else {
			util.Write(res, "Password incorrect, HTTP!\n")
		}
	} else {
		util.Write(res, "User doesn't exist, HTTP!\n")
	}

	util.Info("FORM DATA: ", password+" "+email)
}
