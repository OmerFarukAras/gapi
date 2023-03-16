package user

import (
	"gapi/controllers/user"
	"gapi/models"
	"gapi/util"
	"gapi/util/db"
	"net/http"
)

func RegisterRoute(res http.ResponseWriter, req *http.Request) {
	util.Write(res, "Hello, HTTP!\n")
	util.Write(res, "register, HTTP!\n")

	req.ParseForm()

	username := req.Form.Get("username")
	password := req.Form.Get("password")
	email := req.Form.Get("email")

	cr := user.RegisterController(res, username, email, password)
	if !cr {
		util.Write(res, "Incorrect form data, HTTP!\n")
		return
	}

	util.Info("USER DATA: ", username+" "+password+" "+email)
	db.CreateUser(models.User{
		Username: username,
		Email:    email,
		Password: util.ShaHash(password),
	})

	token := util.CreateToken(email)

	util.Write(res, token+", HTTP!\n")
	util.Info("Form DATA :", req.Form.Has("username"))
}
