package post

import (
	"gapi/models"
	"gapi/util"
	"gapi/util/db"
	"net/http"
)

func DeleteRoute(res http.ResponseWriter, req *http.Request, user *models.User) {
	util.Write(res, "Hello, HTTP!\n")
	util.Write(res, "delete, HTTP!\n")
	if user == nil {
		util.Write(res, "We not have a active user, HTTP!\n")
	} else {
		util.Write(res, user.Username+", HTTP!\n")
		util.Info("USER ", user)

		req.ParseForm()

		if req.Form.Has("id") {

			postID := req.Form.Get("id")

			cr, post := db.DeletePost(postID)
			if !cr {
				util.Write(res, "DB Error in delete, HTTP!\n")
				return
			}
			util.Write(res, "Post Deleted, HTTP!\n")
			util.Write(res, "Post ID: "+post.CID+", HTTP!\n")
		} else {
			util.Write(res, "Incorrect form data, HTTP!\n")
			return
		}

	}
}
