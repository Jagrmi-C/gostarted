package handlers

import (
	// "strconv"
	"fmt"
	"context"
	"encoding/json"
	"net/http"
	"time"
	// "io/ioutil"
	"github.com/gorilla/mux"
	"github.com/Jagrmi-C/gostarted/project/db"
	"github.com/Jagrmi-C/gostarted/project/models"
)

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

func GetTaskHandler(w http.ResponseWriter, req *http.Request) {
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

func UpdateTaskHandler(w http.ResponseWriter, req *http.Request) {
	// create the postgres db connection
	conn := db.CreateConnection()

	// close the db connection
	defer conn.Close(context.Background())

	vars := mux.Vars(req)
	fmt.Println(vars, vars["uuid"])
	// uuid, err := strconv.ParseUint(vars["uuid"], 10, 32)
	uuid, ok := vars["uuid"]
	if ok {
		fmt.Println("Exists!")
	}
	fmt.Println(uuid)
	fmt.Println(req.Body)
	// body, err := ioutil.ReadAll(req.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	// responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }
	// fmt.Println("BODY", body)
	task := models.Task{}
	err := json.NewDecoder(req.Body).Decode(&task);
    if err != nil {
		fmt.Println("X", err)
	}
	fmt.Println(task)
	// err = json.Unmarshal(body, &task)
	// if err != nil {
		// responses.ERROR(w, http.StatusUnprocessableEntity, err)
		// return
	// }
	// fmt.Println(task)
	// fmt.1
}

func Default(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Ok! This is an example HTTPS server!\n")
}

func DefaultTest(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is SPARTA!\n")
}
