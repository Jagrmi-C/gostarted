package router

import (
	"net/http"
    "github.com/Jagrmi-C/gostarted/project/handlers"
    "github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

    router := mux.NewRouter()

    // router.HandleFunc("/api/user/{id}", middleware.GetUser).Methods("GET", "OPTIONS")
    // router.HandleFunc("/api/user", middleware.GetAllUser).Methods("GET", "OPTIONS")
    // router.HandleFunc("/api/newuser", middleware.CreateUser).Methods("POST", "OPTIONS")
    // router.HandleFunc("/api/user/{id}", middleware.UpdateUser).Methods("PUT", "OPTIONS")
    // router.HandleFunc("/api/deleteuser/{id}", middleware.DeleteUser).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/v0/", handlers.DefaultTest).Methods(http.MethodGet)
	router.HandleFunc("/", handlers.Default).Methods(http.MethodGet)

	router.HandleFunc("/groups/", handlers.Default).Methods(http.MethodGet)
	router.HandleFunc("/groups/", handlers.Default).Methods(http.MethodPost)
	router.HandleFunc("/tasks/", handlers.Default).Methods(http.MethodGet)
	router.HandleFunc("/tasks/", handlers.Default).Methods(http.MethodPost)

	router.HandleFunc("/timeframes/", handlers.Default).Methods(http.MethodPost)
	router.HandleFunc("/tasks/{id}", handlers.GetTaskHandler).Methods(http.MethodGet)
	router.HandleFunc("/tasks/{uuid}", handlers.UpdateTaskHandler).Methods(http.MethodPut)
	router.HandleFunc("/tasks/{id}", handlers.Default).Methods(http.MethodDelete)

	router.HandleFunc("/groups/{id}", handlers.Default).Methods(http.MethodPut)
	router.HandleFunc("/groups/{id}", handlers.Default).Methods(http.MethodDelete)

    return router
}