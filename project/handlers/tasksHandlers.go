package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	lr "github.com/sirupsen/logrus"

	"github.com/Jagrmi-C/gostarted/project/db"
	"github.com/Jagrmi-C/gostarted/project/models"
	"github.com/Jagrmi-C/gostarted/project/logger"
)

func init()  {
	logger.LoggerInitialization()
}

type ResponseCustom struct {
	Q1	   string `json:"q1"`
	Q2	   time.Time `json:"q2"`
	Host       string `json:"host"`
	UserAgent  string `json:"user_agent"`
	RequestURI string `json:"request_uri"`
	Headers    struct {
		Accept    []string `json:"Accept"`
		UserAgent []string `json:"User-Agent"`
	} `json:"headers"`
}

func GetTaskHandlerMy(w http.ResponseWriter, req *http.Request) {
	// create the postgres db connection
	conn := db.CreateConnection()

	// close the db connection
	defer conn.Close(context.Background())

	var group string
	var dt time.Time

	res := ResponseCustom{}

	res.Host = req.Host
	res.UserAgent = req.Header["User-Agent"][0]
	res.RequestURI = req.RequestURI

	res.Headers.Accept = req.Header["Accept"]
	res.Headers.UserAgent = req.Header["User-Agent"]

	err := conn.QueryRow(
		context.Background(),
		"select group_uuid, dt from tasks where title=$1",
		"mouse",
	).Scan(&group, &dt)

	if err != nil {
		fmt.Println("err db select")
	}

	res.Q1 = group
	res.Q2 = dt
	js, _ := json.Marshal(res)
	lenght, err := w.Write(js)
	if err != nil {
		panic("Panic")
	}
	fmt.Println("result", lenght)
}

func GetTaskHandler(w http.ResponseWriter, req *http.Request) {
	uuid := mux.Vars(req)["uuid"]
	var task models.Task

	err := db.GetTask(uuid, &task)
	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lr.Info("Get task from DB:", task.UUID)

	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func GetTasksHandler(w http.ResponseWriter, req *http.Request) {
	tasks, err := db.GetTasks()
	lr.Info("Get all tasks from DB")

	var bodyStruct models.TasksStruct

	bodyStruct.Task = tasks
	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(bodyStruct)
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

func Default(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Ok! This is an example HTTPS server!\n")
}

func DefaultTest(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is SPARTA!\n")
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
