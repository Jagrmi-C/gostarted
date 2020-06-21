package router

import (
	"net/http"
    "github.com/Jagrmi-C/gostarted/project/handlers"
    "github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

    router := mux.NewRouter()

	router.HandleFunc("/api/v0/", handlers.DefaultTest).Methods(http.MethodGet)
	router.HandleFunc("/", handlers.Default).Methods(http.MethodGet)

	router.HandleFunc("/tasks", handlers.GetTasksHandler).Methods(http.MethodGet)
	router.HandleFunc("/tasks/", handlers.CreateTaskHandler).Methods(http.MethodPost)
	router.HandleFunc("/tasks/{uuid}", handlers.GetTaskHandler).Methods(http.MethodGet)
	router.HandleFunc("/tasks/{uuid}", handlers.UpdateTaskHandler).Methods(http.MethodPut)
	router.HandleFunc("/tasks/{uuid}", handlers.DeleteTaskHandler).Methods(http.MethodDelete)

	router.HandleFunc("/groups", handlers.GetGroupsHandler).Methods(http.MethodGet)
	router.HandleFunc("/groups/", handlers.CreateGroupHandler).Methods(http.MethodPost)
	router.HandleFunc("/groups/{uuid}", handlers.GetGroupHandler).Methods(http.MethodGet)
	router.HandleFunc("/groups/{uuid}", handlers.UpdateGroupHandler).Methods(http.MethodPut)
	router.HandleFunc("/groups/{uuid}", handlers.DeleteGroupHandler).Methods(http.MethodDelete)

	router.HandleFunc("/timeframes/", handlers.CreateTimeframeHandler).Methods(http.MethodPost)
	router.HandleFunc("/timeframes/{uuid}", handlers.GetTimeframeHandler).Methods(http.MethodGet)
	router.HandleFunc("/timeframes/{uuid}", handlers.DeleteTimeframeHandler).Methods(http.MethodDelete)

	router.HandleFunc("/tasksgo/", handlers.GetGoTasksHandler).Methods(http.MethodGet)
    return router
}