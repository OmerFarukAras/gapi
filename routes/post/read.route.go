package post

import (
	"gapi/util"
	"gapi/util/db"
	"net/http"
	"strconv"
)

func ReadRoute(res http.ResponseWriter, req *http.Request) {
	itemCount := 10
	page := 1
	all := 0
	util.Info("data", itemCount+page)

	res.WriteHeader(http.StatusCreated)
	res.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)

	if req.URL.Query().Has("page") {
		page, _ = strconv.Atoi(req.URL.Query().Get("page"))
		if page <= 0 {
			page = 1
		}
	}

	if req.URL.Query().Has("all") {
		all, _ = strconv.Atoi(req.URL.Query().Get("all"))
		if all != 1 {
			all = 0
		}
	}

	posts, ok := db.GetAllPosts()

	if !ok {
		resp["error"] = "DB error"
		util.JsonWrite(res, resp)
		return
	}
	if all == 1 {
		str, _ := util.PostsArrayToString(posts)
		resp["data"] = str
		util.JsonWrite(res, resp)
		return
	} else if all == 0 {

		postsLength := len(posts)

		start := -10
		end := 0

		start += page * 10
		end += page * 10

		if start > postsLength {
			resp["error"] = "Page error"
			util.JsonWrite(res, resp)
			return
		} else {
			if end > postsLength {
				end = postsLength
			}
			str, _ := util.PostsArrayToString(posts[start:end])
			resp["data"] = str
			util.JsonWrite(res, resp)
			return
		}
	}
}
