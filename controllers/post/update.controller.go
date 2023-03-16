package post

import (
	"gapi/util"
	"net/http"
)

func UpdateController(res http.ResponseWriter, title string, content string) bool {
	if len(title) < 5 && len(title) > 100 {
		util.Write(res, "Title length 5 - 100 , HTTP!\n")
		return false
	}
	if len(content) < 10 && len(content) > 1000 {
		util.Write(res, "Content length 10 - 1000 , HTTP!\n")
		return false
	}

	return true
}
