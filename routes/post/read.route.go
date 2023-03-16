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
		util.Write(res, "DB Error in read, HTTP!\n")
		return
	}
	if all == 1 {
		str, _ := util.PostsArrayToString(posts)
		util.Write(res, str)
	} else if all == 0 {

		postsLength := len(posts)

		start := -10
		end := 0

		start += page * 10
		end += page * 10

		if start > postsLength {
			util.Write(res, "Page error")
		} else {
			if end > postsLength {
				end = postsLength
			}
			str, _ := util.PostsArrayToString(posts[start:end])
			util.Write(res, str)
		}
	}
}
