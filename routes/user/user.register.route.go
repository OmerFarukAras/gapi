package user

import (
	"gapi/controllers/user"
	"gapi/models"
	"gapi/util"
	"gapi/util/db"
	"net/http"
)

func RegisterRoute(res http.ResponseWriter, req *http.Request) {
	headerContentType := req.Header.Get("Content-Type")
	if headerContentType != "application/x-www-form-urlencoded" {
		res.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	req.ParseForm()

	res.WriteHeader(http.StatusCreated)
	res.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)

	username := req.Form.Get("username")
	password := req.Form.Get("password")
	email := req.Form.Get("email")

	cr := user.RegisterController(res, username, email, password)
	if !cr {
		return
	}

	util.Info("USER DATA: ", username+" "+password+" "+email)
	db.CreateUser(models.User{
		Username: username,
		Email:    email,
		Password: util.ShaHash(password),
	})

	token := util.CreateToken(email)

	resp["token"] = token
	util.JsonWrite(res, resp)
	util.Info("Form DATA :", req.Form.Has("username"))
}
