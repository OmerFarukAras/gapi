package post

import (
	"gapi/controllers/post"
	"gapi/models"
	"gapi/util"
	"gapi/util/db"
	"net/http"
)

func CreateRoute(res http.ResponseWriter, req *http.Request, user *models.User) {
	util.Write(res, "Hello, HTTP!\n")
	util.Write(res, "create, HTTP!\n")
	if user == nil {
		util.Write(res, "We not have a active user, HTTP!\n")
	} else {
		util.Write(res, user.Username+", HTTP!\n")
		util.Info("USER ", user)

		req.ParseForm()

		if req.Form.Has("title") && req.Form.Has("content") {

			title := req.Form.Get("title")
			content := req.Form.Get("content")

			ok := post.CreateController(res, title, content)
			if !ok {
				util.Write(res, "Incorrect form data, HTTP!\n")
				return
			}

			cr, post := db.CreatePost(title, content, user.CID)
			if !cr {
				util.Write(res, "DB Error in create, HTTP!\n")
				return
			}
			util.Write(res, "Post Created, HTTP!\n")
			util.Write(res, "Post ID: "+post.CID+", HTTP!\n")
		} else {
			util.Write(res, "Incorrect form data, HTTP!\n")
			return
		}

	}
}
