package main

import (
	"blog-backend/db"
	"blog-backend/handler"
	"blog-backend/services/bsm"
	"database/sql"
	"net/http"
)

const (
	CREATE_BLOG = "POST /create/blog"
	GET_BLOG_BY_ID = "POST /get/blog/id"
	GET_BLOG_BY_FILTER = "POST /get/blog/tags"
	DELETE_BLOG_BY_ID = "POST /delete/blog/id"
	UPDATE_BLOG_BY_ID = "POST /update/blog/id"
)

type Server struct{
	Path		string
	Port		string
	DBClient	*sql.DB
}

func NewServer(path string, port string, db *sql.DB) *Server{
	return &Server{
		Path:  		path,
		Port:		port,
		DBClient:	db,
	}
}

func (s *Server) ServiceSetup() error{
	serveMux := http.NewServeMux()

	// Create an instance of the services to be used
	blogRepository := db.NewBlogRepository(s.DBClient)
	blogService := bsm.NewBlogService(blogRepository)

	// Add the service handler here
	serveMux.HandleFunc(CREATE_BLOG, handler.NewCreateBlogHandler(blogService).Handler)
	serveMux.HandleFunc(GET_BLOG_BY_ID, handler.NewGetBlogByIDHandler(blogService).Handler)
	serveMux.HandleFunc(GET_BLOG_BY_FILTER, handler.NewGetBlogByTagsHandler(blogService).Handler)
	serveMux.HandleFunc(DELETE_BLOG_BY_ID, handler.NewDeleteBlogHandler(blogService).Handler)
	serveMux.HandleFunc(UPDATE_BLOG_BY_ID, handler.NewUpdateBlogHandler(blogService).Handler)

	return http.ListenAndServe(s.Path + ":" + s.Port, serveMux)
}