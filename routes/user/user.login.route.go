package user

import (
	"gapi/controllers/user"
	"gapi/util"
	"gapi/util/db"
	"net/http"
)

func LoginRoute(res http.ResponseWriter, req *http.Request) {
	headerContentType := req.Header.Get("Content-Type")
	if headerContentType != "application/x-www-form-urlencoded" {
		res.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	err := req.ParseForm()
	if err != nil {
		return
	}

	resp := make(map[string]string)

	password := req.Form.Get("password")
	email := req.Form.Get("email")

	cr := user.LoginController(res, email, password)
	if !cr {
		return
	}

	userData, ok := db.FindUserByEmail(email)
	if ok {
		if userData.Password == util.ShaHash(password) {
			token := util.CreateToken(email)
			resp["token"] = token
		} else {
			resp["error"] = "Password incorrect."
		}
	} else {
		resp["error"] = "User doesn't exist, HTTP!\n"
	}

	util.JsonWrite(res, resp)
	util.Info("FORM DATA: ", password+" "+email)
}
