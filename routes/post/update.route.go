package post

import (
	"gapi/controllers/post"
	"gapi/models"
	"gapi/util"
	"gapi/util/db"
	"net/http"
)

func UpdateRoute(res http.ResponseWriter, req *http.Request, user *models.User) {
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

		err := req.ParseForm()
		if err != nil {
			return
		}

		if req.Form.Has("title") && req.Form.Has("content") && req.Form.Has("id") {

			title := req.Form.Get("title")
			content := req.Form.Get("content")
			postID := req.Form.Get("id")

			postData, ok := db.FindPostByID(postID)

			if ok {
				ok := post.UpdateController(res, title, content)
				if !ok {
					return
				}
				if user.CID == postData.Author {
					cr, postData := db.UpdatePost(title, content, postData)
					if !cr {
						resp["error"] = "DB Error in update."
					}

					resp["message"] = "Post updated."
					resp["post"] = postData.CID
					resp["author"] = postData.Author
					resp["createdAt"] = postData.CreatedAt
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
