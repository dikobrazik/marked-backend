package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.Use(CommonMiddleware)

	//router.HandleFunc("/", middleware.TestAPI).Methods("GET", "OPTIONS")
	router.HandleFunc("/register", RegisterHandler).Methods("POST", "OPTIONS")
	// router.HandleFunc("/login", middleware.LoginHandler).Methods("POST", "OPTIONS")

	secure := router.PathPrefix("/auth").Subrouter()
	secure.Use(JwtVerify)

	// secure.HandleFunc("/api/task", middleware.GetAllTask).Methods("GET", "OPTIONS")
	// secure.HandleFunc("/api/task", middleware.CreateTask).Methods("POST", "OPTIONS")
	// secure.HandleFunc("/api/task/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	// secure.HandleFunc("/api/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	// secure.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	// secure.HandleFunc("/api/deleteAllTask", middleware.DeleteAllTask).Methods("DELETE", "OPTIONS")
	return router
}

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
