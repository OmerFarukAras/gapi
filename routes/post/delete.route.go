package post

import (
	"gapi/models"
	"gapi/util"
	"gapi/util/db"
	"net/http"
)

func DeleteRoute(res http.ResponseWriter, req *http.Request, user *models.User) {
	headerContentType := req.Header.Get("Content-Type")
	if headerContentType != "application/x-www-form-urlencoded" {
		res.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	err := req.ParseForm()
	if err != nil {
		return
	}

	res.WriteHeader(http.StatusCreated)
	res.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)

	if user == nil {
		resp["error"] = "We not have a active user"
	} else {
		if req.Form.Has("id") {
			postID := req.Form.Get("id")
			post, ok := db.FindPostByID(postID)
			if ok {
				if user.CID == post.Author {
					cr, post := db.DeletePost(post)
					if !cr {
						resp["error"] = "DB Error in delete."
					}
					resp["message"] = "Post deleted."
					resp["post"] = post.CID
					resp["author"] = post.Author
					resp["createdAt"] = post.CreatedAt
				} else {
					resp["error"] = "you do not own this post."
				}
			} else {
				resp["error"] = "We dont have this post."
			}
		} else {
			resp["error"] = "Incorrect form data."
		}
	}
	util.JsonWrite(res, resp)
}
