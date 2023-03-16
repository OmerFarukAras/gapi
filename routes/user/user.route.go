package user

import (
	"gapi/util"
	"net/http"
	"strings"
)

func RoutePost(res http.ResponseWriter, req *http.Request) {
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
