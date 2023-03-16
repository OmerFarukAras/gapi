package routes

import (
	"gapi/util"
	"net/http"
)

func E404Route(res http.ResponseWriter, req *http.Request) {
	util.Write(res, "Hello, HTTP!\n")
	util.Write(res, "Code 404, we can't fount this page. Try again later!\n")
}
