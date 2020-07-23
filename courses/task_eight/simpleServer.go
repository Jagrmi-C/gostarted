package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Jagrmi-C/gostarted/courses/lib"
)

type ResponseCustom struct {
	Host       string `json:"host"`
	UserAgent  string `json:"user_agent"`
	RequestURI string `json:"request_uri"`
	Headers    struct {
		Accept    []string `json:"Accept"`
		UserAgent []string `json:"User-Agent"`
	} `json:"headers"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	res := ResponseCustom{}

	res.Host = r.Host
	res.UserAgent = r.Header["User-Agent"][0]
	res.RequestURI = r.RequestURI

	res.Headers.Accept = r.Header["Accept"]
	res.Headers.UserAgent = r.Header["User-Agent"]

	js, err := json.Marshal(res)
	lib.OnErrFail(err)

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	lib.OnErrFail(err)
}

func main() {
	http.HandleFunc("/", jsonHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
