package main

import (
    "github.com/gorilla/mux"
	"log"
    "net/http"
    "github.com/Jagrmi-C/gostarted/courses/task_eight_GP_server/handlers"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/api/v0/", handlers.GETRootHandler).Methods(http.MethodGet)
    router.HandleFunc("/api/v0/", handlers.POSTRootHandler).Methods(http.MethodPost)
    router.HandleFunc("/api/v0/logout/", handlers.DeleteCookie).Methods(http.MethodGet)

    log.Println("Starting API server on 3003")
	if err := http.ListenAndServe("127.0.0.1:3003", router); err != nil {
		log.Fatal(err)
    }
}
