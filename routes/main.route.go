package routes

import (
	"gapi/util"
	"net/http"
)

func MainRoute(res http.ResponseWriter, req *http.Request) {
	util.Write(res, "Hello, HTTP!\n")
	util.Write(res, "/, HTTP!\n")
	util.Write(res, "api v1, register/login - create/read/update/delete CRUD!\n")
}
