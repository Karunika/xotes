package main

import (
	"fmt"
	"net/http"

	"server/internal/db"
	"server/internal/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-chi/valve"
)

// //go:embed build/*
// var staticFiles embed.FS

func main() {
	dbURL := db.GenerateDBURL()
	dbService, _ := db.ConnectDB(dbURL)
	defer dbService.CloseDB()

	server := chi.NewRouter()
	server.Use(middleware.Logger)

	// server.Route("*", staticHandler())
	router := routes.NewRouterService(dbService)
	server.Mount("/api/v", router.APIRouter())

	fmt.Println("server started at port :3000")
	http.ListenAndServe(":3000", server)
}

// func staticHandler() http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		path := r.URL.Path
// 		if path == "/" {
// 			path = "/index.html"
// 		}
// 		content, err := staticFiles.ReadFile("build" + path)
// 		if err != nil {
// 			http.NotFound(w, r)
// 			return
// 		}
// 		contentType := http.DetectContentType(content)
// 		w.Header().Set("Content-Type", contentType)
// 		w.Write(content)
// 	})
// }
