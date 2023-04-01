package post

import (
	"gapi/controllers/post"
	"gapi/models"
	"gapi/util"
	"gapi/util/db"
	"net/http"
)

func CreateRoute(res http.ResponseWriter, req *http.Request, user *models.User) {
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
		resp["error"] = "User not exist."
		util.JsonWrite(res, resp)
	} else {
		util.Info("USER ", user)
		if req.Form.Has("title") && req.Form.Has("content") {
			title := req.Form.Get("title")
			content := req.Form.Get("content")

			ok := post.CreateController(res, title, content)
			if !ok {
				//resp["error"] = "Invalid form data."
				//util.JsonWrite(res, resp)
				return
			}

			cr, postData := db.CreatePost(title, content, user.CID)
			if !cr {
				resp["error"] = "DB Error in create, HTTP!\n"
				util.JsonWrite(res, resp)
				return
			}
			resp["postData"] = postData.CID
			resp["author"] = postData.Author
			resp["createdAt"] = postData.CreatedAt
			util.JsonWrite(res, resp)
			return
		} else {
			resp["error"] = "Invalid form data."
			util.JsonWrite(res, resp)
			return
		}

	}
}
