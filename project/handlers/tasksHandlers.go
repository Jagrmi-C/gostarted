package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	lr "github.com/sirupsen/logrus"

	"github.com/Jagrmi-C/gostarted/project/db"
	"github.com/Jagrmi-C/gostarted/project/logger"
	"github.com/Jagrmi-C/gostarted/project/models"
)

func init() {
	logger.LoggerInitialization()
}

func GetTaskHandler(w http.ResponseWriter, req *http.Request) {
	uuid := mux.Vars(req)["uuid"]
	var taskInfo models.TaskInformation

	err := db.GetTaskInfo(uuid, &taskInfo)
	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lr.Info("Get task from DB:", taskInfo.UUID)

	err = json.NewEncoder(w).Encode(taskInfo)
	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func GetTasksHandler(w http.ResponseWriter, req *http.Request) {
	var tasks []models.TaskInformation
	err := db.GetTasksInfo(&tasks)
	lr.Info("Get all tasks from DB")

	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func CreateTaskHandler(w http.ResponseWriter, req *http.Request) {
	var task models.Task
	err := json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.CreateTask(&task)

	lr.Info("Get task from DB:", task.UUID)
	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(task)

	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func UpdateTaskHandler(w http.ResponseWriter, req *http.Request) {
	uuid := mux.Vars(req)["uuid"]

	var task models.Task
	err := json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task.UUID = uuid

	err = db.UpdateTask(&task)
	lr.Info("Update task:", task)
	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func DeleteTaskHandler(w http.ResponseWriter, req *http.Request) {
	uuid := mux.Vars(req)["uuid"]

	err := db.DeleteTask(uuid)
	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lr.Info("Delete task from DB:", uuid)

	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
