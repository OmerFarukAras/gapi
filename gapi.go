package main

import (
	"gapi/routes"
	"gapi/routes/post"
	"gapi/routes/user"
	"gapi/util"
	"gapi/util/db"
	"net/http"
)

type HttpHandler struct{}

func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	datalog := req.Method + " Request : " + req.URL.Path + " - " + req.RemoteAddr + " :"
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Credentials", "true")
	res.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	db.AddLog(datalog)
	util.Info(datalog, nil)

	User := util.GetUser(req)

	url := req.URL

	if url.Path == "/" {
		routes.MainRoute(res, req)
	} else {
		if req.Method == "POST" {
			if url.Path == "/user/register" {
				user.RegisterRoute(res, req)
			} else if url.Path == "/user/login" {
				user.LoginRoute(res, req)
			} else if url.Path == "/user" {
				user.RoutePost(res, req)
			} else if url.Path == "/post/create" {
				post.CreateRoute(res, req, User)
			} else if url.Path == "/post/update" {
				post.UpdateRoute(res, req, User)
			} else if url.Path == "/post/delete" {
				post.DeleteRoute(res, req, User)
			} else {
				routes.E404Route(res, req)
			}
		} else if req.Method == "GET" {
			if url.Path == "/user" {
				user.RouteGet(res, req)
			} else if url.Path == "/post/read" {
				post.ReadRoute(res, req)
			} else {
				routes.E404Route(res, req)
			}
		} else {
			routes.E404Route(res, req)
		}
	}
}

func main() {
	handler := HttpHandler{}
	util.Info("Server started on : 8080", nil)
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		return
	}
}
