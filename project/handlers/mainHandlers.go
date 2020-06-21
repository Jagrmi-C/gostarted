package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	lr "github.com/sirupsen/logrus"

	"github.com/Jagrmi-C/gostarted/project/db"
	"github.com/Jagrmi-C/gostarted/project/logger"
	"github.com/Jagrmi-C/gostarted/project/models"
)

func init() {
	logger.LoggerInitialization()
}

func Default(w http.ResponseWriter, req *http.Request) {
	testTime := time.RFC3339
	fmt.Println(testTime)

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
	uuid := mux.Vars(req)["uuid"]
	taskChan := make(chan models.TaskInformation)
	timeFrameChan := make(chan []models.TaskTimeFrame)

	go db.GetGoTask(uuid, taskChan)
	go db.GetGoTimeFramesByTask(uuid, timeFrameChan)

	taskInfo := <- taskChan
	taskInfo.TimeFrames = <- timeFrameChan

	lr.Info("Get all tasks use goroutines from DB")

	err := json.NewEncoder(w).Encode(taskInfo)
	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
