package main

import (
    "github.com/gorilla/mux"
	"log"
    "net/http"
    "github.com/Jagrmi-C/gostarted/courses/task_eight_server/handlers"
)

func main() {
    router := mux.NewRouter()
    log.Println(router)
    router.HandleFunc("/", handlers.HiHandler).Methods(http.MethodGet)
    router.HandleFunc("/increase", handlers.IncreseX2Handler).Methods(http.MethodPost)

    log.Println("Starting API server on 3001")
	if err := http.ListenAndServe("127.0.0.1:3001", router); err != nil {
		log.Fatal(err)
    }
}
