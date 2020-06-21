package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	lr "github.com/sirupsen/logrus"

	"github.com/Jagrmi-C/gostarted/project/db"
	"github.com/Jagrmi-C/gostarted/project/models"
	"github.com/Jagrmi-C/gostarted/project/logger"
)

func init()  {
	logger.LoggerInitialization()
}

func Default(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Ok! This is an example HTTPS server!\n")
}

func DefaultTest(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is SPARTA!\n")
}


func GetGoTasksHandler(w http.ResponseWriter, req *http.Request) {
	output := make(chan []models.TaskInformation)

	go db.GetGoTasks(output)
	lr.Info("Get all tasksGO from DB")

	err := json.NewEncoder(w).Encode(<-output)
	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func GetGoTaskHandler(w http.ResponseWriter, req *http.Request) {
	output := make(chan []models.TaskInformation)

	go db.GetGoTasks(output)
	lr.Info("Get all tasksGO from DB")

	err := json.NewEncoder(w).Encode(<-output)
	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}