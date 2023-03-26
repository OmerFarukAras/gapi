package post

import (
	"gapi/util"
	"gapi/util/db"
	"net/http"
)

func CreateController(res http.ResponseWriter, title string, content string) bool {
	resp := make(map[string]string)
	if len(title) < 5 && len(title) > 100 {
		resp["error"] = "Title length 5 - 100"
		util.JsonWrite(res, resp)
		return false
	}
	if len(content) < 10 && len(content) > 1000 {
		resp["error"] = "Content length 10 - 1000"
		util.JsonWrite(res, resp)
		return false
	}
	r, ok := db.FindPostByTitle(title)
	if ok && r.Title != "" {
		resp["error"] = "Title already using, try another title"
		util.JsonWrite(res, resp)
	}
	return true
}
