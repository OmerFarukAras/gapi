package user

import (
	"gapi/util"
	"gapi/util/db"
	"net/http"
	"strings"
)

func RoutePost(res http.ResponseWriter, req *http.Request) {
	headerContentType := req.Header.Get("Content-Type")
	if headerContentType != "application/x-www-form-urlencoded" {
		res.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	req.ParseForm()

	res.WriteHeader(http.StatusCreated)
	res.Header().Set("Content-Type", "application/json")

	tokenString := req.Form.Get("Authorization")

	tokenString = strings.Split(tokenString, " ")[1]
	data, ok := util.ParseToken(tokenString)
	resp := make(map[string]string)
	if ok {
		data, ok := db.FindUserByEmail(data.Email)
		if ok {
			resp["user"] = "active"
			resp["email"] = data.Email
			resp["username"] = data.Username
			resp["id"] = data.CID
			resp["role"] = data.Role
		} else {
			resp["error"] = "User doesn't found"
		}
	} else {
		resp["error"] = "User doesn't found"
	}
	util.JsonWrite(res, resp)
}

func RouteGet(res http.ResponseWriter, req *http.Request) {
	util.Write(res, "Hello, HTTP!\n")
	util.Write(res, "user, HTTP!\n")

	tokenString := req.Header.Get("Authorization")

	util.Info("HI", tokenString)

	tokenString = strings.Split(tokenString, " ")[1]
	data, ok := util.ParseToken(tokenString)
	if ok {
		util.Write(res, data.Email+", HTTP!\n")
	} else {
		util.Write(res, "We not have a active user, HTTP!\n")
	}
}
